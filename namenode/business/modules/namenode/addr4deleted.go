// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 已删除的文件地址

package namenode

import (
	"database/sql"
	"errors"
	"namenode/ifilestorage"
	"strings"
	"time"

	"github.com/wup364/pakku/utils/strutil"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// Addr4Deleted 已删除的文件地址
type Addr4Deleted struct {
	table string
	db    *sql.DB
}

// Initial 初始化配置
func (d *Addr4Deleted) Initial(table string, st ifilestorage.DBSetting) (err error) {
	if nil == d.db {
		d.table = table
		d.db, err = sql.Open(st.DriverName, st.DataSourceName)
		if nil == err {
			if st.DriverName == "sqlite3" {
				// database is locked
				// https://github.com/mattn/go-sqlite3/issues/209
				d.db.SetMaxOpenConns(1)
			} else {
				d.db.SetMaxIdleConns(250)
				d.db.SetConnMaxLifetime(time.Hour)
			}
		}
	}
	return err
}

// GetSqlTx 获取事物
func (d *Addr4Deleted) GetSqlTx() (*sql.Tx, error) {
	return d.db.Begin()
}

// GetSqlTx 获取数据库对象
func (d *Addr4Deleted) GetDB() *sql.DB {
	return d.db
}

// DeleteByID 删除节点
func (d *Addr4Deleted) DeleteByID(conn *sql.Tx, id string) (err error) {
	if len(id) == 0 {
		return errors.New("id is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("delete from " + d.table + " where id=?")
	if nil == err {
		_, err = stmt.Exec(id)
	}
	return err
}

// DeleteInID 删除节点
func (d *Addr4Deleted) DeleteInID(conn *sql.Tx, ids []string) (err error) {
	if len(ids) == 0 {
		return errors.New("id is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("delete from " + d.table + " where id in ('" + strings.Join(ids, "','") + "')")
	if nil == err {
		var res sql.Result
		if res, err = stmt.Exec(); nil == err {
			if count, _ := res.RowsAffected(); count == 0 {
				return errors.New("delete failed")
			}
		}
	}
	return err
}

// InsertItem 新增节点
func (d *Addr4Deleted) InsertItem(conn *sql.Tx, addr string) (err error) {
	var stmt *sql.Stmt
	if stmt, err = conn.Prepare("insert into " + d.table + "(id, ctime, nodeno, fid) values(?, ?, ?, ?)"); nil == err {
		s := strings.Split(addr, "@")
		_, err = stmt.Exec(strutil.GetUUID(), time.Now().UnixMilli(), s[0], s[1])
	}
	return err
}

// FindByNodeNo 根据nodeno查找
func (d *Addr4Deleted) FindByNodeNo(conn *sql.DB, nodeno string, ctime int64, limit int) ([]ifilestorage.DeletedAddr, error) {
	if len(nodeno) == 0 {
		return nil, errors.New("nodeno is empty")
	}
	sql := "select id, ctime, fid from " + d.table + " where nodeno=? and ctime<? order by ctime limit ?"
	rows, err := conn.Query(sql, nodeno, ctime, limit)
	if nil != err {
		return nil, err
	}
	defer rows.Close()
	var nodes []ifilestorage.DeletedAddr
	for rows.Next() {
		node := ifilestorage.DeletedAddr{NodeNo: nodeno}
		if err := rows.Scan(&node.Id, &node.Ctime, &node.Fid); nil != err {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

// ListNodeNos 列出所有nodeno
func (d *Addr4Deleted) ListNodeNos(conn *sql.DB) ([]string, error) {
	sql := "select nodeno from " + d.table + " group by nodeno"
	rows, err := conn.Query(sql)
	if nil != err {
		return nil, err
	}
	defer rows.Close()
	var nodes []string
	for rows.Next() {
		var nodeno string
		if err := rows.Scan(&nodeno); nil != err {
			return nil, err
		}
		nodes = append(nodes, nodeno)
	}
	return nodes, nil
}

// CreateTables 初始化结构
func (d *Addr4Deleted) CreateTables(conn *sql.Tx) (err error) {
	dropSql := "drop table  `" + d.table + "`"
	tableSql := "create table if not exists `" + d.table + "`  (`id` char(36)  not null,`ctime` bigint not null,`nodeno` char(64) not null,`fid` char(64) not null)"
	IndexSql := []string{
		"create index " + d.table + "_index_fid on " + d.table + " (fid)",
		"create index " + d.table + "_index_ctime on " + d.table + " (ctime)",
		"create index " + d.table + "_index_nodeno on " + d.table + " (nodeno)",
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
