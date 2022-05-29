// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 命名空间管理

package namenode

import (
	"database/sql"
	"errors"
	"namenode/ifilestorage"
	"pakku/ipakku"
	"pakku/utils/fileutil"
	"pakku/utils/logs"
	"pakku/utils/strutil"
	"strings"
	"time"
)

const (
	CacheLib_StreamToken          = "namenode.streamtoken"
	CacheLib_StreamTokenExpSecond = 60 * 15 // Token过期时间
)

// NameNodeImpl 命名空间管理
type NameNodeImpl struct {
	c  ipakku.AppCache        `@autowired:"AppCache"`
	e  ipakku.AppSyncEvent    `@autowired:"AppEvent"`
	fn ifilestorage.FileNames `@autowired:"FileNames"`
	d  ifilestorage.Conn4DataNodes
	da *Addr4Deleted
}

// DoAskWriteToken DoAskWriteToken
func (n *NameNodeImpl) DoAskWriteToken(src string) (*ifilestorage.StreamToken, error) {
	dnode, err := n.d.GetRandomDNode()
	if nil != err {
		return nil, err
	}
	token := strutil.GetUUID()
	st := &ifilestorage.StreamToken{
		FilePath: src,
		FileID:   token,
		Token:    token,
		CTime:    time.Now().UnixMilli(),
		MTime:    time.Now().UnixMilli(),
		NodeNo:   dnode.GetNodeNo(),
		EndPoint: dnode.GetStreamEndpoint(),
		Type:     ifilestorage.StreamTokenType_Write,
	}
	if err := n.c.Set(CacheLib_StreamToken, token, st); nil == err {
		return st, nil
	} else {
		return nil, err
	}
}

// DoAskReadToken DoAskReadToken
func (n *NameNodeImpl) DoAskReadToken(src string) (*ifilestorage.StreamToken, error) {
	var node *ifilestorage.TNode
	if node = n.GetNode(src); nil == node || node.Flag != FLAG_FILE {
		return nil, fileutil.PathNotExist("askReadToken", src)
	}
	nodeNo := GetNodeNoByAddr(node.Addr)
	if dnode, err := n.d.GetDNode(nodeNo); nil != err {
		return nil, err
	} else {
		token := strutil.GetUUID()
		st := &ifilestorage.StreamToken{
			NodeNo:   nodeNo,
			Token:    token,
			FilePath: src,
			FileSize: node.Size,
			CTime:    time.Now().UnixMilli(),
			MTime:    time.Now().UnixMilli(),
			FileID:   GetFIDByAddr(node.Addr),
			EndPoint: dnode.GetStreamEndpoint(),
			Type:     ifilestorage.StreamTokenType_Read,
		}
		if err := n.c.Set(CacheLib_StreamToken, token, st); nil == err {
			return st, nil
		} else {
			return nil, err
		}
	}
}

// DoQueryToken 查出token
func (n *NameNodeImpl) DoQueryToken(token string) (*ifilestorage.StreamToken, error) {
	var val ifilestorage.StreamToken
	if err := n.c.Get(CacheLib_StreamToken, token, &val); nil == err {
		return &val, nil
	} else if err == ipakku.ErrNoCacheHit {
		return nil, ErrInvalidToken
	} else {
		return nil, err
	}

}

// DoRefreshToken DoRefreshToken
func (n *NameNodeImpl) DoRefreshToken(token string) (st *ifilestorage.StreamToken, err error) {
	if st, err = n.DoQueryToken(token); nil == err {
		st.MTime = time.Now().UnixMilli()
		if err = n.c.Set(CacheLib_StreamToken, token, st); nil != err {
			logs.Errorln(err)
		}
	} else {
		return nil, err
	}
	return n.DoQueryToken(token)
}

