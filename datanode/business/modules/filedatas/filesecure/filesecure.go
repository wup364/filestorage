// Copyright (C) 2023 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 文件安全

package filesecure

import (
	"crypto/sha512"
	"datanode/business/modules/filedatas/ifiledatas"
	"io"
)

// 常量
const (
	ENCRYPTIONMODE_NONE string = ""    // 加密模式 - 无加密
	ENCRYPTIONMODE_XOR  string = "XOR" // 加密模式 - XOR
)

// NewFileSecure 实例化
func NewFileSecure(pwd, encryptionMode string) ifiledatas.FileSecure {
	fc := &FileSecure{model: encryptionMode}
	if fc.model != ENCRYPTIONMODE_NONE {
		fc.setPasswd(pwd)
	}
	return fc
}

// FileSecure 文件安全模块
type FileSecure struct {
	passwd []byte
	model  string
}

func (fst *FileSecure) setPasswd(passwd string) {
	hx := sha512.New()
	hx.Write([]byte(passwd))
	fst.passwd = hx.Sum(nil)
}

func (fst *FileSecure) EncodeWrapper(ioReader io.Reader) io.Reader {
	if fst.model == ENCRYPTIONMODE_XOR {
		return &Reader4XOR{
			passwd: fst.passwd,
			reader: ioReader,
		}
	}
	return ioReader
}

func (fst *FileSecure) DecodeWrapper(ioReader io.ReadCloser, offset int64) io.ReadCloser {
	if fst.model == ENCRYPTIONMODE_XOR {
		rr := &Reader4XOR{
			passwd: fst.passwd,
			reader: ioReader,
		}
		rr.SetOffset(offset)
		return rr
	}
	return ioReader
}

// Reader4XOR XOR加密模式的reader
type Reader4XOR struct {
	passwd []byte
	buffer []byte
	rindex int64
	reader io.Reader
	closed bool
}

func (r *Reader4XOR) initBuffer(p []byte) {
	if len(r.buffer) == 0 {
		r.buffer = make([]byte, len(p))
	}
}

// SetOffset 设置偏移量
func (r *Reader4XOR) SetOffset(offset int64) {
	if offset > -1 {
		r.rindex = offset
	}
}

// Read Read & XOR
func (r *Reader4XOR) Read(p []byte) (n int, err error) {
	r.initBuffer(p)
	n, err = r.reader.Read(r.buffer)
	if n > 0 {
		lenri := int64(len(r.passwd))
		for i := 0; i < n; i++ {
			p[i] = r.buffer[i] ^ r.passwd[(r.rindex%lenri)]
			r.rindex++
		}
	}
	return n, err
}

// Close 关闭reader
func (r *Reader4XOR) Close() (err error) {
	if rr, ok := r.reader.(io.ReadCloser); ok {
		err = rr.Close()
	}
	if nil == err {
		r.closed = true
		r.reader = nil
	}
	return
}
