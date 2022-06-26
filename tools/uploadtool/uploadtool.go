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
	"flag"
	"io"
	"path/filepath"
	"strings"

	"github.com/wup364/filestorage/opensdk"
	"github.com/wup364/filestorage/opensdk/utils"

	"github.com/wup364/pakku/modules/appconfig/jsonconfig"
	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

var scandir *string
var destdir *string
var override *bool
var random *bool
var logger *string
var loglevel *string
var logdir *string
var conf *jsonconfig.Config

func init() {
	conf = &jsonconfig.Config{}
	conf.Init("uploadtool")
}

func main() {
	parseFlag()
	settingLogs()
	//
	logs.Infoln("Start upload!")
	o := getSDK()
	// 执行上传
	if *random {
		for {
			doUploadRandom(*destdir, o)
		}
	} else {
		doUploadDir(*destdir, *scandir, *scandir, o)
	}
	logs.Infoln("End upload!")
}

func parseFlag() {
	scandir = flag.String("scandir", "", "The local folder that needs to be scanned, default null ")
	destdir = flag.String("destdir", "/", "The destination folder to upload to, default / ")
	override = flag.Bool("override", false, "Whether to overwrite existing files, default false ")
	random = flag.Bool("random", false, "Whether to overwrite existing files, default false ")
	logger = flag.String("logger", "console", "logger: console, file or unset, default console")
	loglevel = flag.String("loglevel", "debug", "loglevel: debug, info, error or none, default debug")
	logdir = flag.String("logdir", "./logs", "default ./logs/{application name}.log")
	flag.Parse()
	//
	if !*random && *scandir == "" {
		panic("scandir can not empty!")
	}
}

// doUploadDir 上传文件夹
func doUploadDir(destDir, loaclBaseDir, localDir string, o opensdk.IOpenApi) {
	baseLen := len(loaclBaseDir)
	dirs, err := fileutil.GetDirList(localDir)
	checkError(err)
	for _, val := range dirs {
		localFile := filepath.Clean(localDir + "/" + val)
		if fileutil.IsDir(localFile) {
			doUploadDir(destDir, loaclBaseDir, localFile, o)
		} else {
			uploadPath := strutil.Parse2UnixPath(destDir + "/" + (localFile[baseLen:]))
			if nil == override || !*override {
				if ok, err := o.IsExist(uploadPath); nil == err && ok {
					logs.Infoln("Skip existing", uploadPath)
					continue
				} else if nil != err {
					logs.Errorln(err)
					continue
				}
			}
			token, err := o.DoAskWriteToken(uploadPath)
			checkError(err)
			err = opensdk.FileUploader(localFile, token, 128*1024*1024, o.DoWriteToken)
			checkError(err)
			node, err := o.DoSubmitWriteToken(token.Token, *override)
			checkError(err)
			logs.Infoln("Uploaded successfully", uploadPath, node)
		}
	}
}

// doUploadRandom 随机上传文件
func doUploadRandom(destDir string, o opensdk.IOpenApi) {
	random := strutil.GetRandom(16)
	logs.Infoln(random)
	uploadPath := strutil.Parse2UnixPath(destDir + "/" + random[:3] + "/" + random[3:6] + "/" + random[6:9] + "/" + random[9:])
	token, err := o.DoAskWriteToken(uploadPath)
	checkError(err)
	err = streamUploader(random, token, o.DoWriteToken)
	checkError(err)
	node, err := o.DoSubmitWriteToken(token.Token, true)
	checkError(err)
	// err = o.DoDelete(uploadPath)
	// checkError(err)
	logs.Infoln("Uploaded successfully", uploadPath, node)
}

type doWriteTokenFunc func(nodeNo, token, endpoint string, pieceNumber int, sha256 string, reader io.Reader) (err error)

// streamUploader streamUploader
func streamUploader(random string, token *opensdk.StreamToken, doWriteToken doWriteTokenFunc) (err error) {
	var sha256 string
	reader := strings.NewReader(random)
	if sha256, err = utils.GetFileSHA256(reader); nil == err {
		reader.Seek(0, io.SeekStart)
		err = doWriteToken(token.NodeNo, token.Token, token.EndPoint, 1, sha256, reader)
	}
	return err
}

func getSDK() opensdk.IOpenApi {
	o := opensdk.NewOpenApi(conf.GetConfig("rpc.address").ToString("127.0.0.1:5051"), opensdk.User{
		User:   conf.GetConfig("auth.user").ToString("OPENAPI"),
		Passwd: conf.GetConfig("auth.pwd").ToString(""),
	})
	// datanode 路由
	if datanodes := conf.GetConfig("datanodes").ToStrMap(nil); len(datanodes) == 0 {
		o.SetDataNodeDNS(map[string]string{"DN101": "http://127.0.0.1:5062"})
	} else {
		t := make(map[string]string)
		for key, val := range datanodes {
			t[key] = val.(string)
		}
		o.SetDataNodeDNS(t)
	}
	return o
}
func settingLogs() {
	if *logger == "file" {
		f, err := fileutil.GetWriter(*logdir + "/uploadtool.log")
		checkError(err)
		logs.SetOutput(f)
	}
	switch *loglevel {
	case "none":
		logs.SetLoggerLevel(logs.NONE)
	case "error":
		logs.SetLoggerLevel(logs.ERROR)
	case "info":
		logs.SetLoggerLevel(logs.INFO)
	default:
		logs.SetLoggerLevel(logs.DEBUG)
	}
}
func checkError(err error) {
	if nil != err {
		panic(err)
	}
}
