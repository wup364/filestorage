// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 数据存储节点-文件归档
package datanode

import (
	"datanode/ifilestorage"
	"errors"
	"io"
	"strconv"
	"time"

	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

const (
	DIRBase_Temp      = "/temp"
	DIRBase_Archived  = "/archived"
	DIRBase_Deleteing = "/deleteing"
	tokenexp          = 60 * 30 // 每一片30分钟
)

// NewDataNodeFiles 实例一个文件块归档管理
func NewDataNodeFiles(fds ifilestorage.FileDatas, dhc *HashDataCtrl) *DataNodeFiles {
	return &DataNodeFiles{
		fds: fds,
		dhc: dhc,
	}
}

// DataNodeFiles 数据存储管理
type DataNodeFiles struct {
	dhc *HashDataCtrl
	fds ifilestorage.FileDatas
}

// DoWriteStream 通过token计算写入位置, 并记录写入信息
func (dnm *DataNodeFiles) DoWriteStream(opts *ifilestorage.StreamWriteOpts, reader io.Reader) (err error) {
	var wpath string
	if wpath, err = getTempDIR4UploadToken(*opts); nil == err {
		var newReader *WriteReader
		if newReader, err = dnm.dhc.GetWriteReader(opts, reader); nil == err {
			defer newReader.Close()
			// 只要被锁定过, 那么这个文件很有可能正在被删除, 保险起见还是重新上传一遍
			if len(opts.Sha256) > 0 && dnm.fds.IsFile(getArchivedPath4Hash(opts.Sha256)) && dnm.dhc.CanQuoteHash(opts.Sha256) {
				// 这个文件已经有了, 不需要再传递了[信任传递的sha256情况]
				newReader.SetHash(opts.Sha256)
			} else {
				if !dnm.fds.IsExist(wpath) {
					err = dnm.fds.DoMkDir(wpath)
				}
				if nil == err {
					tmpPath := wpath + "/" + strconv.Itoa(opts.PieceNumber) + ".tmp"
					if err = dnm.fds.DoWrite(tmpPath, newReader); nil == err {
						if vhash := newReader.GetHash(); len(opts.Sha256) > 0 && opts.Sha256 != vhash {
							err = errors.New("hash verification failed, " + opts.Sha256 + " != " + vhash)
						} else {
							opts.Sha256 = vhash
							err = dnm.fds.DoMove(tmpPath, wpath+"/"+vhash, true)
						}
					}
				}
			}
		}
	}
	return err
}

// DoReadStream DoReadStream
func (dnm *DataNodeFiles) DoReadStream(opts *ifilestorage.StreamReadOpts, offset int64) (io.Reader, error) {
	r, err := dnm.dhc.GetReader(*opts)
	if nil == err {
		r.OnPiecesReader(func(hash string, offset int64) (io.Reader, error) {
			path := getArchivedPath4Hash(hash)
			if !dnm.fds.IsFile(path) {
				return nil, fileutil.PathNotExist("PiecesReader", path)
			}
			return dnm.fds.DoRead(path, offset)
		})
		r.OnPiecesSize(func(hash string) (int64, error) {
			path := getArchivedPath4Hash(hash)
			if !dnm.fds.IsFile(path) {
				return 0, fileutil.PathNotExist("PiecesReader", path)
			}
			return dnm.fds.GetFileSize(path), nil
		})
		if opts.Offset > 0 {
			r.SetOffset(opts.Offset)
		}
	}
	return r, err
}

// DoDestroyToken 销毁token
func (dnm *DataNodeFiles) DoDestroyToken(token string) {
	dnm.dhc.wtokens.DestroyToken(token)
}

// DoArchiveToken 提交token, 表示写入完成
func (dnm *DataNodeFiles) DoArchiveToken(token string) (fnode *ifilestorage.DNode, err error) {
	// 检查所有文件是否都上来了
	var wopts []ifilestorage.StreamWriteOpts
	if wopts, err = dnm.dhc.GetPiecesAndSucceed(token); nil == err {
		res := ifilestorage.DNode{
			Size:   0,
			Pieces: make([]string, len(wopts)),
		}
		// 把文件归档 hash 位置
		for i := 0; i < len(wopts); i++ {
			src, _ := getTempDIR4UploadToken(wopts[i])
			dst := getArchivedPath4Hash(wopts[i].Sha256)
			// 等待可以归档hash文件, 如果在删除队列中, 则需要等待
			dnm.dhc.LockHashOnArchive(wopts[i].Sha256)
			if !dnm.fds.IsFile(dst) {
				tsrc := src + "/" + wopts[i].Sha256
				if err = dnm.fds.DoMove(tsrc, dst, false); nil != err {
					if !dnm.fds.IsFile(dst) {
						for i := 0; i < 10; i++ {
							time.Sleep(time.Second)
							if dnm.fds.IsFile(dst) {
								err = nil
								break
							}
							if err = dnm.fds.DoMove(tsrc, dst, false); nil == err {
								break
							}
							logs.Errorln("DoArchiveToken.DoMove.Retry", err)
						}
					} else if dnm.fds.IsFile(dst) {
						err = nil
					}
				}
			}
			if nil != err {
				// TODO 失败了需要回滚, 这种情况视为异常, 先记录错误
				logs.Errorln("DoArchiveToken.DoMove", err)
				return fnode, err
			}
			//
			res.Pieces[i] = wopts[i].Sha256
			res.Size += dnm.fds.GetFileSize(dst)
		}
		if len(res.Pieces) == 1 {
			res.Hash = res.Pieces[0]
		}
		if temp := getTempDIR4Upload(token); dnm.fds.IsExist(temp) {
			if err = dnm.fds.DoDelete(temp); nil != err {
				for i := 0; i < 10; i++ {
					if !dnm.fds.IsExist(temp) {
						err = nil
						break
					}
					if err = dnm.fds.DoDelete(temp); nil == err {
						break
					}
					logs.Errorln("DoArchiveToken.DoDelete.Retry", err)
				}
			}
		}
		if nil == err {
			fnode = &res
		}
	}
	return fnode, err
}

// startUploadTempCleaner 启动'temp/upload文件'维护线程
func (dnm *DataNodeFiles) startUploadTempCleaner() {
	maxTime := int64(1 * 24 * 60 * 60 * 1000)
	baseDIR := getTempDIR4Upload("")
	logs.Infof("startUploadTempCleaner path=%s, exp=%d\r\n", baseDIR, maxTime)
	for {
		if dirs := dnm.fds.GetDirList(baseDIR, -1, -1); len(dirs) > 0 {
			nowTime := time.Now().UnixMilli()
			for j := 0; j < len(dirs); j++ {
				if temp := baseDIR + "/" + dirs[j]; dnm.fds.IsExist(temp) {
					if fnode := dnm.fds.GetNode(temp); nil != fnode {
						if nowTime-fnode.Mtime > maxTime {
							dnm.fds.DoDelete(temp)
						}
					}
				}
			}
		}
		time.Sleep(time.Hour)
	}
}

// getArchivedPath4Hash 分片数据归档路径
func getArchivedPath4Hash(sha256 string) string {
	if len(sha256) == 64 {
		// e/0d/afb6109ade/198327e54c/04b9e92ba9/25f29292f3/16210f4a98/e0dafb6109ade198327e54c04b9e92ba925f29292f316210f4a988c0851ea9b8
		return getArchivedDIR4Hash(sha256[0:1] + "/" + sha256[1:3] + "/" + sha256[3:13] + "/" + sha256[13:23] + "/" + sha256[23:33] + "/" + sha256[33:43] + "/" + sha256[43:53] + "/" + sha256)
	}
	return ""
}

// getTempDIR4UploadToken 上传零时存放目录, // temp/upload/tokenxxx/1
func getTempDIR4UploadToken(opts ifilestorage.StreamWriteOpts) (string, error) {
	if len(opts.Token) > 0 {
		// temp/upload/tokenxxx/1
		return getTempDIR4Upload(opts.Token + "/" + strconv.Itoa(opts.PieceNumber)), nil
	}
	return "", ErrInvalidToken
}

// getArchivedDIR4Hash 分片的存储位置, 存放以hash为唯一依据的文件
func getArchivedDIR4Hash(path string) string {
	return strutil.Parse2UnixPath(DIRBase_Archived + "/" + path)
}

// getTempDIR4Upload 上传临时目录路径
func getTempDIR4Upload(path string) string {
	return strutil.Parse2UnixPath(DIRBase_Temp + "/upload/" + path)
}
