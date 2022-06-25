// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package names

import (
	"errors"
	"fmt"
	"math/rand"
	"namenode/ifilestorage"
	"testing"
	"time"

	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/strutil"
	"github.com/wup364/pakku/utils/upool"

	_ "github.com/mattn/go-sqlite3"
)

func TestNameImage(t *testing.T) {
	fn := NewNameImage()
	checkError(fn.doCreateRootNode(ifilestorage.TNode4New{
		Id:    strutil.GetUUID(),
		Name:  "/",
		Addr:  strutil.GetUUID(),
		Flag:  FLAG_DIR,
		Ctime: time.Now().UnixMilli(),
		Mtime: time.Now().UnixMilli(),
	}))
	//
	worker := 10
	fileNum := 500000
	rootDir := "/" + strutil.GetUUID()
	//
	var err error
	start := time.Now()
	completeCount := 0
	fw := upool.NewGoWorker(worker, worker)
	// 新建文件|文件夹
	beforCount := fn.NodeCount()
	for i := 0; i < worker; i++ {
		randomStr := strutil.GetUUID()
		baseDir := rootDir + "/" + randomStr[0:8] + "/" + randomStr[8:16] + "/" + randomStr[16:24] + "/" + randomStr[24:32]
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
	fmt.Printf("doCreateFile %d-%d=%d --> %d(milli) \r\n", fn.NodeCount(), beforCount, fn.NodeCount()-beforCount, time.Now().UnixMilli()-start.UnixMilli())
	root := fn.GetRoot()
	fmt.Println("root:", root)
	// 文件夹大小
	start = time.Now()
	fmt.Println("GetFileSize", fn.GetFileSize(root.Id))
	fmt.Printf("GetFileSize(%d) --> %d(milli) \r\n", fn.NodeCount()-beforCount, time.Now().UnixMilli()-start.UnixMilli())
	if rsize := fn.GetFileSize(root.Id); rsize != root.Size {
		panic(fmt.Errorf("文件夹大小更新异常: %d - %d", root.Size, rsize))
	}
	// 查询列表
	list, _ := fn.GetNodeChildsByID(fn.GetNodeByPath(rootDir).Id, -1, -1)
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
				checkError(fn.RemoveNode(id))
				delCount++
				if delCount >= fileNum {
					// 所有线程都结束了
					fw.CloseGoWorker()
				}
			})(id)
		}
		fmt.Printf("doCreateAndDelete -> %d(milli) \r\n", time.Now().UnixMilli()-start.UnixMilli())
	}, rootDir+"/"+list[0].Name, 1000))
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
	// 删除目录
	fmt.Println(fn.GetNodeChildsByID(fn.GetNodeByPath(rootDir).Id, -1, -1))
}

// 递归移动-怎么费时怎么来
func doMove(src, dst string, fn *NameImage) (err error) {
	dstparent := strutil.GetPathParent(dst)
	srcNode := fn.GetNodeByPath(src)
	destNode := fn.GetNodeByPath(dstparent)
	fn.MoveNode(srcNode.Id, strutil.GetPathName(dst), destNode.Id)
	return err
}

// 递归拷贝-怎么费时怎么来
func doCopy(src, dst string, fn *NameImage) (err error) {
	dirList, count := fn.GetNodeChildsByID(fn.GetNodeByPath(src).Id, -1, -1)
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
			if node.Flag == FLAG_DIR {
				err = doCopy(src+"/"+dirList[i].Name, dst+"/"+dirList[i].Name, fn)
			} else {
				node.Id = strutil.GetUUID()
				node.Pid = destid
				node.Name = dirList[i].Name
				node.Mtime = time.Now().UnixMilli()
				err = fn.AddNode(ifilestorage.TNode4New{}.ReverseTNode(*node))
			}
		}
		if nil != err {
			break
		}
	}
	return err
}

// doCreateFile 创建文件
func doCreateFile(src, name string, fn *NameImage) (newid string, err error) {
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
		newid = strutil.GetUUID()
		err = fn.AddNode(ifilestorage.TNode4New{
			Id:    newid,
			Pid:   pid,
			Name:  name,
			Size:  int64(rand.Int31()),
			Addr:  strutil.GetUUID(),
			Flag:  FLAG_FILE,
			Ctime: time,
			Mtime: time,
		})
	}
	return newid, err
}

// doMkDir 创建多级文件夹
func doMkDir(src string, fn *NameImage) (newid string, err error) {
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
		newid = strutil.GetUUID()
		err = fn.AddNode(ifilestorage.TNode4New{
			Id:    newid,
			Pid:   pid,
			Name:  makedirs[i],
			Flag:  FLAG_DIR,
			Ctime: time,
			Mtime: time,
		})
		if nil != err {
			newid = ""
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
