// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package filenames

import (
	"errors"
	"fmt"
	"namenode/biz/modules/filenames/names"
	"namenode/ifilestorage"
	"testing"
	"time"

	"github.com/wup364/pakku"
	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/strutil"
	"github.com/wup364/pakku/utils/upool"

	_ "github.com/mattn/go-sqlite3"
)

func TestFileNames(t *testing.T) {
	fn := new(FileNames)
	pakku.NewApplication("filenames-test").EnableCoreModule().BootStart().LoadModule(fn)
	//
	worker := 10
	fileNum := 100
	rootDir := "/" + strutil.GetUUID()
	//
	var err error
	start := time.Now()
	completeCount := 0
	fw := upool.NewGoWorker(worker, worker)
	// 新建文件|文件夹
	beforCount := fn.NodeCount()
	for i := 0; i < worker; i++ {
		randomStr := strutil.GetRandom(16)
		baseDir := rootDir + "/" + randomStr[0:4] + "/" + randomStr[4:8] + "/" + randomStr[8:12] + "/" + randomStr[12:16]
		_, err = doMkDir(baseDir, fn)
		checkError(err)
		fw.AddJob(upool.NewSimpleJob(func(sj *upool.SimpleJob) {
			fileNum := sj.Args.(int)
			for i := 0; i < fileNum; i++ {
				_, err = doCreateFile(sj.ID, strutil.GetUUID(), fn)
				checkError(err)
			}
			// 所有线程都结束了
			completeCount++
			if completeCount >= worker {
				fw.CloseGoWorker()
			}
		}, baseDir, fileNum))
	}
	fw.WaitGoWorkerClose()
	fmt.Printf("doCreateFile(%d) --> %d(milli) \r\n", fn.NodeCount()-beforCount, time.Now().UnixMilli()-start.UnixMilli())
	// 查询列表
	list, _ := fn.ListNodeChilds(fn.GetNodeByPath(rootDir).Id, -1, -1)
	fmt.Println(list)
	// 新建文件 & 删除文件
	start = time.Now()
	beforCount = fn.NodeCount()
	fw = upool.NewGoWorker(worker, worker)
	fw.AddJob(upool.NewSimpleJob(func(sj *upool.SimpleJob) {
		fileNum := sj.Args.(int)
		delCount := 0
		for i := 0; i < fileNum; i++ {
			id, err := doCreateFile(sj.ID, strutil.GetUUID(), fn)
			checkError(err)
			go (func(id string) {
				checkError(fn.DelNode(id))
				delCount++
				if delCount >= fileNum {
					// 所有线程都结束了
					fw.CloseGoWorker()
				}
			})(id)
		}
		fmt.Printf("doCreateAndDelete -> %d(milli) \r\n", time.Now().UnixMilli()-start.UnixMilli())
	}, rootDir+"/"+list[0].Name, fileNum))
	fw.WaitGoWorkerClose()
	if fn.NodeCount() != beforCount {
		panic(fmt.Errorf("删除功能故障: %d", fn.NodeCount()-beforCount))
	}
	fmt.Printf("doCreateAndDelete --> %d(milli) \r\n", time.Now().UnixMilli()-start.UnixMilli())
	// 拷贝目录
	start = time.Now()
	beforCount = fn.NodeCount()
	checkError(doCopy(rootDir+"/"+list[0].Name, rootDir+"/"+list[0].Name+"_copyed", fn))
	fmt.Printf("doCopyDIR add=%d --> %d(milli) \r\n", fn.NodeCount()-beforCount, time.Now().UnixMilli()-start.UnixMilli())
	// 移动目录
	start = time.Now()
	beforCount = fn.NodeCount()
	checkError(doMove(rootDir+"/"+list[0].Name+"_copyed", rootDir+"/"+list[0].Name+"_moved", fn))
	fmt.Printf("doMoveDIR add=%d --> %d(milli) \r\n", fn.NodeCount()-beforCount, time.Now().UnixMilli()-start.UnixMilli())
}

