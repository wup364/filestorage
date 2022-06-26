// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 文件夹上传工具 .\uploadtool.exe --scandir=C:\Program_UnZip --destdir=/RPC/W

package main

import (
	"testing"

	"github.com/wup364/filestorage/opensdk"

	"github.com/wup364/pakku/utils/strutil"
)

func TestDoUploadDir(t *testing.T) {
	o := opensdk.NewOpenApi(conf.GetConfig("rpc.address").ToString("127.0.0.1:5051"), opensdk.User{
		User:   conf.GetConfig("auth.user").ToString("OPENAPI"),
		Passwd: conf.GetConfig("auth.pwd").ToString(""),
	})

	datanodes := conf.GetConfig("datanodes").ToStrMap(nil)
	if nil == datanodes || len(datanodes) == 0 {
		o.SetDataNodeDNS(map[string]string{"DN101": "http://1270.0.1:5062"})
	} else {
		t := make(map[string]string)
		for key, val := range datanodes {
			t[key] = val.(string)
		}
		o.SetDataNodeDNS(t)
	}
	ov := true
	override = &ov
	doUploadDir("/RPC/w/"+strutil.GetRandom(5), "C:\\Program_UnZip", "C:\\Program_UnZip", o)
}

func TestDoUploadRandom(t *testing.T) {
	o := opensdk.NewOpenApi(conf.GetConfig("rpc.address").ToString("127.0.0.1:5051"), opensdk.User{
		User:   conf.GetConfig("auth.user").ToString("OPENAPI"),
		Passwd: conf.GetConfig("auth.pwd").ToString(""),
	})
	datanodes := conf.GetConfig("datanodes").ToStrMap(nil)
	if nil == datanodes || len(datanodes) == 0 {
		o.SetDataNodeDNS(map[string]string{"DN101": "http://1270.0.1:5062"})
	} else {
		t := make(map[string]string)
		for key, val := range datanodes {
			t[key] = val.(string)
		}
		o.SetDataNodeDNS(t)
	}
	ov := true
	override = &ov
	doUploadRandom("/RPC/w/"+strutil.GetRandom(5), o)
}
