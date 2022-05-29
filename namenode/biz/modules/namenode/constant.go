// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package namenode

import (
	"errors"
	"strings"
)

const (
	FLAG_DIR  = 0
	FLAG_FILE = 1
)

// ErrInvalidToken 无效的token
var ErrInvalidToken = errors.New("invalid token")

// ErrDataNodeNotAlive datanode 没有启动
var ErrDataNodeNotAlive = errors.New("datanode not alive")

// GetNodeNoByAddr 截取datanode文件地址中的datanode编号 DN101@xxxxxxxx => DN101
func GetNodeNoByAddr(addr string) string {
	if len(addr) > 0 {
		if array := strings.Split(addr, "@"); len(array) > 1 {
			return array[0]
		}
	}
	return ""
}

// GetFIDByAddr 截取datanode文件地址中的datanode文件id DN101@xxxxxxxx => xxxxxxxx
func GetFIDByAddr(addr string) string {
	if len(addr) > 0 {
		if array := strings.Split(addr, "@"); len(array) > 1 {
			return array[1]
		}
	}
	return ""
}

// BuildDataNodeAddr 生产datanode文件地址
func BuildDataNodeAddr(nodeno, fid string) string {
	return nodeno + "@" + fid
}
