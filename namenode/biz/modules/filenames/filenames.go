// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 文件命名空间管理, 文件操作(新建、删除、移动、复制等)

package filenames

import (
	"namenode/ifilestorage"
	"pakku/ipakku"
	"pakku/utils/fileutil"
	"pakku/utils/logs"
	"strconv"
	"time"
)

// FileNames 文件命名空间管理
type FileNames struct {
	NameManage
	c   ipakku.AppConfig    `@autowired:"AppConfig"`
	ev1 ipakku.AppSyncEvent `@autowired:"AppEvent"`
}

// AsModule 作为一个模块
func (fns *FileNames) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Name:        "FileNames",
		Version:     1.0,
		Description: "文件命名空间模块",
		OnReady: func(mctx ipakku.Loader) {
			fns.ev = fns.ev1
			deftDataSource := "./.datas/" + mctx.GetParam(ipakku.PARAMKEY_APPNAME).ToString("app") + "#filenames?cache=shared"
			confDataSource := fns.c.GetConfig("store.filenames.datasource").ToString(deftDataSource)
			if deftDataSource == confDataSource {
				fns.mkSqliteDIR()
			}
			fns.Setting(ifilestorage.DBSetting{
				DriverName:     fns.c.GetConfig("store.filenames.driver").ToString("sqlite3"),
				DataSourceName: confDataSource,
			})
		},
		OnSetup: func() {
			if err := fns.Install(); nil != err {
				logs.Panicln(err)
			}
		},
		OnInit: func() {
			start := time.Now()
			logs.Infoln("Start load names images")
			if err := fns.LoadCache(); nil != err {
				logs.Panicln(err)
			}
			spend := time.Now().Unix() - start.Unix()
			logs.Infoln("Load names images, spend " + strconv.Itoa(int(spend)) + " s")
			logs.Infoln("Loaed node count: " + strconv.Itoa(int(fns.NodeCount())))
			go fns.StartClearRunnable()
		},
	}
}

// mkSqliteDIR 创建sqlite文件存放目录
func (fns *FileNames) mkSqliteDIR() {
	if !fileutil.IsExist("./.datas") {
		if err := fileutil.MkdirAll("./.datas"); nil != err {
			logs.Panicln(err)
		}
	}
}