// 递归移动-怎么费时怎么来
func doMove(src, dst string, fn *FileNames) (err error) {
	dstparent := strutil.GetPathParent(dst)
	srcNode := fn.GetNodeByPath(src)
	destNode := fn.GetNodeByPath(dstparent)
	fn.MoveNode(srcNode.Id, strutil.GetPathName(dst), destNode.Id)
	return err
}

// 递归拷贝-怎么费时怎么来
func doCopy(src, dst string, fn *FileNames) (err error) {
	dirList, count := fn.ListNodeChilds(fn.GetNodeByPath(src).Id, -1, -1)
	if count == 0 {
		_, err = doMkDir(dst, fn)
		return err
	}
	destid := ""
	dstparent := strutil.GetPathParent(dst)
	if destNode := fn.GetNodeByPath(dstparent); nil == destNode {
		if destid, err = doMkDir(dstparent, fn); nil != err {
			return err
		}
	} else {
		destid = destNode.Id
	}
	for i := 0; i < count; i++ {
		if fn.GetNodeByPath(dst+"/"+dirList[i].Name) != nil {
			continue
		}
		if node := fn.GetNodeByPath(src + "/" + dirList[i].Name); node != nil {
			if node.Flag == names.FLAG_DIR {
				err = doCopy(src+"/"+dirList[i].Name, dst+"/"+dirList[i].Name, fn)
			} else {
				_, err = fn.CopyNode(node.Id, dirList[i].Name, destid)
			}
		}
		if nil != err {
			break
		}
	}
	return err
}

// doCreateFile 创建文件
func doCreateFile(src, name string, fn *FileNames) (newid string, err error) {
	// fmt.Println(src + "  " + name)
	path := strutil.Parse2UnixPath(src + "/" + name)
	if fn.GetNodeByPath(path) != nil {
		return newid, errors.New("文件已存在, path:" + path)
	}
	var pid = ""
	if pnode := fn.GetNodeByPath(src); pnode == nil {
		pid, err = doMkDir(src, fn)
	} else {
		pid = pnode.Id
	}
	if nil == err {
		time := time.Now().UnixMilli()
		newid, err = fn.AddNode(ifilestorage.TNode4New{
			Pid:   pid,
			Name:  name,
			Addr:  strutil.GetUUID(),
			Flag:  ifilestorage.FLAG_NODETYPE_FILE,
			Ctime: time,
			Mtime: time,
		})
	}
	return newid, err
}

// doMkDir 创建多级文件夹
func doMkDir(src string, fn *FileNames) (newid string, err error) {
	src = strutil.Parse2UnixPath(src)
	if len(src) == 0 {
		return newid, errors.New("path is empty")
	}
	if src == "/" {
		return newid, errors.New("no access path '/'")
	}
	temp := src
	pid := ""
	makedirs := make([]string, 0)
	for len(temp) > 0 {
		node := fn.GetNodeByPath(temp)
		if nil == node {
			makedirs = append(makedirs, strutil.GetPathName(temp))
			temp = strutil.GetPathParent(temp)
			if len(temp) == 0 || temp == "/" {
				rnode := fn.GetNodeByPath("/")
				if nil == rnode {
					return newid, fileutil.PathNotExist("mkdir", "/")
				}
				pid = rnode.Id
				break
			}
		} else {
			pid = node.Id
			break
		}
	}
	if len(pid) == 0 {
		return newid, errors.New("make dir failed")
	}
	if len(makedirs) == 0 {
		return newid, errors.New("path is exist, path: " + src)
	}
	time := time.Now().UnixMilli()
	for i := len(makedirs) - 1; i > -1; i-- {
		newid, err = fn.AddNode(ifilestorage.TNode4New{
			Pid:   pid,
			Name:  makedirs[i],
			Flag:  ifilestorage.FLAG_NODETYPE_DIR,
			Ctime: time,
			Mtime: time,
		})
		if nil != err {
			break
		}
	}
	return newid, err
}

func checkError(err error) {
	if nil != err {
		panic(err)
	}
}