// DoSubmitWriteToken 归档数据
func (n *NameNodeImpl) DoSubmitWriteToken(token string, override bool) (node *ifilestorage.TNode, err error) {
	var st *ifilestorage.StreamToken
	if st, err = n.DoQueryToken(token); nil == err && st != nil {
		path := st.FilePath
		isExist := n.IsExist(path)
		if !override && isExist {
			// 是否需要覆盖
			return nil, fileutil.PathExist("submitToken", path)
		}
		// RPC datanode 递交数据, 返回文件ID
		var dnc ifilestorage.RPC4DataNode
		if dnc, err = n.d.GetDNode(st.NodeNo); nil != err {
			return nil, err
		}
		var dnode *ifilestorage.DNode
		if dnode, err = dnc.DoSubmitWriteToken(token); nil == err {
			// 保存&更新元数据信息
			if isExist {
				err = n.DoUpdateFile(path, BuildDataNodeAddr(st.NodeNo, dnode.Id), dnode.Size)
			} else {
				_, err = n.DoMkFile(path, BuildDataNodeAddr(st.NodeNo, dnode.Id), dnode.Size)
			}
			if nil == err {
				node = n.GetNode(path)
			}
		}
	} /* else {
		err = ErrInvalidToken
	}*/
	return node, err
}

// DoRename DoRename
func (n *NameNodeImpl) DoRename(src, newName string) error {
	src = strutil.Parse2UnixPath(src)
	var cnode *ifilestorage.TNode
	if cnode = n.fn.GetNodeByPath(src); nil == cnode {
		return fileutil.PathNotExist("rename", src)
	}
	newPath := strutil.GetPathParent(src) + "/" + newName
	if nnode := n.fn.GetNodeByPath(newPath); nil != nnode {
		return fileutil.PathExist("rename", newPath)
	}
	cnode.Name = newName
	return n.fn.UpdateNode(cnode.Id, ifilestorage.TNode4Update{}.ReverseTNode(*cnode))
}

// DoMkFile DoMkFile
func (n *NameNodeImpl) DoMkFile(src, addr string, size int64) (newid string, err error) {
	if src = strutil.Parse2UnixPath(src); len(src) == 0 {
		return newid, errors.New("path is empty")
	}
	if src == "/" {
		return newid, errors.New("no access path '/'")
	}
	if err = n.checkPathSafety(src); nil != err {
		return newid, err
	}
	var pid string
	psrc := strutil.GetPathParent(src)
	if pnode := n.GetNode(psrc); nil == pnode {
		if pid, err = n.DoMkDir(psrc); nil != err {
			return newid, err
		}
	} else {
		pid = pnode.Id
	}
	if len(pid) == 0 {
		return newid, fileutil.PathNotExist("mkfile", psrc)
	}
	newid, err = n.fn.AddNode(ifilestorage.TNode4New{
		Addr:  addr,
		Pid:   pid,
		Name:  strutil.GetPathName(src),
		Flag:  FLAG_FILE,
		Size:  size,
		Ctime: time.Now().UnixMilli(),
		Mtime: time.Now().UnixMilli(),
	})
	return newid, err
}

// DoUpdateFile DoUpdateFile
func (n *NameNodeImpl) DoUpdateFile(src, addr string, size int64) (err error) {
	src = strutil.Parse2UnixPath(src)
	if len(src) == 0 {
		return errors.New("path is empty")
	}
	if src == "/" {
		return errors.New("no access path '/'")
	}
	if node := n.GetNode(src); nil != node {
		var conn *sql.Tx
		if conn, err = n.da.GetSqlTx(); nil == err {
			if err = n.da.InsertItem(conn, node.Addr); nil == err {
				node.Addr = addr
				node.Size = size
				node.Mtime = time.Now().UnixMilli()
				if err = n.fn.UpdateNode(node.Id, ifilestorage.TNode4Update{}.ReverseTNode(*node)); nil == err {
					err = conn.Commit()
				} else {
					conn.Rollback()
				}
			}
		}
	} else {
		err = fileutil.PathNotExist("updatefile", src)
	}
	return err
}

