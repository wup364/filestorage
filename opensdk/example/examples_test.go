// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package example

import (
	"fmt"
	"opensdk"
	"os"
	"testing"
	"time"
)

func TestUploadFile(t *testing.T) {
	o := getOpenAPISDK()
	// get /opensdk/test/VSCode-win32-x64-1.66.1.zip C:\Users\wupen\Desktop\VSCode-win32-x64-1.66.1.zip -override
	if err := UploadFile("/opensdk/test/", "C:\\Users\\wupen\\Downloads\\VSCode-win32-x64-1.66.1.zip", true, o); nil != err {
		panic(err)
	}
}

func TestDownload(t *testing.T) {
	o := getOpenAPISDK()
	serverPath := "/opensdk/test/01.mp4"
	localPath := "C:/Users/wupen/Downloads/01.mp4"
	if err := DownloadFile(serverPath, localPath, o); nil != err {
		panic(err)
	}
}
func TestUploadRandom(t *testing.T) {
	o := getOpenAPISDK()
	uploadRandom("/opensdk/test", o)
}

func TestDownloadRandom(t *testing.T) {
	o := getOpenAPISDK()
	serverPath := uploadRandom("/opensdk/test", o)
	localPath := os.TempDir() + "/" + serverPath
	if err := DownloadFile(serverPath, localPath, o); nil != err {
		panic(err)
	}
}

func TestLoop(t *testing.T) {
	o := getOpenAPISDK()
	loopCount := 200
	c := make(chan int, loopCount)
	start := time.Now().UnixMilli()
	for i := 0; i < loopCount; i++ {
		go func(ci int) {
			for i := 0; i < 1000; i++ {
				o.IsDir("/")
			}
			c <- ci
		}(i)
	}
	for i := 0; i < loopCount; i++ {
		fmt.Println(<-c)
	}
	fmt.Printf("end-01: %d \r\n", time.Now().UnixMilli()-start)
	//
	start = time.Now().UnixMilli()
	for i := 0; i < 1000; i++ {
		o.IsDir("/")
	}
	fmt.Printf("end-02: %d \r\n", time.Now().UnixMilli()-start)
}

func TestDoUpdatePWD(t *testing.T) {
	o := getMGAPISDK()
	if ok, err := o.DoUpdatePWD("NAMENODE", ""); nil != err {
		panic(err)
	} else {
		fmt.Println(ok)
	}
}

// getOpenAPISDK 获取sdk实例
func getOpenAPISDK() opensdk.IOpenApi {
	o := opensdk.NewOpenApi(
		"127.0.0.1:5051",
		opensdk.User{
			User:   "OPENAPI",
			Passwd: "",
		})
	// o.SetDataNodeDNS(map[string]string{"DN101": "http://127.0.0.1:5062"})
	return o
}

// getMGAPISDK 获取sdk实例
func getMGAPISDK() opensdk.IServerMG {
	o := opensdk.NewSvrMngApi(
		"127.0.0.1:5051",
		opensdk.User{
			User:   "NAMENODE",
			Passwd: "",
		})
	return o
}
