// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 树

package filenames

import (
	"database/sql"
	"errors"
	"namenode/biz/modules/filenames/names"
	"namenode/ifilestorage"
	"pakku/ipakku"
	"pakku/utils/logs"
	"pakku/utils/strutil"
	"pakku/utils/upool"
	"time"
)

// NameManage 树
type NameManage struct {
	image *names.NameImage
	story *names.NameStory
	fw    *upool.GoWorker
	ev    ipakku.AppSyncEvent
}

// Setting 设置DB, 文件和目录可以分开存
func (nm *NameManage) Setting(st ifilestorage.DBSetting) (err error) {
	if nil == nm.story {
		nm.story = &names.NameStory{}
		err = nm.story.Initial("filenames", st)
	}
	if nil == nm.image {
		nm.image = names.NewNameImage()
	}
	if nil == nm.fw {
		nm.fw = upool.NewGoWorker(50, 100)
		go func() {
			nm.fw.WaitGoWorkerClose()
		}()
	}
	return err
}

// LoadCache 初始化缓存
func (nm *NameManage) LoadCache() (err error) {
	root, err := nm.story.GetRoot(nm.story.GetDB(), AREA_ROOT_MAIN)
	if nil != err {
		return err
	}
	if nil != root {
		if nil == nm.image.GetRoot() {
			root.Pid = ""
			err = nm.image.AddNode(ifilestorage.TNode4New{}.ReverseTNode(*root))
		}
		if nil == err {
			err = nm.story.WalkNodeAllChilds(nm.story.GetDB(), root.Id, func(node *ifilestorage.TNode) error {
				return nm.image.AddNode(ifilestorage.TNode4New{}.ReverseTNode(*node))
			})
		}
	}
	return err
}

// AddNode 插入节点
func (nm *NameManage) AddNode(node ifilestorage.TNode4New) (res string, err error) {
	node.Id = strutil.GetUUID()
	if err = node.CheckNodeNotNull(node.Name == "/"); nil == err {
		var conn *sql.Tx
		conn, err = nm.story.GetSqlTx()
		if nil == err {
			if err = nm.story.InsertNode(conn, node); nil == err {
				if err = nm.image.AddNode(node); nil == err {
					if err = conn.Commit(); nil == err {
						res = node.Id
					}
				}
			}
			if nil != err {
				if err1 := conn.Rollback(); nil != err1 {
					err = err1
				}
			}
		}
	}
	return res, err
}

// GetNodeByPath 获得节点 根据路径
func (nm *NameManage) GetNodeByPath(path string) (res *ifilestorage.TNode) {
	if len(path) == 0 {
		return nm.image.GetRoot()
	}
	res = nm.image.GetNodeByPath(path)
	return res
}

// GetNode 获得节点 根据node.id
func (nm *NameManage) GetNode(id string) (res *ifilestorage.TNode) {
	return nm.image.GetNodeByID(id)
}

// ListNodeChilds 获得子节点 根据node.id
func (nm *NameManage) ListNodeChilds(id string, limit int, offset int) (nodes []ifilestorage.TNode, count int) {
	return nm.image.GetNodeChildsByID(id, limit, offset)
}

// DelNode 删除节点
func (nm *NameManage) DelNode(id string) (err error) {
	cnode := nm.GetNode(id)
	if nil == cnode {
		return errors.New("node not find")
	}
	var conn *sql.Tx
	conn, err = nm.story.GetSqlTx()
	if nil == err {
		// 移动到待删除区域, 后续异步删除
		if err = nm.story.MoveNode(conn, id, cnode.Name, AREA_ROOT_RECYCLE); nil == err {
			if err = nm.image.RemoveNode(cnode.Id); nil == err {
				err = conn.Commit()
			}
		}
		if nil != err {
			conn.Rollback()
		}
	}
	return err
}

// UpdateNode 更新节点
func (nm *NameManage) UpdateNode(id string, node ifilestorage.TNode4Update) (err error) {
	var conn *sql.Tx
	conn, err = nm.story.GetSqlTx()
	if nil != err {
		return err
	}
	if err = nm.story.UpdateById(conn, id, node); nil == err {
		if err = nm.image.UpdateNode(id, node); nil == err {
			err = conn.Commit()
		}
	}
	if nil != err {
		conn.Rollback()
	}
	return err
}

// CopyNode 复制节点, srcid -> dstid.child
func (nm *NameManage) CopyNode(srcid, name, destid string) (id string, err error) {
	var node *ifilestorage.TNode
	if node = nm.image.GetNodeByID(srcid); nil == node {
		return "", errors.New("src node not exist")
	}
	node.Id = ""
	node.Name = name
	node.Pid = destid
	node.Mtime = time.Now().UnixMilli()
	node.Id, err = nm.AddNode(ifilestorage.TNode4New{}.ReverseTNode(*node))
	return node.Id, err
}

