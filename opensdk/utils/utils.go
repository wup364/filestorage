// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"io"
	"math/rand"
	"os"
	"path"
	"strings"
	"time"
)

// HTTPResponse 接口返回格式约束
type HTTPResponse struct {
	Code int    `json:"code"`
	Flag string `json:"flag"`
	Data string `json:"data"`
}

// Parse2HTTPResponse json转对象
func Parse2HTTPResponse(str string) *HTTPResponse {
	res := &HTTPResponse{}
	if err := json.Unmarshal([]byte(str), res); nil != err {
		return nil
	}
	return res
}

// ReadAsString 从io.Reader读取文字
func ReadAsString(src io.Reader) string {
	if nil == src {
		return ""
	}
	buf := make([]byte, 0)
	for {
		buftemp := make([]byte, 1024)
		nr, er := src.Read(buftemp)
		if nr > 0 {
			buf = append(buf, buftemp[:nr]...)
		}
		if er != nil {
			if er != io.EOF {
				return ""
			}
			break
		}
	}
	return string(buf)
}

// GetStringSHA256 字符转sha256
func GetStringSHA256(str string) string {
	hx := sha256.New()
	hx.Write([]byte(str))
	return hex.EncodeToString(hx.Sum(nil))
}

// GetFileSHA256 获取sha256
func GetFileSHA256(r io.Reader) (string, error) {
	h := sha256.New()
	buf := make([]byte, 1<<20)
	for {
		n, err := io.ReadFull(r, buf)
		if err == nil || err == io.ErrUnexpectedEOF || err == io.EOF {
			if n > 0 {
				if _, err = h.Write(buf[0:n]); err != nil {
					return "", err
				}
			} else {
				break
			}
		} else {
			return "", err
		}
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

// Parse2UnixPath 删除路径后面 /, 把\转换为/
func Parse2UnixPath(str string) string {
	if len(str) == 0 {
		return ""
	}
	return path.Clean(strings.Replace(str, "\\", "/", -1))
}

// GetPathName 截取最后一个'/'后的文字
func GetPathName(path string) string {
	if path == "/" || path == "\\" {
		return ""
	}
	if strings.Contains(path, "\\") {
		return path[strings.LastIndex(path, "\\")+1:]
	}
	return path[strings.LastIndex(path, "/")+1:]
}

// GetPathParent 截取最后一个'/'前的文字
func GetPathParent(path string) string {
	if path == "/" || path == "\\" {
		return ""
	}
	if strings.Contains(path, "\\") {
		return path[:strings.LastIndex(path, "\\")]
	}
	return path[:strings.LastIndex(path, "/")]
}

// getFileInfo 获取文件信息对象
func getFileInfo(path string) (os.FileInfo, error) {
	return os.Stat(path)
}

// OpenFile 获取文件信息对象
func OpenFile(path string) (*os.File, error) {
	return os.Open(path)
}

// GetWriter 获取只写文件对象
// O_RDONLY: 只读模式(read-only)
// O_WRONLY: 只写模式(write-only)
// O_RDWR: 读写模式(read-write)
// O_APPEND: 追加模式(append)
// O_CREATE: 文件不存在就创建(create a new file if none exists.)
// O_EXCL: 与 O_CREATE 一起用, 构成一个新建文件的功能, 它要求文件必须不存在
// O_SYNC: 同步方式打开，即不使用缓存，直接写入硬盘
// O_TRUNC: 打开并清空文件
func GetWriter(path string) (*os.File, error) {
	return os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
}

// IsDir 是否是文件夹
func IsDir(path string) bool {
	stat, err := getFileInfo(path)
	if err == nil {
		return stat.IsDir()
	}
	return false
}

// MkdirAll 创建文件夹-多级
func MkdirAll(path string) error {
	return os.MkdirAll(path, os.ModePerm)
}

// GetDirList 获取一级子目录名字(包含文件|文件夹,无序)
func GetDirList(path string) ([]string, error) {
	f, err := OpenFile(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	list, err := f.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	return list, nil
}

// GetRandom 生成随机字符
func GetRandom(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyz"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
