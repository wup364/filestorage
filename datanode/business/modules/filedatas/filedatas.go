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
	"datanode/business/modules/filedatas/dirmount"
	"datanode/business/modules/filedatas/fsdrivers"
	"datanode/business/modules/filedatas/ifiledatas"
	"datanode/ifilestorage"
	"errors"
	"io"
	"strings"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

// FileDatas 文件数据管理
type FileDatas struct {
	mt ifiledatas.DIRMount
	c  ipakku.AppConfig `@autowired:"AppConfig"`
}

// AsModule 作为一个模块
func (fns *FileDatas) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Name:        "FileDatas",
		Version:     1.0,
		Description: "数据存取模块",
		OnReady: func(mctx ipakku.Loader) {

		},
		OnSetup: func() {
			mount := fns.c.GetConfig(CONFKEY_MOUNT).ToStrMap(make(map[string]interface{}))
			if len(mount) == 0 {
				fns.c.SetConfig(CONFKEY_MOUNT+"./."+CONFKEY_MOUNTTYPE, "LOCAL")
				fns.c.SetConfig(CONFKEY_MOUNT+"./."+CONFKEY_MOUNTADDR, "./datas")
				logs.Infoln("Add mount: type=LOCAL, addr=./datas")
			}
		},
		OnInit: func() {
			mount := fns.c.GetConfig(CONFKEY_MOUNT).ToStrMap(make(map[string]interface{}))
			if len(mount) == 0 {
				logs.Panicln("Not find mount config, in config key: " + CONFKEY_MOUNT)
			}
			fns.mt = new(dirmount.DIRMount)
			// 注册支持的驱动
			fns.mt.RegisterFileDriver(new(fsdrivers.LocalDriver)).LoadAllMount(mount)
		},
	}
}

// getPathDriver getPathDriver
func (fns *FileDatas) getPathDriver(relativePath string) (ifiledatas.FileDriver, error) {
	if relativePath, err := checkPathSafety(relativePath); nil != err {
		return nil, err
	} else {
		return fns.mt.GetFileDriver(relativePath), nil
	}
}

// DoRename DoRename
func (fns *FileDatas) DoRename(relativePath, newName string) error {
	fs, err := fns.getPathDriver(relativePath)
	if nil != err {
		return err
	}
	return fs.DoRename(relativePath, newName)
}

// DoMkDir DoMkDir
func (fns *FileDatas) DoMkDir(relativePath string) error {
	fs, err := fns.getPathDriver(relativePath)
	if nil != err {
		return err
	}
	return fs.DoMkDir(relativePath)
}

// DoDelete 删除文件|文件夹
func (fns *FileDatas) DoDelete(relativePath string) error {
	fs, err := fns.getPathDriver(relativePath)
	if nil != err {
		return err
	}
	return fs.DoDelete(relativePath)
}

// DoMove 移动文件|文件夹
func (fns *FileDatas) DoMove(src, dst string, replace bool) error {
	fs, err := fns.getPathDriver(src)
	if nil != err {
		return err
	}
	dst, err = checkPathSafety(dst)
	if nil != err {
		return err
	}
	return fs.DoMove(src, dst, replace)
}

// DoCopy 复制文件|夹
func (fns *FileDatas) DoCopy(src, dst string, replace bool) error {
	fs, err := fns.getPathDriver(src)
	if nil != err {
		return err
	}
	dst, err = checkPathSafety(dst)
	if nil != err {
		return err
	}
	return fs.DoCopy(src, dst, replace)
}

// DoWrite 写入文件
func (fns *FileDatas) DoWrite(relativePath string, ioReader io.Reader) error {
	fs, err := fns.getPathDriver(relativePath)
	if nil != err {
		return err
	}
	return fs.DoWrite(relativePath, ioReader)
}

// DoRead 读取文件
func (fns *FileDatas) DoRead(relativePath string, offset int64) (io.ReadCloser, error) {
	fs, err := fns.getPathDriver(relativePath)
	if nil != err {
		return nil, err
	}
	return fs.DoRead(relativePath, offset)
}

// IsFile 是否是文件, 如果路径不对或者驱动不对则为 false
func (fns *FileDatas) IsFile(relativePath string) bool {
	fs, err := fns.getPathDriver(relativePath)
	if nil != err {
		return false
	}
	return fs.IsFile(relativePath)
}

// IsDir 是否是文件夹子·, 如果路径不对或者驱动不对则为 false
func (fns *FileDatas) IsDir(relativePath string) bool {
	fs, err := fns.getPathDriver(relativePath)
	if nil != err {
		return false
	}
	return fs.IsDir(relativePath)
}

// IsExist 是否存在, 如果路径不对或者驱动不对则为 false
func (fns *FileDatas) IsExist(relativePath string) bool {
	fs, err := fns.getPathDriver(relativePath)
	if nil != err {
		return false
	}
	return fs.IsExist(relativePath)
}