// DoMkDir DoMkDir
func (n *NameNodeImpl) DoMkDir(src string) (newid string, err error) {
	if src = strutil.Parse2UnixPath(src); len(src) == 0 {
		return newid, errors.New("path is empty")
	}
	if src == "/" {
		return newid, errors.New("no access path '/'")
	}
	if err = n.checkPathSafety(src); nil != err {
		return newid, err
	}
	temp := src
	pid := ""
	makedirs := make([]string, 0)
	for len(temp) > 0 {
		node := n.fn.GetNodeByPath(temp)
		if nil == node {
			makedirs = append(makedirs, strutil.GetPathName(temp))
			temp = strutil.GetPathParent(temp)
			if len(temp) == 0 || temp == "/" {
				rnode := n.fn.GetNodeByPath("/")
				if nil == rnode {
					return newid, fileutil.PathNotExist("mkdir", "/")
				}
				pid = rnode.Id
				break
			}
		} else {
			pid = node.Id
			break
		}
	}
	if len(pid) == 0 {
		return newid, errors.New("make dir failed")
	}
	if len(makedirs) == 0 {
		return newid, errors.New("path is exist, path: " + src)
	}
	time := time.Now().UnixMilli()
	for i := len(makedirs) - 1; i > -1; i-- {
		pid, err = n.fn.AddNode(ifilestorage.TNode4New{
			Pid:   pid,
			Name:  makedirs[i],
			Flag:  FLAG_DIR,
			Ctime: time,
			Mtime: time,
		})
		if nil != err {
			break
		}
	}
	return pid, err
}

// DoDelete 删除文件|文件夹(包含子列表)
func (n *NameNodeImpl) DoDelete(src string) error {
	node := n.GetNode(src)
	if nil == node {
		return fileutil.PathNotExist("delete", src)
	}
	return n.fn.DelNode(node.Id)
}

// DoMove 移动文件|文件夹(包含子列表)到目标文件夹下
func (n *NameNodeImpl) DoMove(src, dst string, override bool) error {
	// 同级目录-重命名
	pdst := strutil.GetPathParent(dst)
	if strutil.GetPathParent(src) == pdst {
		return n.DoRename(src, strutil.GetPathName(dst))
	}
	//
	var srcNode *ifilestorage.TNode
	if srcNode = n.GetNode(src); nil == srcNode {
		return fileutil.PathNotExist("move", src)
	}
	if src == dst {
		return fileutil.PathExist("move", dst)
	}
	var pid string
	if pdnode := n.GetNode(pdst); nil == pdnode {
		if newid, err := n.DoMkDir(pdst); nil != err {
			return fileutil.PathNotExist("move", pdst)
		} else {
			pid = newid
		}
	} else if dstNode := n.GetNode(dst); nil != dstNode {
		if !override || dstNode.Flag == FLAG_DIR {
			return fileutil.PathExist("move", dst)
		}
		var err error
		if err = n.fn.UpdateNode(dstNode.Id, ifilestorage.TNode4Update{
			Addr:  srcNode.Addr,
			Name:  srcNode.Name,
			Size:  srcNode.Size,
			Props: srcNode.Props,
			Mtime: time.Now().UnixMilli(),
		}); nil == err {
			if err = n.fn.DelNode(srcNode.Id); nil != err {
				if n.fn.GetNode(srcNode.Id) != nil {
					return err
				} else {
					err = nil
				}
			}
		}
		return err
	} else {
		pid = pdnode.Id
	}
	return n.fn.MoveNode(srcNode.Id, strutil.GetPathName(dst), pid)
}

