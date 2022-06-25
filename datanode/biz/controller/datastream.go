// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// datanode 文件流http服务

package controller

import (
	"datanode/biz/bizutils"
	"datanode/ifilestorage"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/serviceutil"
	"github.com/wup364/pakku/utils/strutil"
)

const (
	// HEADER4FILENAME 用头信息标记Form表单中文件的FormName
	HEADER4FILENAME = "filename"
	// HEADER4FILENAMEDEFT 默认使用这个作为Form表单中文件的FormName
	HEADER4FILENAMEDEFT = "file"
	refreshMilliSecond  = 5 * 60 * 1000 // 每5分钟刷新一下Token
)

// DataStream DataStream
type DataStream struct {
	dp ifilestorage.DataNode `@autowired:"DataNode"`
}

// AsController 实现 AsController 接口
func (s *DataStream) AsController() ipakku.ControllerConfig {
	return ipakku.ControllerConfig{
		RequestMapping: "/stream",
		RouterConfig: ipakku.RouterConfig{
			ToLowerCase: true,
			HandlerFunc: [][]interface{}{
				{http.MethodPut, "put/:[A-Za-z0-9]+$", s.Put},
				{http.MethodPost, "put/:[A-Za-z0-9]+$", s.Put},
				{http.MethodGet, "read/:[A-Za-z0-9]+$", s.Read},
				{http.MethodHead, "read/:[A-Za-z0-9]+$", s.ReadHead},
			},
		},
	}
}

// Put 文件上传, 支持Form和Body上传方式
// url格式: /stream/put/tokenxxxxxx?number=1&hash=sha256xxxxxxxxxxxxxx
func (s *DataStream) Put(w http.ResponseWriter, r *http.Request) {
	if method := strings.ToLower(r.Method); method != "put" && method != "post" {
		serviceutil.SendErrorAndStatus(w, http.StatusForbidden, "Invalid method")
		return
	}
	var err error
	token, err := s.dp.DoQueryToken(strutil.GetPathName(r.URL.Path))
	if nil != err {
		serviceutil.SendBadRequest(w, err.Error())
		return
	}
	// 五分钟刷一下
	if time.Now().UnixMilli()-token.MTime >= refreshMilliSecond {
		if _, err = s.dp.DoRefreshToken(token.Token); nil != err {
			serviceutil.SendBadRequest(w, err.Error())
			return
		}
	}
	var urlvals url.Values
	if urlvals, err = url.ParseQuery(r.URL.RawQuery); err != nil {
		serviceutil.SendBadRequest(w, err.Error())
		return
	}
	var pieceNumber int64
	if pieceNumber, err = strconv.ParseInt(urlvals.Get("number"), 10, 32); nil != err {
		serviceutil.SendBadRequest(w, err.Error())
		return
	}
	var fieldname string
	if fieldname = r.Header.Get(HEADER4FILENAME); len(fieldname) == 0 {
		fieldname = HEADER4FILENAMEDEFT
	}
	//
	{
		var multipart *multipart.Reader
		if multipart, err = r.MultipartReader(); nil != err {
			serviceutil.SendBadRequest(w, err.Error())
			return
		}
		hasfile := false
		for {
			p, err := multipart.NextPart()
			if nil == p || err == io.ErrUnexpectedEOF || err == io.EOF {
				break
			}
			if !hasfile && p.FormName() == fieldname {
				hasfile = true
				if err = s.dp.DoWriteToken(token.Token, int(pieceNumber), urlvals.Get("hash"), p); nil != err {
					serviceutil.SendBadRequest(w, err.Error())
				} else {
					p.Close()
					w.WriteHeader(http.StatusOK)
				}
				return
			}
		}
		if !hasfile {
			serviceutil.SendBadRequest(w, "file not found from the form")
		}
	}
}

// Read 下载前的信息获取
// url格式: /stream/read/tokenxxxxxx?name=a.txt
func (s *DataStream) ReadHead(w http.ResponseWriter, r *http.Request) {
	// token
	token, err := s.dp.DoQueryToken(strutil.GetPathName(r.URL.Path))
	if nil != err {
		serviceutil.SendBadRequest(w, err.Error())
		return
	}
	//
	start, end, hasRange := serviceutil.GetRequestRange(r, token.FileSize)
	w.Header().Set("Content-Length", strconv.Itoa(int(end-start)))
	if hasRange {
		w.Header().Set("Content-Range", "bytes "+strconv.Itoa(int(start))+"-"+strconv.Itoa(int(end-1))+"/"+strconv.Itoa(int(token.FileSize)))
	}
	w.WriteHeader(http.StatusOK)
}

// Read 下载
// url格式: /stream/read/tokenxxxxxx?name=a.txt
func (s *DataStream) Read(w http.ResponseWriter, r *http.Request) {
	if nil == r.Form {
		if err := r.ParseForm(); nil != err {
			serviceutil.SendErrorAndStatus(w, http.StatusInternalServerError, err.Error())
			return
		}
	}
	// name
	if name := r.FormValue("name"); len(name) > 0 {
		w.Header().Set("Content-Type", strutil.GetMimeTypeBySuffix(name))
		w.Header().Set("Content-Disposition", "attachment; filename="+name)
	}
	// token
	token, err := s.dp.DoQueryToken(strutil.GetPathName(r.URL.Path))
	if nil != err {
		serviceutil.SendBadRequest(w, err.Error())
		return
	}
	//
	start, end, hasRange := serviceutil.GetRequestRange(r, token.FileSize)
	tr, err := s.dp.DoReadToken(token.Token, start)
	if nil != err {
		serviceutil.SendBadRequest(w, err.Error())
		return
	}
	{
		stransSize := end - start
		w.Header().Set("Content-Length", strconv.Itoa(int(stransSize)))
		if hasRange {
			w.Header().Set("Content-Range", "bytes "+strconv.Itoa(int(start))+"-"+strconv.Itoa(int(end-1))+"/"+strconv.Itoa(int(token.FileSize)))
			w.WriteHeader(http.StatusPartialContent)
		}
		if stransSize <= 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		twr := bizutils.NewTokenReaderWarp(token, refreshMilliSecond, tr, s.dp)
		//
		for {
			buf := make([]byte, 16384) // 16k
			n, err := twr.Read(buf)
			if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
				serviceutil.SendErrorAndStatus(w, http.StatusInternalServerError, err.Error())
				break
			} else if n == 0 {
				break
			} else if n > int(stransSize) {
				n = int(stransSize)
			}

			if _, err := w.Write(buf[:n]); nil != err {
				break
			} else {
				if stransSize = stransSize - int64(n); stransSize <= 0 {
					break
				}
			}
		}
	}
}
