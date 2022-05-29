// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package ifilestorage

// DataNodeConn 数据存储节点信息
type DataNodeConn struct {
	IP        string
	RPCURL    string // 远程调用地址
	StreamURL string // 文件流地址
}

// DeletedAddr 已删除的文件地址
type DeletedAddr struct {
	Id     string
	Ctime  int64
	NodeNo string
	Fid    string
}

// NameNode 命名空间节点
type NameNode interface {
	IsDir(src string) bool
	IsFile(src string) bool
	IsExist(src string) bool
	GetFileSize(src string) int64

	GetNode(src string) *TNode
	GetDirNameList(src string, limit, offset int) ([]string, int)
	GetDirNodeList(src string, limit, offset int) ([]TNode, int)

	DoMkDir(path string) (string, error)
	DoDelete(src string) error
	DoRename(src string, dst string) error
	DoCopy(src, dst string, override bool) (string, error)
	DoMove(src, dst string, override bool) error
	DoMkFile(src, addr string, size int64) (string, error)
	DoUpdateFile(src, addr string, size int64) error

	DoQueryToken(token string) (*StreamToken, error)
	DoAskReadToken(src string) (*StreamToken, error)
	DoAskWriteToken(src string) (*StreamToken, error)
	DoRefreshToken(token string) (*StreamToken, error)
	DoSubmitWriteToken(token string, override bool) (node *TNode, err error)

	DoQueryDeletedDataAddr(nodeno string, limit int) ([]DeletedAddr, error)
	DoConfirmDeletedDataAddr(nodenos []string) error
}
