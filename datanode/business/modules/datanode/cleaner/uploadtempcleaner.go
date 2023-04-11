// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 上传缓存清理程序

package cleaner

import (
	"datanode/business/modules/datanode/fio"
	"datanode/ifilestorage"
	"time"

	"github.com/wup364/pakku/utils/logs"
)

// NewUploadTempCleaner 上传缓存清理程序
func NewUploadTempCleaner(fds ifilestorage.FileDatas) *UploadTempCleaner {
	return &UploadTempCleaner{fds: fds}
}

// UploadTempCleaner 上传缓存清理程序
type UploadTempCleaner struct {
	fds     ifilestorage.FileDatas
	started bool
}

// StartCleaner 清理删除数据开始
func (cls *UploadTempCleaner) StartCleaner() {
	if cls.started {
		return
	}
	cls.started = true
	cls.startUploadTempCleaner()
}

// startUploadTempCleaner 启动'temp/upload文件'维护线程
func (cls *UploadTempCleaner) startUploadTempCleaner() {
	maxTime := int64(1 * 24 * 60 * 60 * 1000)
	baseDIR := fio.GetTempDIR4Upload("")
	logs.Infof("startUploadTempCleaner path=%s, exp=%d\r\n", baseDIR, maxTime)
	for {
		if dirs := cls.fds.GetDirList(baseDIR, -1, -1); len(dirs) > 0 {
			nowTime := time.Now().UnixMilli()
			for j := 0; j < len(dirs); j++ {
				if temp := baseDIR + "/" + dirs[j]; cls.fds.IsExist(temp) {
					if fnode := cls.fds.GetNode(temp); nil != fnode {
						if nowTime-fnode.Mtime > maxTime {
							cls.fds.DoDelete(temp)
						}
					}
				}
			}
		}
		time.Sleep(time.Hour)
	}
}
