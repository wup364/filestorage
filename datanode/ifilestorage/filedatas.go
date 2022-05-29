// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 文件操作接口

package ifilestorage

import "io"

// FNode 文件|夹基础属性(filedatas)
type FNode struct {
	Path   string
	Mtime  int64
	IsFile bool
	IsDir  bool
	Size   int64
}

// FileDatas 文件数据块管理模块
type FileDatas interface {
	IsDir(src string) bool
	IsFile(src string) bool
	IsExist(src string) bool

	GetNode(src string) *FNode
	GetNodes(src []string, ignoreNotIsExist bool) ([]FNode, error)
	GetFileSize(src string) int64
	GetDirList(src string, limit, offset int) []string
	GetDirNodeList(src string, limit, offset int) ([]FNode, error)

	DoMkDir(src string) error
	DoDelete(src string) error
	DoRename(src string, dst string) error
	DoCopy(src, dst string, replace bool) error
	DoMove(src, dst string, replace bool) error

	DoWrite(src string, ioReader io.Reader) error
	DoRead(src string, offset int64) (io.ReadCloser, error)
}
