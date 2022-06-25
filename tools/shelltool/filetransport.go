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
	"errors"
	"io"
	"opensdk"
	"path/filepath"

	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

// doUploadDir 上传文件夹
func doUploadDir(dest, loaclBase, localScan string, override bool, o opensdk.IOpenApi) error {
	baseLen := len(loaclBase)
	if dirs, err := fileutil.GetDirList(localScan); nil == err && len(dirs) > 0 {
		for _, val := range dirs {
			localFile := filepath.Clean(localScan + "/" + val)
			if fileutil.IsDir(localFile) {
				doUploadDir(dest, loaclBase, localFile, override, o)
			} else {
				uploadPath := strutil.Parse2UnixPath(dest + "/" + (localFile[baseLen:]))
				if !override {
					if ok, err := o.IsExist(uploadPath); nil == err && ok {
						Println("Skip existing", uploadPath)
						logs.Infoln("Skip existing", uploadPath)
						continue
					} else if nil != err {
						return err
					}
				}
				var err error
				var token *opensdk.StreamToken
				if token, err = o.DoAskWriteToken(uploadPath); nil == err {
					if err = opensdk.FileUploader(localFile, token, 128*1024*1024, o.DoWriteToken); nil == err {
						if _, err = o.DoSubmitWriteToken(token.Token, override); nil == err {
							Println("Uploaded successfully", uploadPath)
							logs.Infoln("Uploaded successfully", uploadPath)
						}
					}
				}
				if nil != err {
					return err
				}
			}
		}
	} else {
		return err
	}
	return nil
}

// doDownloadDir 下载文件夹
func doDownloadDir(locl, base, src string, override bool, o opensdk.IOpenApi) (err error) {
	baseLen := len(base)
	offset := 0
	var nodes *opensdk.DirNodeListDto
	for nil == err {
		if nodes, err = o.GetDirNodeList(src, 100, offset); nil != err {
			break
		}
		if nil == nodes || nodes.Total == 0 || len(nodes.Datas) == 0 {
			break
		}
		for i := 0; i < len(nodes.Datas); i++ {
			node := nodes.Datas[i]
			if node.Flag == 0 {
				if err = doDownloadDir(locl, base, strutil.Parse2UnixPath(src+"/"+node.Name), override, o); nil != err {
					break
				}
			} else {
				src := strutil.Parse2UnixPath(src + "/" + node.Name)
				dest := filepath.Clean(locl + "/" + src[baseLen:])
				Println("Downloaded successfully", src)
				logs.Infoln("download src=" + src + ", dest=" + dest)
				if err = doDownloadFile(dest, src, override, o); nil != err {
					break
				}
			}
		}
		offset += 100
		if nil != nodes && nodes.Total <= offset {
			break
		}
	}
	if nil != err {
		logs.Errorln(err)
	}
	return err
}

// doDownloadFile 下载文件
func doDownloadFile(locl, src string, override bool, o opensdk.IOpenApi) error {
	if fileutil.IsExist(locl) {
		if !override {
			return fileutil.PathExist("download", locl)
		}
		if err := fileutil.RemoveFile(locl); nil != err {
			return err
		}
	} else if parent := strutil.GetPathParent(locl); !fileutil.IsExist(parent) {
		if err := fileutil.MkdirAll(parent); nil != err {
			return err
		}
	}
	if token, err := o.DoAskReadToken(src); nil != err {
		return err
	} else {
		if r, err := o.DoReadToken(token.NodeNo, token.Token, token.EndPoint, 0); nil != err {
			return err
		} else {
			defer r.Close()
			if w, err := fileutil.GetWriter(locl); nil == err {
				defer w.Close()
				if _, err := io.Copy(w, r); nil != err {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}

// copyNodes 复制文件|文件夹
func copyNodes(src, dest string, o opensdk.IOpenApi) (err error) {
	if len(src) == 0 || len(dest) == 0 {
		return errors.New("src or dest is nil")
	}
	var srcNode *opensdk.TNode
	if srcNode, err = o.GetNode(src); nil == err {
		if srcNode.Flag == 0 {
			offset := 0
			var nodes *opensdk.DirNodeListDto
			for nil == err {
				if nodes, err = o.GetDirNodeList(src, 100, offset); nil != err {
					break
				}
				if nil == nodes || nodes.Total == 0 || len(nodes.Datas) == 0 {
					break
				}
				for i := 0; i < len(nodes.Datas); i++ {
					node := nodes.Datas[i]
					if node.Flag == 0 {
						if err = copyNodes(strutil.Parse2UnixPath(src+"/"+node.Name), strutil.Parse2UnixPath(dest+"/"+node.Name), o); nil != err {
							break
						}
					} else {
						src := strutil.Parse2UnixPath(src + "/" + node.Name)
						dest := strutil.Parse2UnixPath(dest + "/" + node.Name)
						Printfln("Copy node src=%s, dest=%s", src, dest)
						if _, err = o.DoCopy(src, dest, false); nil != err {
							break
						}
					}
				}
				offset += 100
				if nil != nodes && nodes.Total <= offset {
					break
				}
			}
		} else {
			Printfln("Copy node src=%s, dest=%s", src, dest)
			_, err = o.DoCopy(src, dest, false)
		}
	}
	if nil != err {
		logs.Errorln(err)
	}
	return err
}
