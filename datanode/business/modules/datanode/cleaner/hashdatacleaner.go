// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 1. 入库filedatas时同时写入hashdatas表, 按照hash、文件id的方式不去重写入
// 2. 删除文件时, 根据文件id操作hash统计表, 状态设置为删除. 并且同时删除filedatas表的相关数据
// 3. 查询hashdatas status=0的数据, 得到hash值, 去重然后锁定. 当其中部分hash被锁定时跳过他, 但是要记录下来
// 4. 左连接查询出标记为删除, 但是还有引用的数据, 删除这些数据, 并解锁部分hash
// 5. 剩下的除掉锁定失败的, 就是没有引用的hash, 直接删除文件.

package cleaner

import (
	"database/sql"
	"datanode/business/modules/datanode/fio"
	"datanode/business/modules/datanode/repository"
	"datanode/ifilestorage"
	"time"

	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

// NewHashDataCleaner NewHashDataCleaner
func NewHashDataCleaner(nodeno string,
	dhc *fio.HashDataCtrl,
	dns *repository.DataNodeRepo,
	dhs *repository.DataHashRepo,
	fds ifilestorage.FileDatas,
	nrpc ifilestorage.RPC4NameNode,
	delFile bool) Cleaner {
	return &HashDataClear{
		nodeno:  nodeno,
		dhc:     dhc,
		dns:     dns,
		dhs:     dhs,
		fds:     fds,
		nrpc:    nrpc,
		delFile: delFile,
	}
}

// HashDataClear Hash文件数据清理
type HashDataClear struct {
	nodeno  string
	dhc     *fio.HashDataCtrl
	dns     *repository.DataNodeRepo
	dhs     *repository.DataHashRepo
	fds     ifilestorage.FileDatas
	nrpc    ifilestorage.RPC4NameNode
	delFile bool
	started bool
}

// StartCleaner 清理删除数据开始
func (cls *HashDataClear) StartCleaner() {
	if cls.started {
		return
	}
	cls.started = true
	defer func() {
		if err := recover(); err != nil {
			logs.Errorln("HashDataClear [recover]: ", err)
			cls.started = false
			cls.StartCleaner()
		}
	}()
	limit := 200
	for {
		var err error
		var addrs []ifilestorage.DeletedAddr
		// 查询删除列表
		if addrs, err = cls.nrpc.DoQueryDeletedDataAddr(cls.nodeno, limit); nil == err && len(addrs) > 0 {
			fids := make([]string, len(addrs))
			dids := make([]string, len(addrs))
			for i := 0; i < len(addrs); i++ {
				dids[i] = addrs[i].Id
				fids[i] = addrs[i].Fid
			}

			optid := strutil.GetUUID()
			logs.Infof("QueryDeletedDataAddr optid=%s, list=%s \r\n", optid, dids)
			if err = cls.doClearHashRecords(optid, fids); nil == err && cls.delFile {
				err = cls.doClearHashFile()
			}
			if nil == err {
				logs.Infof("ConfirmDeletedDataAddr optid=%s\r\n", optid)
				err = cls.nrpc.DoConfirmDeletedDataAddr(dids)
			}
		}
		if nil != err {
			time.Sleep(time.Minute)
		} else if len(addrs) < limit {
			time.Sleep(time.Minute * 10)
		} else {
			time.Sleep(time.Second)
		}
	}
}

// doClearHashRecords 清理数据库数据
// . 根据文件fid删除 datanodes 表的相关数据
// . 根据文件fid操作 datahashs 表, 状态设置为删除
// . 左连接查询出标记为删除, 但是还有引用的数据, 删除这些数据
func (cls *HashDataClear) doClearHashRecords(optid string, fids []string) (err error) {
	// . 根据文件fid删除 datanodes 表的相关数据
	if err = cls.deleteDataNodeRecords(fids); nil != err {
		logs.Errorf("deleteDataNodeRecords optid=%s, err=%s \r\n", optid, err.Error())
		return // 删不掉不要紧, 只要没提交到namenode就会再次尝试
	}

	// . 根据文件fid操作 datahashs 表, 状态设置为删除
	if err = cls.disableDatahashsRecord(fids); nil != err {
		logs.Errorf("disableDatahashsRecord optid=%s, err=%s \r\n", optid, err.Error())
		return // 删不掉不要紧, 只要没提交到namenode就会再次尝试
	}

	// . 左连接查询出标记为删除, 但是还有引用的数据, 删除这些数据
	if _, err = cls.deleteRepeatedAndDisabledHashDataRecords(); nil != err {
		logs.Errorf("deleteRepeatedAndDisabledHashDataRecords optid=%s, err=%s \r\n", optid, err.Error())
		return // 删不掉不要紧, 只要没提交到namenode就会再次尝试
	}
	return
}

// deleteDataNodeRecords 根据文件fid删除 datanodes 表的相关数据
func (cls *HashDataClear) deleteDataNodeRecords(fids []string) (err error) {
	var tx *sql.Tx
	if tx, err = cls.dns.GetSqlTx(); nil == err {
		if err = cls.dns.DeleteInIDs(tx, fids); nil == err {
			err = tx.Commit()
		} else {
			tx.Rollback()
		}
	}
	return
}

// disableDatahashsRecord 根据文件fid操作 datahashs 表, 状态设置为删除
func (cls *HashDataClear) disableDatahashsRecord(fids []string) (err error) {
	var tx *sql.Tx
	if tx, err = cls.dhs.GetSqlTx(); nil == err {
		if err = cls.dhs.DisableInFIds(tx, fids); nil == err {
			err = tx.Commit()
		} else {
			tx.Rollback()
		}
	}
	return
}

// deleteRepeatedAndDisabledHashDataRecords 左连接查询出标记为删除, 但是还有引用的数据, 删除这些数据
func (cls *HashDataClear) deleteRepeatedAndDisabledHashDataRecords() (hashs []string, err error) {
	var tx *sql.Tx
	var disIds []string
	if disIds, hashs, err = cls.dhs.QueryRepeatedHashAndDisabledIds(cls.dhs.GetDB()); nil == err && len(disIds) > 0 {
		if tx, err = cls.dhs.GetSqlTx(); nil == err {
			if err = cls.dhs.DeleteInIDs(tx, disIds); nil == err {
				err = tx.Commit()
			} else {
				tx.Rollback()
			}
		}
	}
	return
}

// doClearHashFile 清理冗余的hash文件
// . 查询 datahashs status=0 的数据, 得到hash值, 去重然后锁定.  当其中部分hash被其他线程锁定时跳过他
// . 再次执行, 左连接查询出标记为删除, 但是还有引用的数据, 删除这些数据, 并解锁这部分hash
// . 剩下的锁定成功的, 就是没有引用的hash, 直接删除文件并解锁
// . 循环锁定失败的部分, 执行上述动作
func (cls *HashDataClear) doClearHashFile() (err error) {
	var disabledNodes []ifilestorage.HNode
	// . 查询 datahashs status=0 的数据, 得到hash值, 去重然后锁定.  当其中部分hash被其他线程锁定时跳过他
	if disabledNodes, err = cls.dhs.ListDisabled(cls.dhs.GetDB()); nil != err || len(disabledNodes) == 0 {
		return
	}

	// 初步标记即将删除的hash
	needlockhash := make([]string, 0)
	allhashMap := make(map[string]uint8)
	for i := 0; i < len(disabledNodes); i++ {
		if _, ok := allhashMap[disabledNodes[i].Hash]; !ok {
			allhashMap[disabledNodes[i].Hash] = 0
			needlockhash = append(needlockhash, disabledNodes[i].Hash)
		}
	}
	markedhash := cls.dhc.MarkHashOnDelete(needlockhash)
	defer func() {
		for i := 0; i < len(markedhash); i++ {
			cls.dhc.ReleaseDeleteLock(markedhash[i])
		}
	}()

	// . 左连接查询出标记为删除, 但是还有引用的数据, 删除这些数据, 并解锁这部分hash
	if hashs, err := cls.deleteRepeatedAndDisabledHashDataRecords(); nil != err {
		return err
	} else if len(hashs) > 0 {
		for i := 0; i < len(hashs); i++ {
			cls.dhc.RemoveHashDeleteMark(hashs[i])
		}
	}

	// . 剩下的锁定成功的, 就是没有引用的hash, 直接删除文件并解锁
	if candelhash := cls.tryDeleteAndUnlokHashFile(markedhash); len(candelhash) > 0 {
		// 删除这些已经删除的hash
		err = cls.deleteDataHashRecord(candelhash)
	}
	// . 循环锁定失败的部分, 执行上述动作 - //
	return err
}

// deleteDataHashRecord 删除hash片记录
func (cls *HashDataClear) deleteDataHashRecord(hashs []string) (err error) {
	var tx *sql.Tx
	if tx, err = cls.dhs.GetSqlTx(); nil == err {
		if err = cls.dhs.DeleteInHashs(tx, hashs); nil == err {
			err = tx.Commit()
		} else {
			tx.Rollback()
		}
	}
	return
}

// tryDeleteAndUnlokHashFile 尝试删除hash文件, 并解锁, 返回删除成功的列表
func (cls *HashDataClear) tryDeleteAndUnlokHashFile(hashs []string) (succeed []string) {
	succeed = make([]string, 0)
	for i := 0; i < len(hashs); i++ {
		if !cls.dhc.LockHashOnDelete(hashs[i]) {
			continue
		}
		if temp := fio.GetArchivedPath4Hash(hashs[i]); cls.fds.IsFile(temp) {
			if err := cls.fds.DoDelete(temp); nil == err {
				succeed = append(succeed, hashs[i])
			} else {
				logs.Errorf("tryDeleteAndUnlokHashFile path=%s, err=%s \r\n", temp, err.Error())
			}
		}
		cls.dhc.ReleaseDeleteLock(hashs[i])
	}
	return
}
