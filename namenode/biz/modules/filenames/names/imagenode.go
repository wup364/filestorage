package names

import (
	"errors"
	"namenode/ifilestorage"
	"pakku/utils/utypes"
	"sort"
	"sync"
)

// TMNode 元数据信息(内存树)
type TMNode struct {
	id     string
	addr   string
	name   string
	flag   int8
	size   int64
	ctime  int64
	mtime  int64
	props  string
	parent *TMNode
	lock   *sync.RWMutex
	child  *utypes.SafeMap // map[name]TMNode
}

// Lock 锁定节点-独写
func (node *TMNode) Lock() {
	node.lock.Lock()
}

// RLock 锁定节点-共读
func (node *TMNode) RLock() {
	node.lock.RLock()
}

// Unlock 解锁-独写
func (node *TMNode) UnLock() {
	node.lock.Unlock()
}

// RUnlock 解锁-共读
func (node *TMNode) RUnLock() {
	node.lock.RUnlock()
}

// GetChildByName 根据名字获取子节点
func (node *TMNode) GetChildByName(name string) *TMNode {
	node.RLock()
	defer node.RUnLock()
	if len(name) > 0 && nil != node.child {
		if val, ok := node.child.Get(name); ok {
			return val.(*TMNode)
		}
	}
	return nil
}

// GetChilds 获取子列表, limit<=0时返回所有数据
func (node *TMNode) GetChilds(limit int, offset int) (nodes []*TMNode, count int) {
	node.RLock()
	defer node.RUnLock()
	if nil != node.child {
		count = node.child.Size()
		if limit <= 0 {
			limit = count
			offset = 0
		} else if limit > count {
			limit = count
		}
		if offset < 0 {
			offset = 0
		}
		if offset < count {
			index := 0
			keys := make([]string, count)
			node.child.DoRange(func(key, _ interface{}) error {
				keys[index] = key.(string)
				index++
				return nil
			})
			sort.Strings(keys)
			index = 0
			nodes = make([]*TMNode, limit)
			for i := offset; i < count; i++ {
				if val, ok := node.child.Get(keys[i]); ok {
					nodes[index] = val.(*TMNode)
					index++
					if index == limit {
						break
					}
				}
			}
			if index < limit {
				nodes = nodes[:index]
			}
		}
	}
	return nodes, count
}

// GetFileSize 同步阻塞获取文件|文件夹大小(重新计算)
func (node *TMNode) GetFileSize() int64 {
	if node.flag == FLAG_DIR {
		if nil == node.child {
			return 0
		}
		size := int64(0)
		node.child.DoRange(func(_, val interface{}) error {
			cnode := val.(*TMNode)
			if cnode.flag == FLAG_DIR {
				size += cnode.GetFileSize()
			} else {
				size += cnode.size
			}
			return nil
		})
		return size
	}
	return node.size
}

// AddChild 添加子节点
func (node *TMNode) AddChild(nodedto ifilestorage.TNode4New) (*TMNode, error) {
	if node.flag == FLAG_DIR {
		node.Lock()
		defer node.UnLock()
		if nil == node.child {
			node.child = utypes.NewSafeMap()
		} else {
			if node.child.ContainsKey(nodedto.Name) {
				return nil, errors.New("node.Name is exist, name: " + nodedto.Name)
			}
		}
		insert := new(TMNode).ReverseNode4New(node, nodedto)
		if err := node.child.PutX(nodedto.Name, insert); nil == err {
			node.updateDirSize(nodedto.Size, "")
			return insert, nil
		} else {
			return nil, err
		}
	}
	return nil, errors.New("node is not dir")
}