// DoCopy 复制文件|夹(不包含子列表)到目标文件夹
func (n *NameNodeImpl) DoCopy(src, dst string, override bool) (string, error) {
	var srcNode *ifilestorage.TNode
	if srcNode = n.GetNode(src); nil == srcNode {
		return "", fileutil.PathNotExist("copy", src)
	}
	if src == dst {
		return "", fileutil.PathExist("copy", dst)
	}
	var pid string
	pdst := strutil.GetPathParent(dst)
	if pdnode := n.GetNode(pdst); nil == pdnode {
		if newid, err := n.DoMkDir(pdst); nil != err {
			return "", fileutil.PathNotExist("copy", pdst)
		} else {
			pid = newid
		}
	} else if dstNode := n.GetNode(dst); nil != dstNode {
		if !override {
			return "", fileutil.PathExist("copy", dst)
		}
		updateVo := ifilestorage.TNode4Update{
			Name:  srcNode.Name,
			Props: srcNode.Props,
			Mtime: time.Now().UnixMilli(),
		}
		if dstNode.Flag == FLAG_FILE {
			updateVo.Addr = srcNode.Addr
			updateVo.Size = srcNode.Size
		}
		return dstNode.Id, n.fn.UpdateNode(dstNode.Id, updateVo)
	} else {
		pid = pdnode.Id
	}
	return n.fn.CopyNode(srcNode.Id, strutil.GetPathName(dst), pid)
}

// IsFile 是否是文件
func (n *NameNodeImpl) IsFile(src string) bool {
	node := n.GetNode(src)
	return nil != node && node.Flag == FLAG_FILE
}

// IsDir 是否是文件夹
func (n *NameNodeImpl) IsDir(src string) bool {
	node := n.GetNode(src)
	return nil != node && node.Flag == FLAG_DIR
}

// IsExist 是否存在
func (n *NameNodeImpl) IsExist(src string) bool {
	node := n.GetNode(src)
	return nil != node
}

// GetFileSize 获取文件大小
func (n *NameNodeImpl) GetFileSize(src string) int64 {
	if node := n.GetNode(src); nil != node {
		return node.Size
	}
	return 0
}

// GetDirNameList 获取文件夹列表
func (n *NameNodeImpl) GetDirNameList(src string, limit int, offst int) ([]string, int) {
	if node := n.GetNode(src); nil != node {
		if nodes, count := n.fn.ListNodeChilds(node.Id, limit, offst); nil != nodes {
			res := make([]string, len(nodes))
			for i := 0; i < len(nodes); i++ {
				res[i] = nodes[i].Name
			}
			return res, count
		}
	}
	return nil, 0
}

// GetNode 获取node
func (n *NameNodeImpl) GetNode(src string) *ifilestorage.TNode {
	return n.fn.GetNodeByPath(strutil.Parse2UnixPath(src))
}

// GetNode 获取node
func (n *NameNodeImpl) GetNodeById(id string) *ifilestorage.TNode {
	return n.fn.GetNode(id)
}

// GetDirNodeList 获取node的子节点信息
func (n *NameNodeImpl) GetDirNodeList(src string, limit int, offst int) ([]ifilestorage.TNode, int) {
	if node := n.GetNode(src); nil != node {
		return n.fn.ListNodeChilds(node.Id, limit, offst)
	}
	return nil, 0
}

// DoQueryDeletedDataAddr 根据nodeno查找, 只能查找一天前的
func (n *NameNodeImpl) DoQueryDeletedDataAddr(nodeno string, limit int) ([]ifilestorage.DeletedAddr, error) {
	ctime := time.Now().UnixMilli() - int64(time.Hour*24)
	return n.da.FindByNodeNo(n.da.GetDB(), nodeno, ctime, limit)
}

// DoConfirmDeletedDataAddr 根据nodeno查找
func (n *NameNodeImpl) DoConfirmDeletedDataAddr(nodenos []string) (err error) {
	var tx *sql.Tx
	if tx, err = n.da.GetSqlTx(); nil == err {
		if err = n.da.DeleteInID(tx, nodenos); nil == err {
			err = tx.Commit()
		} else {
			tx.Rollback()
		}
	}
	return err
}

// checkPathSafety 路径合规检查, 避免 ../ ./之类的路径
func (n *NameNodeImpl) checkPathSafety(path string) error {
	if path == ".." || strings.Contains(path, "../") || strings.Contains(path, "..\\") {
		return errors.New("unsupported path format '../'")
	} else if path == "." || strings.Contains(path, "./") || strings.Contains(path, ".\\") {
		return errors.New("unsupported path format './'")
	}
	return nil
}
