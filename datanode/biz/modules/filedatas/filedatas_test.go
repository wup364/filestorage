// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 文件管理模块, 文件操作(新建、删除、移动、复制等)、虚拟分区挂载

package filedatas

import (
	"datanode/ifilestorage"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/wup364/pakku"
	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/modules/appconfig"
	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

func TestFileDatas(t *testing.T) {
	app := pakku.NewApplication("filedatas-test").EnableCoreModule().BootStart()
	var conf ipakku.AppConfig
	app.GetModuleByName(new(appconfig.AppConfig).AsModule().Name, &conf)
	// 挂载目录
	conf.SetConfig(CONFKEY_MOUNT+"./."+CONFKEY_MOUNTTYPE, "LOCAL")
	conf.SetConfig(CONFKEY_MOUNT+"./."+CONFKEY_MOUNTADDR, os.TempDir())
	conf.SetConfig(CONFKEY_MOUNT+"./test."+CONFKEY_MOUNTTYPE, "LOCAL")
	conf.SetConfig(CONFKEY_MOUNT+"./test."+CONFKEY_MOUNTADDR, os.TempDir()+"/"+app.GetInstanceID())
	// conf.SetConfig(CONFKEY_MOUNT+"./test."+CONFKEY_MOUNTADDR, "D:\\"+app.GetInstanceID())
	// 获取对象
	var fsm ifilestorage.FileDatas
	app.LoadModule(new(FileDatas)).GetModuleByName(new(FileDatas).AsModule().Name, &fsm)
	defer fileutil.RemoveAll(os.TempDir() + "/" + app.GetInstanceID())
	defer fileutil.RemoveAll(os.TempDir() + "/.sys")
	// 测试 GetDirList
	rootdirs := fsm.GetDirList("/", -1, -1)
	if len(rootdirs) == 0 {
		panic("dir: " + os.TempDir() + " is empty")
	}
	wkdir := "/test"
	// DoWrite
	writeFile := wkdir + "/txt.txt"
	defer fsm.DoDelete(writeFile)
	err := fsm.DoWrite(writeFile, strings.NewReader(strutil.GetUUID()))
	checkErr(err)
	wkdirs := fsm.GetDirList(wkdir, -1, -1)
	if len(wkdirs) == 0 {
		panic("DoWrite: " + writeFile + " failed")
	}
	fmt.Println(wkdirs)
	// GetDirListInfo
	// fsinfo, err := fsm.GetDirNodeList(wkdir)
	// checkErr(err)
	// logs.Infoln(fsinfo)
	// DoCopy
	copyFile := "/copy.txt"
	err = fsm.DoCopy(writeFile, copyFile, true)
	checkErr(err)
	// IsFile
	if !fsm.IsFile(copyFile) {
		logs.Panicln("DoCopy failed")
	}
	// DoMove
	moveFile := wkdir + "/move.txt"
	err = fsm.DoMove(copyFile, moveFile, true)
	checkErr(err)
	// IsFile
	if !fsm.IsFile(moveFile) {
		logs.Panicln("DoMove failed")
	}
	// 复制多个
	moveDir := "/movedir"
	defer fsm.DoDelete(moveDir)
	for i := 0; i < 10; i++ {
		err = fsm.DoCopy(moveFile, moveDir+"/"+strconv.Itoa(i)+".txt", true)
		checkErr(err)
	}
	// DoDelete
	checkErr(fsm.DoDelete(moveFile))
	// 复制多个
	copyDir := wkdir + "/" + moveDir
	defer fsm.DoDelete(copyDir)
	err = fsm.DoCopy(moveDir, copyDir, true)
	checkErr(err)
	// 移动多个
	err = fsm.DoMove(moveDir, copyDir, true)
	checkErr(err)
}

func checkErr(err error) {
	if nil != err {
		panic(err)
	}
}
