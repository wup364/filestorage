// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 树内存映射

package names

import (
	"errors"
	"namenode/ifilestorage"
	"strings"

	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
	"github.com/wup364/pakku/utils/utypes"
)

// NewNameImage NewNameImage
func NewNameImage() *NameImage {
	return &NameImage{
		allIndex: utypes.NewSafeMap(),
	}
}

// NameImage 树内存映射
type NameImage struct {
	rootId   string          // 根节点ID
	allIndex *utypes.SafeMap // 所有ID的索引, 方便快速ID检索
}

// GetRoot 根节点
func (nm *NameImage) GetRoot() *ifilestorage.TNode {
	return nm.GetNodeByID(nm.rootId)
}

// NodeCount 统计有多少个节点
func (nm *NameImage) NodeCount() int64 {
	return int64(nm.allIndex.Size())
}

// GetFileSize 统计文件|文件夹有多少个大(同步阻塞) 50W/0.7s
func (nm *NameImage) GetFileSize(id string) int64 {
	if node := nm.doGetTMNode(id); nil != node {
		return node.GetFileSize()
	}
	return 0
}

// GetNodeByID 根据ID查找node
func (nm *NameImage) GetNodeByID(id string) *ifilestorage.TNode {
	if len(id) > 0 {
		if node := nm.doGetTMNode(id); nil != node {
			return node.Clone2TNode()
		}
	}
	return nil
}

// GetNodeByPath 根据Path查找node
func (nm *NameImage) GetNodeByPath(path string) *ifilestorage.TNode {
	if len(path) == 0 {
		path = "/"
	}
	var root *TMNode
	if root = nm.doGetTMNode(nm.rootId); nil == root {
		return nil
	}
	if path = strutil.Parse2UnixPath(path); path == "/" {
		return root.Clone2TNode()
	}
	var names []string
	if names = strings.Split(path, "/"); len(names) < 2 {
		return nil
	}
	var res *TMNode
	nextNode := root.child
	for i := 1; i < len(names); i++ {
		if nil == nextNode {
			return nil
		}
		if node, ok := nextNode.Get(names[i]); !ok || nil == node {
			return nil
		} else {
			res = node.(*TMNode)
			nextNode = res.child
		}
	}
	return res.Clone2TNode()
}

// GetNodeChildsByID 根据ID查找子节点
func (nm *NameImage) GetNodeChildsByID(id string, limit int, offset int) ([]ifilestorage.TNode, int) {
	if len(id) > 0 {
		if node := nm.doGetTMNode(id); nil != node && nil != node.child {
			if childs, count := node.GetChilds(limit, offset); count > 0 {
				nodes := make([]ifilestorage.TNode, len(childs))
				for i := 0; i < len(childs); i++ {
					nodes[i] = *childs[i].Clone2TNode()
				}
				return nodes, count
			}
		}
	}
	return make([]ifilestorage.TNode, 0), 0
}

// AddNode 插入节点, 文件|文件夹
func (nm *NameImage) AddNode(node ifilestorage.TNode4New) error {
	// 根节点, pid是空&是第一次
	if len(node.Pid) == 0 && len(nm.rootId) == 0 {
		return nm.doCreateRootNode(node)
	}
	//
	if err := node.CheckNodeNotNull(false); nil != err {
		return err
	}
	if nm.allIndex.ContainsKey(node.Id) {
		return errors.New("node.Id is exist, id: " + node.Id)
	}
	// 获取父节点
	var parent *TMNode
	if parent = nm.doGetTMNode(node.Pid); nil == parent {
		return errors.New("node.Pid is not exist, Pid: " + node.Pid)
	}
	if child, err := parent.AddChild(node); nil != err {
		return err
	} else {
		if err := nm.allIndex.PutX(node.Id, child); nil != err {
			if err := parent.DeleteChild(node.Name); nil != err {
				logs.Panicln(err)
			}
			return errors.New("node.Id is exist, id: " + node.Id)
		}
	}
	return nil
}

// UpdateNode 更改某个节点信息, 文件|文件夹
func (nm *NameImage) UpdateNode(id string, node ifilestorage.TNode4Update) error {
	if len(id) > 0 {
		if cnode := nm.doGetTMNode(id); nil != cnode {
			return cnode.Update(node)
		}
	}
	return errors.New("nodeId is not exist, id: " + id)
}

// MoveNode 移动节点, 文件|文件夹 id -> pid.child
func (nm *NameImage) MoveNode(id, name, pid string) error {
	if cnode := nm.doGetTMNode(id); nil == cnode {
		return errors.New("src node not exist")
	} else {
		return cnode.MoveChild(nm.doGetTMNode(pid), name)
	}
}

// RemoveNode 删除节点, 文件|文件夹
func (nm *NameImage) RemoveNode(id string) error {
	if cnode := nm.doGetTMNode(id); nil != cnode {
		if err := cnode.Delete(); nil == err {
			nm.allIndex.Delete(cnode.id)
			return nil
		} else {
			return err
		}
	}
	return errors.New("id is not exist, id: " + id)
}

// doCreateRootNode 创建跟节点
func (nm *NameImage) doCreateRootNode(insert ifilestorage.TNode4New) error {
	if err := insert.CheckNodeNotNull(true); nil != err {
		return err
	}
	root := new(TMNode).ReverseNode4New(nil, insert)
	if err := nm.allIndex.PutX(insert.Id, root); nil != err {
		return err
	}
	nm.rootId = insert.Id
	return nil
}

// doGetTMNode 根据ID查找node, 内存原值
func (nm *NameImage) doGetTMNode(id string) *TMNode {
	if node, ok := nm.allIndex.Get(id); !ok || nil == node {
		return nil
	} else {
		return node.(*TMNode)
	}
}
