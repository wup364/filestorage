// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 文件操作接口

package ifiledatas

import "io"

// FileDriver 文件管理接口
type FileDriver interface {
	GetDriverType() string
	// 状态判断
	IsDir(src string) bool
	IsFile(src string) bool
	IsExist(src string) bool
	// 信息读取
	GetNode(src string) *Node
	GetNodes(src []string, ignoreNotIsExist bool) ([]Node, error)
	GetFileSize(src string) int64
	GetModifyTime(src string) int64
	GetDirList(src string, limit, offset int) []string
	GetDirNodeList(src string, limit, offset int) ([]Node, error)
	// 目录操作
	DoMkDir(path string) error
	DoDelete(src string) error
	DoRename(src string, dst string) error
	DoCopy(src, dst string, replace bool) error
	DoMove(src, dst string, repalce bool) error
	// 流操作
	DoWrite(src string, ioReader io.Reader) error
	DoRead(src string, offset int64) (io.ReadCloser, error)
}

// Node 文件|夹基础属性
type Node struct {
	Path   string
	Mtime  int64
	IsFile bool
	IsDir  bool
	Size   int64
}