// GetFileSize 获取文件大小
func (fns *FileDatas) GetFileSize(relativePath string) int64 {
	fs, err := fns.getPathDriver(relativePath)
	if nil != err {
		return -1
	}
	return fs.GetFileSize(relativePath)
}

// GetDirList 获取文件夹列表
func (fns *FileDatas) GetDirList(relativePath string, limit int, offset int) []string {
	fs, err := fns.getPathDriver(relativePath)
	if nil != err {
		return make([]string, 0)
	}
	return fs.GetDirList(relativePath, limit, offset)
}

// GetNode 获取文件的基本信息
func (fns *FileDatas) GetNode(relativePath string) *ifilestorage.FNode {
	fs, err := fns.getPathDriver(relativePath)
	if nil != err {
		return nil
	}
	if node := fs.GetNode(relativePath); nil != node {
		return &ifilestorage.FNode{
			Path:   node.Path,
			IsDir:  node.IsDir,
			IsFile: node.IsFile,
			Mtime:  node.Mtime,
			Size:   node.Size,
		}
	}
	return nil
}

// GetNodes 批量获取文件基本信息
func (fns *FileDatas) GetNodes(src []string, ignoreNotIsExist bool) (res []ifilestorage.FNode, err error) {
	// 分组
	if group, err := fns.groupPathByDriver(src); nil == err && len(group) > 0 {
		merge := make([]ifiledatas.Node, 0)
		for _, paths := range group {
			if fs, err := fns.getPathDriver(paths[0]); nil == err {
				if nodes, err := fs.GetNodes(paths, ignoreNotIsExist); nil != err {
					return res, err
				} else if len(nodes) > 0 {
					merge = append(merge, nodes...)
				}
			} else {
				return res, err
			}
		}
		// dto
		res = make([]ifilestorage.FNode, len(merge))
		for i := 0; i < len(merge); i++ {
			res[i] = ifilestorage.FNode{
				Path:   merge[i].Path,
				IsDir:  merge[i].IsDir,
				IsFile: merge[i].IsFile,
				Mtime:  merge[i].Mtime,
				Size:   merge[i].Size,
			}
		}
	} else {
		return res, err
	}
	return res, err
}

// GetDirNodeList 获取文件夹下文件的基本信息
func (fns *FileDatas) GetDirNodeList(relativePath string, limit int, offset int) (res []ifilestorage.FNode, err error) {
	if fs, err := fns.getPathDriver(relativePath); nil != err {
		return res, err
	} else {
		if nodes, err := fs.GetDirNodeList(relativePath, limit, offset); nil != err {
			return res, err
		} else {
			// 挂载目录
			if mLS := fns.mt.ListChildMount(relativePath); len(mLS) > 0 {
				for i := 0; i < len(mLS); i++ {
					existedIndex := -1
					for j := 0; j < len(nodes); j++ {
						if nodes[j].Path == mLS[i] {
							existedIndex = j
							break
						}
					}
					//
					childPath := mLS[i]
					if relativePath != "/" {
						childPath = relativePath + childPath
					}
					mtfs := fns.mt.GetFileDriver(childPath)
					if node := mtfs.GetNode(childPath); nil != node {
						if existedIndex == -1 {
							nodes = append(nodes, *node)
						} else {
							nodes[existedIndex] = *node
						}
					}
				}
			}
			// 分组 + 合并
			if len(nodes) > 0 {
				files := make([]ifilestorage.FNode, 0)
				folders := make([]ifilestorage.FNode, 0)
				for i := 0; i < len(nodes); i++ {
					node := ifilestorage.FNode{
						Path:   nodes[i].Path,
						IsDir:  nodes[i].IsDir,
						IsFile: nodes[i].IsFile,
						Mtime:  nodes[i].Mtime,
						Size:   nodes[i].Size,
					}
					if node.IsFile {
						files = append(files, node)
					} else {
						folders = append(folders, node)
					}
				}
				// 合并
				res = append(res, folders...)
				res = append(res, files...)
			}
		}
	}
	return res, err
}

// groupPathByDriver 根据驱动类型分类
func (fns *FileDatas) groupPathByDriver(srcList []string) (map[string][]string, error) {
	result := make(map[string][]string)
	for i := 0; i < len(srcList); i++ {
		if relativePath, err := checkPathSafety(srcList[i]); nil != err {
			return nil, err
		} else if mnode := fns.mt.GetMountNode(relativePath); nil != mnode {
			if val, ok := result[mnode.Type]; ok {
				result[mnode.Type] = append(val, relativePath)
			} else {
				result[mnode.Type] = []string{relativePath}
			}
		}
	}
	return result, nil
}

// checkPathSafety 路径合规检查, 避免 ../ ./之类的路径
func checkPathSafety(path string) (string, error) {
	if len(path) == 0 {
		return "/", nil
	}
	path = strings.Replace(path, "\\", "/", -1)
	if strings.Contains(path, "../") {
		return "", errors.New("unsupported path format '../'")
	}
	if strings.Contains(path, "./") {
		return "", errors.New("unsupported path format './'")
	}
	return strutil.Parse2UnixPath(path), nil
}
