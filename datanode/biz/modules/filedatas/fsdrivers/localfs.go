// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 本地文件挂载操作驱动

package fsdrivers

import (
	"datanode/biz/modules/filedatas/ifiledatas"
	"errors"
	"io"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

const (
	sysDir      = ".sys"                // 系统文件存放位置
	tempDir     = sysDir + "/.cache"    // 系统文件-缓存位置
	deletingDir = sysDir + "/.deleting" // 系统文件-待删除文件位置
)

// LocalDriver 本地文件挂载操作驱动
type LocalDriver struct {
	mtn *ifiledatas.MountNode
	mtm ifiledatas.DIRMount
}

// GetDriverType 当驱动注册时调用
func (locl LocalDriver) GetDriverType() string {
	return "LOCAL"
}

// InstanceDriver 当驱动实例化时调用
func (locl LocalDriver) InstanceDriver(dirMount ifiledatas.DIRMount, mtnode *ifiledatas.MountNode) (ifiledatas.FileDriver, error) {
	// 格式化 mtnode.Addr
	if !filepath.IsAbs(mtnode.Addr) {
		if path, err := filepath.Abs(mtnode.Addr); nil != err {
			return nil, err
		} else {
			mtnode.Addr = path
		}
	}
	mtnode.Addr = filepath.Clean(mtnode.Addr)
	// 初始化必要的文件夹
	if loclTemp := filepath.Clean(mtnode.Addr + "/" + tempDir); !fileutil.IsDir(loclTemp) {
		if err := fileutil.MkdirAll(loclTemp); nil != err {
			return nil, errors.New("Create Folder Failed, Path: " + loclTemp + ", " + err.Error())
		}
	}
	// 初始化'删除缓存'文件夹
	if loclDeleting := filepath.Clean(mtnode.Addr + "/" + deletingDir); !fileutil.IsDir(loclDeleting) {
		if err := fileutil.MkdirAll(loclDeleting); nil != err {
			return nil, errors.New("Create Folder Failed, Path: " + loclDeleting + ", " + err.Error())
		}
	}
	// 实例化
	instance := &LocalDriver{
		mtn: mtnode,
		mtm: dirMount,
	}
	// 启动'temp文件'维护线程
	go instance.startCacheCleaner(mtnode)
	// 启动'删除缓存'维护线程
	go instance.startDeleteCacheCleaner(mtnode)
	return instance, nil
}

// IsExist 文件是否存在
func (locl *LocalDriver) IsExist(relativePath string) bool {
	absPath, _, err := locl.getAbsolutePath(locl.mtn, relativePath)
	if nil != err {
		logs.Errorln(err)
		return false
	}
	return fileutil.IsExist(absPath)
}

// IsDir IsDir
func (locl *LocalDriver) IsDir(relativePath string) bool {
	absPath, _, err := locl.getAbsolutePath(locl.mtn, relativePath)
	if nil != err {
		logs.Errorln(err)
		return false
	}
	return fileutil.IsDir(absPath)
}

// IsFile IsFile
func (locl *LocalDriver) IsFile(relativePath string) bool {
	absPath, _, err := locl.getAbsolutePath(locl.mtn, relativePath)
	if nil != err {
		logs.Errorln(err)
		return false
	}
	return fileutil.IsFile(absPath)
}

// GetNode GetNode
func (locl *LocalDriver) GetNode(src string) *ifiledatas.Node {
	absPath, _, err := locl.getAbsolutePath(locl.mtn, src)
	if nil != err {
		logs.Errorln(err)
		return nil
	}
	//
	node := ifiledatas.Node{
		Path:   strutil.Parse2UnixPath(src),
		IsFile: fileutil.IsFile(absPath),
		IsDir:  fileutil.IsDir(absPath),
	}
	if !node.IsFile && !node.IsDir {
		return nil
	}
	if t, err := fileutil.GetModifyTime(absPath); nil == err {
		node.Mtime = t.UnixMilli()
	}
	if node.IsFile {
		if s, err := fileutil.GetFileSize(absPath); nil == err {
			node.Size = s
		}
	}
	return &node
}

// GetNodes GetNodes
func (locl *LocalDriver) GetNodes(src []string, ignoreNotIsExist bool) ([]ifiledatas.Node, error) {
	if len(src) == 0 {
		return nil, nil
	}
	if ignoreNotIsExist {
		res := make([]ifiledatas.Node, 0)
		for i := 0; i < len(src); i++ {
			if node := locl.GetNode(src[i]); nil != node {
				res = append(res, *locl.GetNode(src[i]))
			}
		}
		return res, nil
	} else {
		res := make([]ifiledatas.Node, len(src))
		for i := 0; i < len(src); i++ {
			if node := locl.GetNode(src[i]); nil != node {
				res[i] = *locl.GetNode(src[i])
			} else {
				return nil, fileutil.PathNotExist("GetNodes", src[i])
			}
		}
		return res, nil
	}
}

// GetDirNodeList 获取node列表
func (locl *LocalDriver) GetDirNodeList(src string, limit, offset int) ([]ifiledatas.Node, error) {
	dirNames := locl.GetDirList(src, limit, offset)
	if len(dirNames) == 0 {
		return nil, nil
	}
	for i := 0; i < len(dirNames); i++ {
		dirNames[i] = strutil.Parse2UnixPath(src + "/" + dirNames[i])
	}
	return locl.GetNodes(dirNames, true)
}

// GetDirList 获取路径列表, 返回相对路径
func (locl *LocalDriver) GetDirList(relativePath string, limit, offset int) []string {
	if absPath, mRelativePath, err := locl.getAbsolutePath(locl.mtn, relativePath); err == nil {
		if ls, err := fileutil.GetDirList(absPath); nil == err {
			// 如果是挂载目录根目录, 需要处理 缓存目录
			if mRelativePath == "/" && len(ls) > 0 {
				res := make([]string, 0)
				for i := 0; i < len(ls); i++ {
					// 如果是挂载目录根目录, 忽略系统目录
					if sysDir == ls[i] {
						continue
					}
					res = append(res, ls[i])
				}
				ls = res
			}
			// limit offset
			if count := len(ls); count > 0 && limit > 0 && offset > -1 {
				if offset < count {
					if limit < count {
						limit = count
					}
					index := 0
					ls = make([]string, limit)
					for i := offset; i < count; i++ {
						ls[index] = ls[i]
						index++
						if index == limit {
							break
						}
					}
					if index < limit {
						ls = ls[:index]
					}
				} else {
					ls = make([]string, 0)
				}
			}
			return ls
		}
	}
	return make([]string, 0)
}

// GetFileSize GetFileSize
func (locl *LocalDriver) GetFileSize(relativePath string) int64 {
	if absPath, _, err := locl.getAbsolutePath(locl.mtn, relativePath); nil == err {
		if fileutil.IsFile(absPath) {
			if size, err := fileutil.GetFileSize(absPath); nil == err {
				return size
			}
		} else {
			return 0
		}
	}
	return -1
}

// GetModifyTime GetModifyTime
func (locl *LocalDriver) GetModifyTime(relativePath string) int64 {
	if absPath, _, err := locl.getAbsolutePath(locl.mtn, relativePath); nil == err {
		if time, err := fileutil.GetCreateTime(absPath); nil == err {
			return time.UnixMilli()
		}
	}
	return -1
}

// DoMove 移动文件|夹
func (locl *LocalDriver) DoMove(src string, dst string, replace bool) error {
	if locl.mtn.Path == src {
		return locl.wrapError(src, dst, errors.New("Does not allow access: "+src))
	}
	absSrc, _, err := locl.getAbsolutePath(locl.mtn, src)
	if nil != err {
		return locl.wrapError(src, dst, err)
	}
	if locl.mtn.Addr == absSrc {
		return locl.wrapError(src, dst, errors.New(src+" is mount root, cannot move"))
	}
	// 目标位置驱动接口
	dstMountItem := locl.mtm.GetMountNode(dst)
	switch dstMountItem.Type {
	case locl.GetDriverType():
		{ // 本地存储
			absDst, _, err := locl.getAbsolutePath(dstMountItem, dst)
			if nil != err {
				return locl.wrapError(src, dst, err)
			}
			err = fileutil.MoveFilesAcrossDisk(absSrc, absDst, replace, false, func(srcPath, dstPath string, err error) error {
				rSrc := locl.getRelativePath(locl.mtn, srcPath)
				rDst := locl.getRelativePath(dstMountItem, dstPath)
				return locl.wrapError(rSrc, rDst, errors.New("Move failed: "+rSrc+" -> "+rDst))
			})
			return locl.wrapError(src, dst, err)
		}
	default:
		{ // 不支持的分区挂载类型
			return locl.wrapError(src, dst, errors.New("Unsupported partition mount type: "+dstMountItem.Type))
		}
	}
}

// DoRename 重命名文件|文件夹
func (locl *LocalDriver) DoRename(relativePath string, newName string) error {
	if locl.mtn.Path == relativePath {
		return locl.wrapError(relativePath, "", errors.New("Does not allow access: "+relativePath))
	}
	absSrc, _, err := locl.getAbsolutePath(locl.mtn, relativePath)
	if nil != err {
		return locl.wrapError(relativePath, "", err)
	}
	if len(newName) > 0 {
		return locl.wrapError(relativePath, "", fileutil.Rename(absSrc, newName))
	}
	return nil
}

// DoMkDir 新建文件夹
func (locl *LocalDriver) DoMkDir(relativePath string) error {
	if locl.mtn.Path == relativePath {
		return locl.wrapError(relativePath, "", errors.New("Does not allow access: "+relativePath))
	}
	absSrc, _, err := locl.getAbsolutePath(locl.mtn, relativePath)
	if nil != err {
		return locl.wrapError(relativePath, "", err)
	}
	return locl.wrapError(relativePath, "", fileutil.MkdirAll(absSrc))
}

// DoDelete 删除文件|文件夹
func (locl *LocalDriver) DoDelete(relativePath string) error {
	if locl.mtn.Path == relativePath {
		return locl.wrapError(relativePath, "", errors.New("Does not allow access: "+relativePath))
	}
	absSrc, _, err := locl.getAbsolutePath(locl.mtn, relativePath)
	if nil != err {
		return locl.wrapError(relativePath, "", err)
	}
	deletingPath := locl.getAbsoluteDeletingPath(locl.mtn)
	for fileutil.IsExist(deletingPath) {
		deletingPath = locl.getAbsoluteDeletingPath(locl.mtn)
	}
	// 移动到删除零时目录, 如果存在则覆盖
	// 通过这种方式可以减少函数等待时间, 但是如果线程删除失败则可能导致文件无法删除
	// 所以再启动或者周期性的检擦删除零时目录, 进行清空
	mvErr := fileutil.MoveFiles(absSrc, deletingPath, true, false)
	return locl.wrapError(relativePath, "", mvErr)
}

// DoCopy 拷贝文件
func (locl *LocalDriver) DoCopy(src, dst string, replace bool) error {
	absSrc, _, err := locl.getAbsolutePath(locl.mtn, src)
	if nil != err {
		return locl.wrapError(src, dst, err)
	}
	// 目标位置驱动接口
	dstMountItem := locl.mtm.GetMountNode(dst)
	switch dstMountItem.Type {
	case locl.GetDriverType():
		{ // 本地存储
			absDst, _, err := locl.getAbsolutePath(dstMountItem, dst)
			if nil != err {
				return locl.wrapError(src, dst, err)
			}
			if fileutil.IsFile(absSrc) {
				return locl.wrapError(src, dst, fileutil.CopyFile(absSrc, absDst, replace, false))
			}
			err = fileutil.CopyFiles(absSrc, absDst, replace, false, func(srcPath, dstPath string, err error) error {
				rSrc := locl.getRelativePath(locl.mtn, srcPath)
				rDst := locl.getRelativePath(dstMountItem, dstPath)
				return locl.wrapError(rSrc, rDst, err)
			})
			return locl.wrapError(src, dst, err)
		}
	default:
		{ // 不支持的分区挂载类型
			return locl.wrapError(src, dst, errors.New("Unsupported partition mount type: "+dstMountItem.Type))
		}
	}
}

// DoRead 读取文件, 需要手动关闭流
func (locl *LocalDriver) DoRead(relativePath string, offset int64) (io.ReadCloser, error) {
	absDst, _, err := locl.getAbsolutePath(locl.mtn, relativePath)
	if nil != err {
		return nil, locl.wrapError(relativePath, "", err)
	}
	fs, err := fileutil.OpenFile(absDst)
	if nil != err {
		return nil, locl.wrapError(relativePath, "", err)
	}
	_, err = fs.Seek(offset, io.SeekStart)
	if nil != err {
		return nil, locl.wrapError(relativePath, "", err)
	}
	return fs, nil
}

// DoWrite 写入文件， 先写入临时位置, 然后移动到正确位置
func (locl *LocalDriver) DoWrite(relativePath string, ioReader io.Reader) error {
	if ioReader == nil {
		return locl.wrapError(relativePath, "", errors.New("IO Reader is nil"))
	}
	absDst, _, err := locl.getAbsolutePath(locl.mtn, relativePath)
	if nil != err {
		return locl.wrapError(relativePath, "", err)
	}
	tempPath := locl.getAbsoluteTempPath(locl.mtn)
	fs, wErr := fileutil.GetWriter(tempPath)
	if wErr != nil {
		return locl.wrapError(relativePath, "", wErr)
	}
	_, cpErr := io.Copy(fs, ioReader)
	if nil == cpErr {
		fsCloseErr := fs.Close()
		if fsCloseErr == nil {
			return locl.wrapError(relativePath, "", fileutil.MoveFiles(tempPath, absDst, true, false))
		}
		return locl.wrapError(relativePath, "", fsCloseErr)
	}
	fsCloseErr := fs.Close()
	if nil != fsCloseErr {
		return locl.wrapError(relativePath, "", fsCloseErr)
	}
	if fileutil.IsFile(tempPath) {
		rmErr := fileutil.RemoveFile(tempPath)
		if rmErr != nil {
			return locl.wrapError(relativePath, "", rmErr)
		}
	}
	return locl.wrapError(relativePath, "", cpErr)
}

// getAbsolutePath (mnode, 虚拟路径)(绝对位置, 挂载位置, 错误)处理路径拼接
func (locl *LocalDriver) getAbsolutePath(mtn *ifiledatas.MountNode, relativePath string) (abs string, rlPath string, err error) {
	rlPath = relativePath
	if mtn.Path != "/" {
		rlPath = relativePath[len(mtn.Path):]
		if rlPath == "" {
			rlPath = "/"
		}
	}
	// /Mount/.sys/.cache=>/.sys/.cache
	if rlPath == sysDir ||
		rlPath == "/"+sysDir ||
		strings.Index(rlPath, "/"+sysDir+"/") == 0 {
		return abs, rlPath, errors.New("Does not allow access: " + rlPath)
	}
	abs = filepath.Clean(mtn.Addr + rlPath)
	//fmt.Println( "getAbsolutePath: ", rlPath, abs )
	return
}

// getRelativePath 获取相对路径
func (locl *LocalDriver) getRelativePath(mti *ifiledatas.MountNode, absolute string) string {
	// fmt.Println("locl.getRelativePath: ", mti.Addr, absolute)
	absolute = filepath.Clean(absolute)
	if filepath.IsAbs(absolute) {
		if strings.HasPrefix(absolute, mti.Addr) {
			return strutil.Parse2UnixPath(mti.Path + "/" + absolute[len(mti.Addr):])
		} else if strings.HasPrefix(mti.Addr, absolute) {
			return mti.Path
		}
	}
	return absolute
}

// getAbsoluteTempPath 获取该分区下的缓存目录
func (locl *LocalDriver) getAbsoluteTempPath(mtn *ifiledatas.MountNode) string {
	return filepath.Clean(mtn.Addr + "/" + tempDir + "/" + strutil.GetUUID())
}

// getAbsoluteDeletingPath 获取一个放置删除文件的目录
func (locl *LocalDriver) getAbsoluteDeletingPath(mtn *ifiledatas.MountNode) string {
	return filepath.Clean(mtn.Addr + "/" + deletingDir + "/" + strutil.GetRandom(6) + strconv.FormatInt(time.Now().UnixNano(), 10))
}

// wrapLocalError 重新包装本地驱动错误信息, 避免真实路径暴露
func (locl *LocalDriver) wrapError(src, dst string, err error) error {
	if nil != err {
		return errors.New(locl.clearMountAddr(locl.mtm.GetMountNode(src), locl.mtm.GetMountNode(dst), err))
	}
	return nil
}

// clearMountAddr 去除挂载目录的位置信息
func (locl *LocalDriver) clearMountAddr(src, dst *ifiledatas.MountNode, err error) string {
	if nil != err {
		errStr := err.Error()
		// src
		srcAddr := ""
		if nil != src {
			srcAddr = filepath.Clean(src.Addr)
			if strings.Contains(errStr, srcAddr) {
				// /root/datas/a/b -> /a/b/a/b
				rStr := src.Path
				if src.Path == "/" {
					rStr = ""
				}
				errStr = strings.Replace(errStr, srcAddr, rStr, -1)
			}
		}
		// dst
		if nil != dst {
			dstAddr := filepath.Clean(dst.Addr)
			if dstAddr != srcAddr && strings.Contains(errStr, dstAddr) {
				rStr := dst.Path
				if dst.Path == "/" {
					rStr = ""
				}
				errStr = strings.Replace(errStr, dstAddr, rStr, -1)
			}
		}
		return errStr
	}
	return ""
}

// startCacheCleaner 启动'temp文件'维护线程
func (locl *LocalDriver) startCacheCleaner(mtnode *ifiledatas.MountNode) {
	maxTime := int64(1 * 24 * 60 * 60 * 1000)
	baseDIR := filepath.Clean(mtnode.Addr + "/" + tempDir)
	logs.Infof("startCacheCleaner dir=%s, exp=%d\r\n", baseDIR, maxTime)
	for {
		var err error
		var dirs []string
		if dirs, err = fileutil.GetDirList(baseDIR); nil == err {
			if len(dirs) > 0 {
				nowTime := time.Now().UnixMilli()
				for j := 0; j < len(dirs); j++ {
					if temp := baseDIR + "/" + dirs[j]; fileutil.IsExist(temp) {
						if mtime, err := fileutil.GetModifyTime(temp); nil == err {
							if nowTime-mtime.UnixMilli() > maxTime {
								fileutil.RemoveAll(temp)
							}
						}
					}
				}
			}
		} else {
			logs.Errorln("DoClearCache", err)
		}
		time.Sleep(time.Hour)
	}
}

// startDeleteCacheCleaner 启动'删除缓存'维护线程
func (locl *LocalDriver) startDeleteCacheCleaner(mtnode *ifiledatas.MountNode) {
	baseDIR := filepath.Clean(mtnode.Addr + "/" + deletingDir)
	logs.Infof("startDeleteCacheCleaner dir=%s\r\n", baseDIR)
	for {
		var err error
		var dirs []string
		if dirs, err = fileutil.GetDirList(baseDIR); nil == err {
			if len(dirs) > 0 {
				for j := 0; j < len(dirs); j++ {
					if temp := baseDIR + "/" + dirs[j]; fileutil.IsExist(temp) {
						fileutil.RemoveAll(temp)
					}
				}
			} else {
				time.Sleep(time.Minute)
			}
		} else {
			logs.Errorln("DoClearDeletedCahce", err)
		}
		time.Sleep(time.Second * 10)
	}
}