// Update 更新节点信息
func (node *TMNode) Update(nodedto ifilestorage.TNode4Update) error {
	node.Lock()
	defer node.UnLock()
	if nodedto.Name != node.name {
		node.parent.Lock()
		defer node.parent.UnLock()
		if err := node.parent.child.PutX(nodedto.Name, node); nil != err {
			return err
		}
		node.parent.child.Delete(node.name)
	}
	node.name = nodedto.Name
	node.mtime = nodedto.Mtime
	node.props = nodedto.Props
	// 文件才更新这些数据
	if node.flag == FLAG_FILE {
		updateSize := int64(nodedto.Size - node.size)
		node.addr = nodedto.Addr
		node.size = nodedto.Size
		node.parent.updateDirSize(updateSize, "")
	}
	return nil
}

// MoveChild 移动节点到另外一个节点下面
func (node *TMNode) MoveChild(destnode *TMNode, newname string) error {
	if nil == destnode {
		return errors.New("dest node is not exists")
	}
	if node.id == destnode.id || node.parent.id == destnode.id {
		return nil
	}
	node.Lock()
	defer node.UnLock()
	node.parent.Lock()
	defer node.parent.UnLock()
	destnode.Lock()
	defer destnode.UnLock()
	if nil == destnode.child {
		destnode.child = utypes.NewSafeMap()
	}
	if err := destnode.child.PutX(newname, node); nil != err {
		return errors.New("dest is exists, name: " + newname)
	} else {
		node.parent.child.Delete(node.name)
		// 需要避免同父级目录锁定
		node.parent.updateDirSize(int64(-node.size), destnode.id)
		destnode.updateDirSize(int64(node.size), node.parent.id)
	}
	node.name = newname
	node.parent = destnode
	return nil
}

// Delete 删除节点
func (node *TMNode) Delete() error {
	node.parent.Lock()
	defer node.parent.UnLock()
	node.parent.child.Delete(node.name)
	node.parent.updateDirSize(int64(-node.size), "")
	return nil
}

// DeleteChild 删除子节点
func (node *TMNode) DeleteChild(name string) error {
	if node.flag == FLAG_DIR {
		node.Lock()
		defer node.UnLock()
		if nil != node.child && node.child.Size() > 0 && node.child.ContainsKey(name) {
			node.child.Delete(name)
			node.updateDirSize(int64(-node.size), "")
		}
		return nil
	}
	return errors.New("node is not dir")
}

// updateDirSize 更新文件夹大小, 操作当前节点无锁, 父级节点有锁 size=+oo|-oo
func (node *TMNode) updateDirSize(size int64, ignore string) {
	if size != 0 {
		node.size += int64(size)
		pnode := node.parent
		for pnode != nil {
			// 对于父级目录, 只要改完数据就解锁
			canlock := len(ignore) == 0 || pnode.id != ignore
			if canlock {
				pnode.Lock()
			}
			pnode.size += int64(size)
			if canlock {
				pnode.UnLock()
			}
			pnode = pnode.parent
		}
	}
}

// Clone2TNode 拷贝对象, 生成一个新的公开对象, 不包含子节点
func (node *TMNode) Clone2TNode() *ifilestorage.TNode {
	node.RLock()
	defer node.RUnLock()
	tnode := ifilestorage.TNode{
		Id:    node.id,
		Addr:  node.addr,
		Flag:  node.flag,
		Name:  node.name,
		Size:  node.size,
		Ctime: node.ctime,
		Mtime: node.mtime,
		Props: node.props,
	}
	if nil != node.parent {
		tnode.Pid = node.parent.id
	}
	return &tnode
}

// ReverseNode4New 从新建对象赋值到实例对象
func (node *TMNode) ReverseNode4New(pnode *TMNode, nodedto ifilestorage.TNode4New) *TMNode {
	node.parent = pnode
	node.id = nodedto.Id
	node.addr = nodedto.Addr
	node.flag = nodedto.Flag
	node.name = nodedto.Name
	node.size = nodedto.Size
	node.ctime = nodedto.Ctime
	node.mtime = nodedto.Mtime
	node.props = nodedto.Props
	node.lock = new(sync.RWMutex)
	return node
}
