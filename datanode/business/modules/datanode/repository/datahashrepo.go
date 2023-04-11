// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 数据存储节点-文件信息记录

package repository

import (
	"database/sql"
	"datanode/ifilestorage"
	"errors"
	"strings"
	"time"
)

const (
	status_hashdata_disabled = 0
	status_hashdata_enable   = 1
)

// NewDataHashRepo NewDataHashRepo
func NewDataHashRepo(table string, dbs ifilestorage.DBSetting) *DataHashRepo {
	dns := &DataHashRepo{table: table}
	dns.Initial(dbs)
	return dns
}

// DataHashRepo hahs记录
type DataHashRepo struct {
	table string
	db    *sql.DB
}

// Initial 初始化配置
func (ds *DataHashRepo) Initial(st ifilestorage.DBSetting) (err error) {
	if nil == ds.db {
		ds.db, err = sql.Open(st.DriverName, st.DataSourceName)
		if nil == err {
			if st.DriverName == "sqlite3" {
				// database is locked
				// https://github.com/mattn/go-sqlite3/issues/209
				ds.db.SetMaxOpenConns(1)
			} else {
				ds.db.SetMaxIdleConns(250)
				ds.db.SetConnMaxLifetime(time.Hour)
			}
		}
	}
	return err
}

// GetSqlTx 获取事物
func (ds *DataHashRepo) GetSqlTx() (*sql.Tx, error) {
	return ds.db.Begin()
}

// GetSqlTx 获取数据库对象
func (ds *DataHashRepo) GetDB() *sql.DB {
	return ds.db
}

