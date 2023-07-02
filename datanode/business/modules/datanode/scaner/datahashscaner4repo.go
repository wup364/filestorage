// Copyright (C) 2023 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package scaner

import (
	"database/sql"
	"datanode/business/modules/datanode/fio"
	"datanode/business/modules/datanode/repository"
	"datanode/ifilestorage"
	"time"

	"github.com/wup364/pakku/utils/logs"
)

// NewDataHashCanerForRepo 从数据库中读取hash, 并验证文件是否还存在
func NewDataHashCanerForRepo(fds ifilestorage.FileDatas, dhs *repository.DataHashRepo) Scaner {
	return &DataHashCanerForRepo{
		fds: fds,
		dhs: dhs,
	}
}

// DataHashCanerForRepo 从数据库中读取hash, 并验证文件是否还存在
type DataHashCanerForRepo struct {
	started            bool
	smallestScanmarker int
	fds                ifilestorage.FileDatas
	dhs                *repository.DataHashRepo
}

// StartScaner 开始扫描hash记录
func (dhc *DataHashCanerForRepo) StartScaner() {
	if dhc.started {
		return
	}
	dhc.started = true

	defer func() {
		if err := recover(); err != nil {
			logs.Errorln("DataHashCanerForRepo [recover]: ", err)
			dhc.started = false
			dhc.StartScaner()
		}
	}()

	if err := dhc.doInitScan(); nil != err {
		logs.Panicln(err)
		time.Sleep(time.Minute * 10)
	}

	logs.Infoln("DataHashCanerForRepo StartScaner")
	for {
		if size, err := dhc.doLimitScan(100); nil != err {
			logs.Errorf("doLimitScan: err=%s \r\n", err.Error())
			time.Sleep(time.Minute)

		} else if size == 0 {
			dhc.started = false
			break
		}
		time.Sleep(time.Millisecond * 200)
	}
	logs.Infoln("DataHashCanerForRepo scan complete")

	time.Sleep(time.Hour * 24)
	go dhc.StartScaner()
}

func (dhc *DataHashCanerForRepo) doInitScan() (err error) {
	dhc.smallestScanmarker, err = dhc.dhs.GetSmallestScanmarker(dhc.dhs.GetDB())
	return
}

// doLimitScan 分批扫描
func (dhc *DataHashCanerForRepo) doLimitScan(limit int) (size int, err error) {
	var hashs []string
	if hashs, err = dhc.dhs.ListEnabledHashByScanmarker(dhc.dhs.GetDB(), dhc.smallestScanmarker, limit, 0); nil != err || len(hashs) == 0 {
		return
	}

	notFound := make([]string, 0)
	isExisted := make([]string, 0)
	for i := 0; i < len(hashs); i++ {
		if temp := fio.GetArchivedPath4Hash(hashs[i]); dhc.fds.IsFile(temp) {
			isExisted = append(isExisted, hashs[i])
		} else {
			notFound = append(notFound, hashs[i])
		}
	}
	if len(isExisted) > 0 {
		if err = dhc.updateIsExistedScanMarkerBySha256(isExisted); nil != err {
			return
		}
	}

	if len(notFound) > 0 {
		logs.Infof("doLimitScan notFound=%s", notFound)
		if err = dhc.markerNotFoundBySha256(notFound); nil != err {
			return
		}
	}

	return len(hashs), err
}

// updateIsExistedScanMarkerBySha256 更新已存在的hash的扫描标记
func (dhc *DataHashCanerForRepo) updateIsExistedScanMarkerBySha256(hashs []string) (err error) {
	var tx *sql.Tx
	if tx, err = dhc.dhs.GetSqlTx(); nil == err {
		if err = dhc.dhs.UpdateScanmarkerBySha256(tx, hashs, dhc.smallestScanmarker+1); nil == err {
			err = tx.Commit()
		} else {
			tx.Rollback()
		}
	}
	return
}

// markerNotFoundBySha256 标记为数据丢失
func (dhc *DataHashCanerForRepo) markerNotFoundBySha256(hashs []string) (err error) {
	var tx *sql.Tx
	if tx, err = dhc.dhs.GetSqlTx(); nil == err {
		if err = dhc.dhs.UpdateScanmarker2NotFoundBySha256(tx, hashs); nil == err {
			err = tx.Commit()
		} else {
			tx.Rollback()
		}
	}
	return
}
