// Copyright (C) 2019 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// HTTP客户端工具

package utils

import (
	"bytes"
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"time"
)

// DefaultClient 默认http client
var DefaultClient = &http.Client{Timeout: time.Second * time.Duration(30)}

// BuildURLWithMap 使用Map结构构件url请求参数
// {key:value} => /url/xxx?key=value
func BuildURLWithMap(url string, params map[string]string) string {
	result := url
	if lenP := len(params); lenP > 0 {
		result += "?"
		for key, val := range params {
			result += key + "=" + val
			if lenP > 1 {
				result += "&"
			}
			lenP--
		}
	}
	return result
}

// BuildURLWithArray 使用二维数组结构构件url请求参数
// [[key,value], ...] => /url/xxx?key=value
func BuildURLWithArray(url string, params [][]string) string {
	result := url
	if lenP := len(params); lenP > 0 {
		result += "?"
		for i := 0; i < lenP; i++ {
			if len(params[i]) >= 2 {
				result += params[i][0] + "=" + params[i][1]
				if i < lenP-1 {
					result += "&"
				}
			}
		}
	}
	return result
}

// Get Get请求
func Get(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	return Request4URL(DefaultClient, http.MethodGet, url, params, headers)
}

// Post Post请求
func Post(url string, params map[string]string, headers map[string]string) (*http.Response, error) {
	return Request4URL(DefaultClient, http.MethodPost, url, params, headers)
}

// PostJSON 通过Post Json 内容发送请求
func PostJSON(client *http.Client, method, url string, params interface{}, headers map[string]string) (*http.Response, error) {
	return Request4JSON(DefaultClient, method, url, params, headers)
}

// Request4URL 发送请求, 认使用 application/x-www-form-urlencoded 方式发送请求
func Request4URL(client *http.Client, method, url string, params, headers map[string]string) (*http.Response, error) {
	return DoRequest(client, method, "application/x-www-form-urlencoded", BuildURLWithMap(url, params), nil, headers)
}

// Request4JSON 通过 Json 内容发送请求
func Request4JSON(client *http.Client, method, url string, params interface{}, headers map[string]string) (*http.Response, error) {
	// build query
	if query, err := json.Marshal(params); err != nil {
		return nil, err
	} else {
		return DoRequest(client, method, "application/json;charset=utf-8", url, bytes.NewBuffer(query), headers)
	}
}

// PostFile 发送文件使用默认的file表单字段
func PostFile(url, filePath string, headers map[string]string) (*http.Response, error) {
	if reader, err := os.Open(filePath); err != nil {
		return nil, err
	} else {
		defer reader.Close()
		return DoUploadFile(DefaultClient, http.MethodPost, url, reader, headers, "", GetPathName(filePath))
	}
}

// PostFile 发送文件使用默认的file表单字段
func PutFile(url, filePath string, headers map[string]string) (*http.Response, error) {
	if reader, err := os.Open(filePath); err != nil {
		return nil, err
	} else {
		defer reader.Close()
		return DoUploadFile(DefaultClient, http.MethodPut, url, reader, headers, "", GetPathName(filePath))
	}
}

// DoUploadFile 上传文件文件 form-data
func DoUploadFile(client *http.Client, method, url string, reader io.Reader, headers map[string]string, fieldname, filename string) (*http.Response, error) {
	bodyBuf := bytes.NewBufferString("")
	bodyWriter := multipart.NewWriter(bodyBuf)
	if len(fieldname) == 0 {
		fieldname = "file"
	}
	if _, err := bodyWriter.CreateFormFile(fieldname, filename); err != nil {
		return nil, err
	}
	//
	boundary := bodyWriter.Boundary()
	body := io.MultiReader(bodyBuf, reader, bytes.NewBufferString("\r\n--"+boundary+"--\r\n"))
	return DoRequest(client, method, "multipart/form-data; boundary="+boundary, url, body, headers)
}

// DoRequest 发送请求(client, 请求方式, Content-Type, url)
func DoRequest(client *http.Client, method, contentType, url string, content io.Reader, header map[string]string) (*http.Response, error) {
	// build request method
	if req, err := http.NewRequest(method, url, content); err != nil {
		return nil, err
	} else {
		if len(contentType) > 0 {
			req.Header.Set("Content-Type", contentType)
		}
		if len(header) > 0 {
			for k, v := range header {
				req.Header.Set(k, v)
			}
		}
		return client.Do(req)
	}
}
