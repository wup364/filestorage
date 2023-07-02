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

// NewDataNodeRepo NewDataNodeRepo
func NewDataNodeRepo(table string, dbs ifilestorage.DBSetting) *DataNodeRepo {
	dns := &DataNodeRepo{table: table}
	dns.Initial(dbs)
	return dns
}

// DataNodeRepo 元数据信息-数据库
type DataNodeRepo struct {
	table string
	db    *sql.DB
}

// Initial 初始化配置
func (ds *DataNodeRepo) Initial(st ifilestorage.DBSetting) (err error) {
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
func (ds *DataNodeRepo) GetSqlTx() (*sql.Tx, error) {
	return ds.db.Begin()
}

// GetSqlTx 获取数据库对象
func (ds *DataNodeRepo) GetDB() *sql.DB {
	return ds.db
}

// UpdateById 更新节点
func (ds *DataNodeRepo) UpdateById(conn *sql.Tx, node ifilestorage.DNode) (err error) {
	if len(node.Id) == 0 {
		return errors.New("id is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("update " + ds.table + " set sha256=? where id=?")
	if nil == err {
		_, err = stmt.Exec(node.Hash, node.Id)
	}
	return err
}

// DeleteByID 删除节点
func (ds *DataNodeRepo) DeleteByID(conn *sql.Tx, id string) (err error) {
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

// DeleteInIDs 删除节点
func (ds *DataNodeRepo) DeleteInIDs(conn *sql.Tx, ids []string) (err error) {
	if len(ids) == 0 {
		return errors.New("id is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("delete from " + ds.table + " where id in ('" + strings.Join(ids, "','") + "')")
	if nil == err {
		_, err = stmt.Exec()
	}
	return err
}

// InsertNode 新增节点
func (ds *DataNodeRepo) InsertNode(conn *sql.Tx, node ifilestorage.DNode) (err error) {
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("insert into " + ds.table + "(id, size, sha256, pieces) values(?, ?, ?, ?)")
	if nil == err {
		_, err = stmt.Exec(node.Id, node.Size, node.Hash, strings.Join(node.Pieces, ","))
	}
	return err
}

// ListPiecesInIDs 获取文件pieces信息
func (ds *DataNodeRepo) ListPiecesInIDs(conn *sql.DB, ids []string) (pieces []string, err error) {
	if len(ids) == 0 {
		return pieces, errors.New("id is empty")
	}
	sqlstr := "select pieces from " + ds.table + " where id in('" + strings.Join(ids, "','") + "')"
	var rows *sql.Rows
	if rows, err = conn.Query(sqlstr); nil == err {
		defer rows.Close()
		for rows.Next() {
			var piece string
			if err = rows.Scan(&piece); nil == err {
				pieces = append(pieces, piece)
			} else {
				pieces = make([]string, 0)
				break
			}
		}
	}
	return pieces, err
}

// GetNode 在数据库中查找节点
func (ds *DataNodeRepo) GetNode(conn *sql.DB, id string) (*ifilestorage.DNode, error) {
	if len(id) == 0 {
		return nil, errors.New("id is empty")
	}
	sql := "select id, size, sha256, pieces from " + ds.table + " where id=?"
	rows, err := conn.Query(sql, id)
	if nil == err {
		defer rows.Close()
		if rows.Next() {
			node := &ifilestorage.DNode{}
			var pieces string
			if err = rows.Scan(&node.Id, &node.Size, &node.Hash, &pieces); nil == err {
				node.Pieces = strings.Split(pieces, ",")
				return node, nil
			}
		}
	}
	return nil, err
}

// CreateTables 初始化结构
func (ds *DataNodeRepo) CreateTables(conn *sql.Tx) (err error) {
	tableSql := "create table if not exists `" + ds.table + "`  (`id` char(36)  not null, `size` bigint not null, `sha256` char(64) null, `pieces` text, primary key (`id`))"
	IndexSql := []string{
		"create index " + ds.table + "_index_id on " + ds.table + " (id)",
		"create index " + ds.table + "_index_sha256 on " + ds.table + " (sha256)",
	}

	_, err = conn.Exec(tableSql)
	if nil == err {
		for i := 0; i < len(IndexSql); i++ {
			if _, err = conn.Exec(IndexSql[i]); nil != err {
				break
			}
		}
	}
	return err
}