// MoveNode 移动节点, srcid -> dstid.child
func (nm *NameManage) MoveNode(srcid, name, destid string) (err error) {
	var conn *sql.Tx
	conn, err = nm.story.GetSqlTx()
	if nil == err {
		if err = nm.story.MoveNode(conn, srcid, name, destid); nil == err {
			if err = nm.image.MoveNode(srcid, name, destid); nil == err {
				err = conn.Commit()
			}
		}
		if nil != err {
			conn.Rollback()
		}
	}
	return err
}

// NodeCount 统计Node
func (nm *NameManage) NodeCount() int64 {
	return nm.image.NodeCount()
}

// Install 第一次初始化安装表
func (nm *NameManage) Install() (err error) {
	conn, err := nm.story.GetSqlTx()
	if nil != err {
		return err
	}
	if err = nm.story.CreateTables(conn); nil == err {
		if err = nm.initialStoryRoot(conn); nil == err {
			err = conn.Commit()
		}
	}
	if nil != err {
		conn.Rollback()
	}
	return err
}

// 扫描需要清理的数据, 并执行
func (nm *NameManage) StartClearRunnable() {
	nm.waitClearTime()
	nodes, err := nm.story.GetNodeChilds(nm.story.GetDB(), AREA_ROOT_RECYCLE, 200)
	if nil != err {
		logs.Panicln(err)
	}
	logs.Infof("扫描到一级待删除数据 %d 条\r\n", len(nodes))
	if len(nodes) > 0 {
		for i := 0; i < len(nodes); i++ {
			var err error
			if nodes[i].Flag == names.FLAG_FILE {
				err = nm.doDelNode(*nodes[i])
			} else {
				err = nm.doDeleteDir(nodes[i])
			}
			if nil != err {
				logs.Errorln(err)
			}
		}
		nm.StartClearRunnable()
	} else {
		time.Sleep(time.Minute * 30)
	}
	go nm.StartClearRunnable()
}

// waitClearTime 30秒内连续创建x个以上则认为当前服务忙, 不适合删除操作
func (nm *NameManage) waitClearTime() {
	t := int64(30)
	max := int64(60) // 每秒 2 个算
	start := time.Now().Unix()
	if count, err := nm.story.CountInTime(nm.story.GetDB(), start-t, start); nil != err {
		logs.Errorln(err)
		time.Sleep(time.Minute)
	} else {
		if count > max {
			logs.Debugf("名称服务当前忙, 等待可用清理窗口, %ds / %d\r\n", t, count)
			time.Sleep(time.Second * 15)
		} else {
			return
		}
	}
	nm.waitClearTime()
}

// doDelNode 物理删除数据
func (nm *NameManage) doDelNode(node ifilestorage.TNode) (err error) {
	// 删除datanode数据
	if node.Flag == names.FLAG_FILE && len(node.Addr) > 0 {
		if err = nm.ev.PublishSyncEvent("filenames", "deletenode", node); nil != err {
			return err
		}
	}
	// 物理删除
	var conn *sql.Tx
	if conn, err = nm.story.GetSqlTx(); nil == err {
		if err = nm.story.DeleteByID(conn, node.Id); nil == err {
			err = conn.Commit()
		}
		if nil != err {
			conn.Rollback()
		}
	}
	logs.Debugf("delete id=%s flag=%d  \r\n", node.Id, node.Flag)
	return err
}
func (nm *NameManage) doDeleteDir(node *ifilestorage.TNode) (err error) {
	nm.waitClearTime()
	var nodes []*ifilestorage.TNode
	if nodes, err = nm.story.GetNodeChilds(nm.story.GetDB(), node.Id, 200); nil == err {
		if len(nodes) > 0 {
			dir := make([]*ifilestorage.TNode, 0)
			for i := 0; i < len(nodes); i++ {
				if nodes[i].Flag == names.FLAG_FILE {
					if err = nm.doDelNode(*nodes[i]); nil != err {
						return err
					}
				} else {
					dir = append(dir, nodes[i])
				}
			}
			if len(dir) > 0 {
				for i := 0; i < len(dir); i++ {
					if err = nm.doDeleteDir(dir[i]); nil != err {
						return err
					}
				}
			}
			//
			err = nm.doDeleteDir(node)
		} else {
			err = nm.doDelNode(*node)
		}
	}
	return err
}

// initialStoryRoot 初始化根结构
func (nm *NameManage) initialStoryRoot(conn *sql.Tx) (err error) {
	err = nm.story.InsertNode(conn, ifilestorage.TNode4New{Id: strutil.GetUUID(), Pid: AREA_ROOT_MAIN, Name: "/", Flag: names.FLAG_DIR, Ctime: time.Now().UnixMilli(), Mtime: time.Now().UnixMilli()})
	return err
}
