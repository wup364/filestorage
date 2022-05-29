// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package ifilestorage

import "errors"

const (
	FLAG_NODETYPE_DIR  int8 = 0
	FLAG_NODETYPE_FILE int8 = 1
)

// TNode 元数据内容(filenames)
type TNode struct {
	Id    string
	Pid   string
	Addr  string
	Flag  int8
	Name  string
	Size  int64
	Ctime int64
	Mtime int64
	Props string
}

// TNode4New 添加新纪录
type TNode4New struct {
	Id    string
	Pid   string
	Addr  string
	Name  string
	Flag  int8
	Size  int64
	Ctime int64
	Mtime int64
	Props string
}

// ReverseTNode 公开对象转内存对象
func (node TNode4New) ReverseTNode(tnode TNode) TNode4New {
	node.Id = tnode.Id
	node.Pid = tnode.Pid
	node.Addr = tnode.Addr
	node.Flag = tnode.Flag
	node.Name = tnode.Name
	node.Size = tnode.Size
	node.Ctime = tnode.Ctime
	node.Mtime = tnode.Mtime
	node.Props = tnode.Props
	return node
}

// CheckNodeNotNull 节点必填校验
func (node *TNode4New) CheckNodeNotNull(isRoot bool) error {
	if len(node.Id) == 0 {
		return errors.New("node.Id is empty")
	}
	if !isRoot && len(node.Pid) == 0 {
		return errors.New("node.Pid is empty")
	}
	if len(node.Name) == 0 {
		return errors.New("node.Name is empty, id: " + node.Id)
	}
	if node.Flag == FLAG_NODETYPE_FILE && len(node.Addr) == 0 {
		return errors.New("node.Addr is empty, id: " + node.Id)
	}
	if isRoot && node.Flag != FLAG_NODETYPE_DIR {
		return errors.New("root node must be a folder")
	}
	return nil
}

// TNode4Update 修改距离
type TNode4Update struct {
	Addr  string
	Name  string
	Size  int64
	Mtime int64
	Props string
}

// ReverseTNode 公开对象转内存对象
func (node TNode4Update) ReverseTNode(tnode TNode) TNode4Update {
	node.Addr = tnode.Addr
	node.Name = tnode.Name
	node.Size = tnode.Size
	node.Mtime = tnode.Mtime
	node.Props = tnode.Props
	return node
}

// FileNames 命名空间管理
type FileNames interface {
	GetNode(id string) *TNode
	AddNode(node TNode4New) (string, error)
	GetNodeByPath(path string) *TNode
	ListNodeChilds(id string, limit int, offst int) ([]TNode, int)
	DelNode(id string) error
	UpdateNode(id string, node TNode4Update) error
	CopyNode(srcid, name, parentid string) (string, error)
	MoveNode(srcid, name, parentid string) error
	NodeCount() int64
}
