// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package opensdk

import (
	"errors"
	"io"
	"net/http"
	"opensdk/utils"
	"os"
	"strconv"
	"time"
)

// NewDataStream NewDataStream
func NewDataStream() *DataStream {
	return &DataStream{
		httpclient: &http.Client{
			Timeout: time.Minute * 5,
		},
		nodes: utils.NewSafeMap(),
	}
}

type doWriteTokenFunc func(nodeNo, token, endpoint string, pieceNumber int, sha256 string, reader io.Reader) (err error)

// DataStream DataStream
type DataStream struct {
	httpclient *http.Client
	nodes      *utils.SafeMap
}

// SetNodeStreamAddr 设置nodeNo的地址映射
func (s *DataStream) SetNodeStreamAddr(nodes map[string]string) {
	s.nodes.Clear()
	for key, val := range nodes {
		s.nodes.Put(key, val)
	}
}

// GetNodeStreamAddr 获取nodeNo的地址映射
func (s *DataStream) GetNodeStreamAddr(nodeNo string, endpoint string) (string, error) {
	if addr, ok := s.nodes.Get(nodeNo); !ok || len(addr.(string)) == 0 {
		if len(endpoint) == 0 {
			return "", errors.New("Cannot resolve the address of nodeNo=" + nodeNo)
		}
		return endpoint, nil
	} else {
		return addr.(string), nil
	}
}

// GetReadStreamURL 获取下载url
func (s *DataStream) GetReadStreamURL(nodeNo, token, endpoint string) (string, error) {
	if addr, err := s.GetNodeStreamAddr(nodeNo, endpoint); nil != err {
		return "", err
	} else {
		return addr + "/stream/read/" + token, nil
	}
}

// GetWriteStreamURL 获取上传url
func (s *DataStream) GetWriteStreamURL(nodeNo, token, endpoint string) (string, error) {
	if addr, err := s.GetNodeStreamAddr(nodeNo, endpoint); nil != err {
		return "", err
	} else {
		return addr + "/stream/put/" + token, nil
	}
}

// DoReadToken /stream/read/tokenxxxx
func (s *DataStream) DoReadToken(nodeNo, token, endpoint string, offset int64) (r io.ReadCloser, err error) {
	var url string
	if url, err = s.GetReadStreamURL(nodeNo, token, endpoint); nil != err {
		return nil, err
	}
	header := map[string]string{"Connection": "Keep-Alive", "Range": "bytes=" + strconv.Itoa(int(offset)) + "-"}
	if response, err := utils.Request4URL(s.httpclient, http.MethodGet, url, nil, header); nil == err {
		if response.StatusCode == http.StatusOK || response.StatusCode == http.StatusPartialContent {
			return response.Body, nil
		} else {
			return nil, errors.New("[" + response.Status + "] " + utils.ReadAsString(response.Body))
		}
	} else {
		return nil, err
	}
}

// DoWriteToken /stream/put/tokenxxxx
func (s *DataStream) DoWriteToken(nodeNo, token, endpoint string, pieceNumber int, sha256 string, reader io.Reader) (err error) {
	var url string
	if url, err = s.GetWriteStreamURL(nodeNo, token, endpoint); nil != err {
		return err
	}
	url = utils.BuildURLWithMap(url, map[string]string{"hash": sha256, "number": strconv.Itoa(pieceNumber)})
	if response, err := utils.DoUploadFile(s.httpclient, http.MethodPut, url, reader, map[string]string{"Connection": "Keep-Alive"}, "", token); nil == err {
		defer response.Body.Close()
		if response.StatusCode == http.StatusOK {
			return nil
		} else {
			return errors.New("[" + response.Status + "] " + utils.ReadAsString(response.Body))
		}
	} else {
		if nil != response {
			defer response.Body.Close()
		}
		return err
	}
}

// FileUploader FileUploader
func FileUploader(path string, token *StreamToken, pieceSize int64, doWriteToken doWriteTokenFunc) error {
	if file, err := os.Open(path); nil != err {
		return err
	} else {
		return TokenWriter(file, token, pieceSize, doWriteToken)
	}
}

// TokenWriter TokenWriter
func TokenWriter(reader io.Reader, token *StreamToken, pieceSize int64, doWriteToken doWriteTokenFunc) (err error) {
	pieceNumber := 1
	limitr := NewLimitedReader(reader, pieceSize)
	for {
		if err = doWriteToken(token.NodeNo, token.Token, token.EndPoint, pieceNumber, "", limitr); nil != err {
			break
		}
		var next bool
		if next, err = limitr.Next(); !next {
			break
		}
		pieceNumber++
	}
	return err
}

// NewLimitedReader NewLimitedReader
func NewLimitedReader(r io.Reader, n int64) *LimitedReader {
	return &LimitedReader{r: r, n: n, l: n}
}

type LimitedReader struct {
	r io.Reader // reader
	t []byte    // test byte
	n int64     // current
	l int64     // limit
}

func (l *LimitedReader) Read(p []byte) (n int, err error) {
	if l.n <= 0 {
		return 0, io.EOF
	}
	if int64(len(p)) > l.n {
		p = p[0:l.n]
	}
	if len(l.t) == 1 {
		p[0] = l.t[0]
		p = p[1:]
	}
	if n, err = l.r.Read(p); err == nil || err == io.EOF {
		if l.t != nil {
			l.t = nil
			n += 1
		}
	}
	l.n -= int64(n)
	return
}

func (l *LimitedReader) Next() (bool, error) {
	if l.n > 0 {
		return false, nil
	}
	l.t = make([]byte, 1)
	if n, err := l.r.Read(l.t); nil != err && err != io.EOF {
		return false, err
		// n>0 nil EOF
	} else if n > 0 {
		l.n = l.l
		return true, nil
	}
	// n==0 nil EOF
	return false, nil
}