// DisableInIds 设置节点为删除
func (ds *DataHashRepo) DisableInIds(conn *sql.Tx, ids []string) (err error) {
	if len(ids) == 0 {
		return errors.New("id is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("update " + ds.table + " set status=? where id in ('" + strings.Join(ids, "','") + "')")
	if nil == err {
		_, err = stmt.Exec(status_hashdata_disabled)
	}
	return err
}

// DisableInFIds 设置节点为删除
func (ds *DataHashRepo) DisableInFIds(conn *sql.Tx, fids []string) (err error) {
	if len(fids) == 0 {
		return errors.New("fid is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("update " + ds.table + " set status=? where fid in ('" + strings.Join(fids, "','") + "')")
	if nil == err {
		_, err = stmt.Exec(status_hashdata_disabled)
	}
	return err
}

// DeleteByFID 删除节点
func (ds *DataHashRepo) DeleteByFID(conn *sql.Tx, fid string) (err error) {
	if len(fid) == 0 {
		return errors.New("fid is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("delete from " + ds.table + " where fid=?")
	if nil == err {
		_, err = stmt.Exec(fid)
	}
	return err
}

// DeleteByID 删除节点
func (ds *DataHashRepo) DeleteByID(conn *sql.Tx, id string) (err error) {
	if len(id) == 0 {
		return errors.New("id is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("delete from " + ds.table + " where id=?")
	if nil == err {
		_, err = stmt.Exec(id)
	}
	return err
}

// DeleteInFIDs 删除节点
func (ds *DataHashRepo) DeleteInFIDs(conn *sql.Tx, fids []string) (err error) {
	if len(fids) == 0 {
		return errors.New("fids is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("delete from " + ds.table + " where fid in ('" + strings.Join(fids, "','") + "')")
	if nil == err {
		_, err = stmt.Exec()
	}
	return err
}

// DeleteInIDs 删除节点
func (ds *DataHashRepo) DeleteInIDs(conn *sql.Tx, ids []string) (err error) {
	if len(ids) == 0 {
		return errors.New("ids is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("delete from " + ds.table + " where id in ('" + strings.Join(ids, "','") + "')")
	if nil == err {
		_, err = stmt.Exec()
	}
	return err
}

// DeleteInHashs 删除节点-只能删除已被标记为删除的
func (ds *DataHashRepo) DeleteInHashs(conn *sql.Tx, hashs []string) (err error) {
	if len(hashs) == 0 {
		return errors.New("hashs is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("delete from " + ds.table + " where sha256 in ('" + strings.Join(hashs, "','") + "') and status=?")
	if nil == err {
		_, err = stmt.Exec(status_hashdata_disabled)
	}
	return err
}

// InsertEnabledHash 新增节点
func (ds *DataHashRepo) InsertEnabledHash(conn *sql.Tx, node ifilestorage.HNode) (err error) {
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("insert into " + ds.table + "(id, status, fid, sha256) values(?, ?, ?, ?)")
	if nil == err {
		_, err = stmt.Exec(node.Id, status_hashdata_enable, node.FId, node.Hash)
	}
	return err
}

// ListDisabled 获取文件pieces信息
func (ds *DataHashRepo) ListDisabled(conn *sql.DB) (nodes []ifilestorage.HNode, err error) {
	sqlstr := "select id, status, fid, sha256 from " + ds.table + " where status=?"
	var rows *sql.Rows
	if rows, err = conn.Query(sqlstr, status_hashdata_disabled); nil == err {
		defer rows.Close()
		for rows.Next() {
			node := ifilestorage.HNode{}
			if err = rows.Scan(&node.Id, &node.Status, &node.FId, &node.Hash); nil == err {
				nodes = append(nodes, node)
			} else {
				nodes = make([]ifilestorage.HNode, 0)
				break
			}
		}
	}
	return nodes, err
}

// QueryRepeatedHashAndDisabledIds 获取hash重复(disabled的和非disabled的数据), 并且已标记为删除的数据
func (ds *DataHashRepo) QueryRepeatedHashAndDisabledIds(conn *sql.DB) (ids []string, hashs []string, err error) {
	sqlstr := "select id, sha256 from ("
	sqlstr += "select tdel.id, tdel.sha256, t.id as jid from (select id, sha256 from " + ds.table + " where status = 0) as tdel "
	sqlstr += "left join (select id, sha256 from " + ds.table + " where sha256 in"
	sqlstr += "(select sha256 from " + ds.table + " where status = 0 group by sha256) and status = 1 group by sha256) as t "
	sqlstr += "on tdel.sha256=t.sha256) where jid notnull"
	var rows *sql.Rows
	if rows, err = conn.Query(sqlstr); nil == err {
		defer rows.Close()
		hashsmap := make(map[string]uint8)
		for rows.Next() {
			var id string
			var hash string
			if err = rows.Scan(&id, &hash); nil == err {
				ids = append(ids, id)
				if _, ok := hashsmap[hash]; !ok {
					hashs = append(hashs, hash)
				}
			} else {
				ids = make([]string, 0)
				hashs = make([]string, 0)
				break
			}
		}
	}
	return ids, hashs, err
}

// GetHash 在数据库中查找节点
func (ds *DataHashRepo) GetHash(conn *sql.DB, id string) (*ifilestorage.HNode, error) {
	if len(id) == 0 {
		return nil, errors.New("id is empty")
	}
	sql := "select id, status, fid, sha256 from " + ds.table + " where id=?"
	rows, err := conn.Query(sql, id)
	if nil == err {
		defer rows.Close()
		if rows.Next() {
			node := ifilestorage.HNode{}
			if err = rows.Scan(&node.Id, &node.Status, &node.FId, &node.Hash); nil == err {
				return &node, nil
			}
		}
	}
	return nil, err
}

// CreateTables 初始化结构
func (ds *DataHashRepo) CreateTables(conn *sql.Tx) (err error) {
	dropSql := "drop table  `" + ds.table + "`"
	tableSql := "create table if not exists `" + ds.table + "`  (`id` char(36)  not null, `status` int not null, `fid` char(36)  not null, `sha256` char(64) default '', primary key (`id`))"
	IndexSql := []string{
		"create index " + ds.table + "_index_id on " + ds.table + " (id)",
		"create index " + ds.table + "_index_fid on " + ds.table + " (fid)",
		"create index " + ds.table + "_index_status on " + ds.table + " (status)",
		"create index " + ds.table + "_index_sha256 on " + ds.table + " (sha256)",
	}

	conn.Exec(dropSql)
	//if nil == err {
	_, err = conn.Exec(tableSql)
	if nil == err {
		for i := 0; i < len(IndexSql); i++ {
			if _, err = conn.Exec(IndexSql[i]); nil != err {
				break
			}
		}
	}
	//}
	return err
}
