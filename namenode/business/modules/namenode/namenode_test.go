// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 树

package namenode

import (
	"fmt"
	"namenode/business/modules/filenames"
	"namenode/business/modules/user4rpc"
	"namenode/ifilestorage"
	"strconv"
	"testing"
	"time"

	"github.com/wup364/pakku"
	"github.com/wup364/pakku/utils/strutil"
	"github.com/wup364/pakku/utils/upool"
)

func TestNameNode(t *testing.T) {
	app := pakku.NewApplication("filenames-test").EnableCoreModule().EnableNetModule().BootStart()
	// 获取对象
	var ft ifilestorage.NameNode
	app.LoadModules(new(user4rpc.User4RPC), new(filenames.FileNames), new(NameNode)).GetModuleByName(new(NameNode).AsModule().Name, &ft)
	jobCount := 20
	completedCount := 0
	beforeTme := time.Now()
	fw := upool.NewGoWorker(20, jobCount)
	for i := 0; i < jobCount; i++ {
		fw.AddJob(upool.NewSimpleJob(func(sj *upool.SimpleJob) {
			parent := sj.ID
			maxdepth := sj.Args.(int)
			for maxdepth > 0 {
				maxdepth--
				_, err := ft.DoMkDir(parent)
				checkError(err)
				parent = parent + "/" + strutil.GetRandom(50) + "_" + strconv.Itoa(maxdepth)
				for i := 0; i < maxdepth; i++ {
					_, err := ft.DoMkDir(parent + "_" + strconv.Itoa(i))
					checkError(err)
				}
			}
			completedCount++
			if completedCount >= jobCount {
				fw.CloseGoWorker()
			}
		}, "/"+strutil.GetUUID(), 20))
	}
	fw.WaitGoWorkerClose()
	fmt.Println("MKDIR-LOOP ---> ", time.Now().UnixMilli()-beforeTme.UnixMilli())
	//
	testPath := strutil.GetUUID()
	newName := strutil.GetUUID()
	beforeTme = time.Now()
	ft.DoMkDir("/" + testPath + "/" + strutil.GetUUID() + "/" + strutil.GetUUID())
	fmt.Println("MKDIR ---> ", time.Now().UnixMilli()-beforeTme.UnixMilli())
	//
	beforeTme = time.Now()
	err := ft.DoRename("/"+testPath, newName)
	fmt.Println("RENAME ---> ", time.Now().UnixMilli()-beforeTme.UnixMilli())
	checkError(err)
	//
	beforeTme = time.Now()
	list, _ := ft.GetDirNameList("/"+newName, -1, -1)
	fmt.Println(list)
	fmt.Println("LIST ---> ", time.Now().UnixMilli()-beforeTme.UnixMilli())
	//
	beforeTme = time.Now()
	err = ft.DoDelete("/" + newName)
	checkError(err)
	fmt.Println("DEL ---> ", time.Now().UnixMilli()-beforeTme.UnixMilli())
	list, _ = ft.GetDirNameList("/"+newName, -1, -1)
	fmt.Println(list)
	//
	cplist, _ := ft.GetDirNameList("/", -1, -1)
	if len(cplist) == 0 {
		panic("get dir list nil")
	}
	fmt.Println(doCopy("/"+cplist[0], "/"+cplist[0], ft))
	//
	beforeTme = time.Now()
	for i := 0; i < 10; i++ {
		cpdst := "/" + cplist[0] + "_" + strutil.GetUUID()
		checkError(doCopy("/"+cplist[0], cpdst, ft))
		if !ft.IsExist(cpdst) {
			panic("拷贝失败, path=" + cpdst)
		}
	}
	fmt.Println("COPY ---> ", time.Now().UnixMilli()-beforeTme.UnixMilli())
	//
	mverr := ft.DoMove("/"+cplist[0], "/"+cplist[0], false)
	fmt.Println(mverr)
	//
	beforeTme = time.Now()
	mvdst := "/" + cplist[0] + "_" + strutil.GetUUID()
	mverr = ft.DoMove("/"+cplist[0], mvdst, false)
	fmt.Println("MOVE ---> ", time.Now().UnixMilli()-beforeTme.UnixMilli())
	checkError(mverr)
	//
	mverr = ft.DoMove("/"+cplist[1], "/"+cplist[1]+"/"+strutil.GetUUID(), false)
	fmt.Println("MOVE1 ---> ", time.Now().UnixMilli()-beforeTme.UnixMilli())
	checkError(mverr)
	//
	beforeTme = time.Now()
	delerr := ft.DoDelete(mvdst)
	fmt.Println("DELETE-DIR ---> ", time.Now().UnixMilli()-beforeTme.UnixMilli())
	checkError(delerr)

	time.Sleep(time.Duration(1) * time.Second)
}

// 递归拷贝
func doCopy(src, dst string, ft ifilestorage.NameNode) (err error) {
	dirList, count := ft.GetDirNameList(src, -1, -1)
	if count == 0 {
		_, err = ft.DoMkDir(dst)
		return err
	}
	for i := 0; i < count; i++ {
		if ft.IsExist(dst + "/" + dirList[i]) {
			continue
		}
		if ft.IsDir(src + "/" + dirList[i]) {
			err = doCopy(src+"/"+dirList[i], dst+"/"+dirList[i], ft)
		} else {
			_, err = ft.DoCopy(src+"/"+dirList[i], dst+"/"+dirList[i], false)
		}
		if nil != err {
			break
		}
	}
	return err
}

func checkError(err error) {
	if nil != err {
		panic(err)
	}
}
