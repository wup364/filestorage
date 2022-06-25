// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 树-数据库中
// 默认sqlite3, sqlite3提供的性能有限(插入120条/S), 可以使用mysql获取更快的性能

package names

import (
	"database/sql"
	"errors"
	"namenode/ifilestorage"
	"time"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// NameStory 树-数据库中
type NameStory struct {
	table string
	db    *sql.DB
}

// Initial 初始化配置
func (nm *NameStory) Initial(table string, st ifilestorage.DBSetting) (err error) {
	if nil == nm.db {
		nm.table = table
		nm.db, err = sql.Open(st.DriverName, st.DataSourceName)
		if nil == err {
			if st.DriverName == "sqlite3" {
				// database is locked
				// https://github.com/mattn/go-sqlite3/issues/209
				nm.db.SetMaxOpenConns(1)
			} else {
				nm.db.SetMaxIdleConns(250)
				nm.db.SetConnMaxLifetime(time.Hour)
			}
		}
	}
	return err
}

// GetSqlTx 获取事物
func (nm *NameStory) GetSqlTx() (*sql.Tx, error) {
	return nm.db.Begin()
}

// GetSqlTx 获取数据库对象
func (nm *NameStory) GetDB() *sql.DB {
	return nm.db
}

// UpdateById 更新节点
func (nm *NameStory) UpdateById(conn *sql.Tx, id string, node ifilestorage.TNode4Update) (err error) {
	if len(id) == 0 {
		return errors.New("id is empty")
	}
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("update " + nm.table + " set name=?, size=?, mtime=?, addr=?, props=? where id=?")
	if nil == err {
		_, err = stmt.Exec(node.Name, node.Size, node.Mtime, node.Addr, node.Props, id)
	}
	return err
}

// MoveNode 更新节点-重新设置pid
func (nm *NameStory) MoveNode(conn *sql.Tx, srcid, name, dstid string) (err error) {
	if len(srcid) == 0 {
		return errors.New("src id is empty")
	}
	if len(dstid) == 0 {
		return errors.New("dst id is empty")
	}
	var stmt *sql.Stmt
	if stmt, err = conn.Prepare("update " + nm.table + " set pid=?, name=? where id=?"); nil == err {
		_, err = stmt.Exec(dstid, name, srcid)
	}
	return err
}

// DeleteByID 删除节点
func (nm *NameStory) DeleteByID(conn *sql.Tx, id string) (err error) {
	if len(id) == 0 {
		return errors.New("id is empty")
	}
	var stmt *sql.Stmt
	if stmt, err = conn.Prepare("delete from " + nm.table + " where id=?"); nil == err {
		_, err = stmt.Exec(id)
	}
	return err
}

// CountReferencedAddr 统计引用的文件地址
func (nm *NameStory) CountReferencedAddr(conn *sql.DB, addr string) (count int64, err error) {
	if len(addr) == 0 {
		return count, errors.New("addr is empty")
	}

	if rows, err := conn.Query("select count(addr) from "+nm.table+" where addr=?", addr); nil != err {
		return count, err
	} else {
		defer rows.Close()
		if rows.Next() {
			if err := rows.Scan(&count); nil != err {
				return count, err
			}
		}
	}
	return count, err
}

// InsertNode 新增节点
func (nm *NameStory) InsertNode(conn *sql.Tx, node ifilestorage.TNode4New) (err error) {
	var stmt *sql.Stmt
	stmt, err = conn.Prepare("insert into " + nm.table + "(id, pid, flag, name, size, ctime, mtime, addr, props) values(?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if nil == err {
		_, err = stmt.Exec(node.Id, node.Pid, node.Flag, node.Name, node.Size, node.Ctime, node.Mtime, node.Addr, node.Props)
	}

	return err
}

// GetRoot 在数据库中查找根节点, pid=root节点的父节点, 可以为空
func (nm *NameStory) GetRoot(conn *sql.DB, pid string) (*ifilestorage.TNode, error) {
	sql := "select id, pid, flag, name, size, ctime, mtime, props from " + nm.table + " where pid=?"
	rows, err := conn.Query(sql, pid)
	if nil != err {
		return nil, err
	}
	defer rows.Close()
	var node *ifilestorage.TNode
	if rows.Next() {
		node = &ifilestorage.TNode{}
		if err := rows.Scan(&node.Id, &node.Pid, &node.Flag, &node.Name, &node.Size, &node.Ctime, &node.Mtime, &node.Props); nil != err {
			return nil, err
		}
	}
	return node, nil
}

// GetNode 在数据库中查找节点
func (nm *NameStory) GetNode(conn *sql.DB, id string) (*ifilestorage.TNode, error) {
	if len(id) == 0 {
		return nil, errors.New("id is empty")
	}
	sql := "select id, pid, flag, name, size, ctime, mtime, addr, props from " + nm.table + " where id=?"
	rows, err := conn.Query(sql, id)
	if nil == err {
		defer rows.Close()
		if rows.Next() {
			node := &ifilestorage.TNode{}
			if err := rows.Scan(&node.Id, &node.Pid, &node.Flag, &node.Name, &node.Size, &node.Ctime, &node.Mtime, &node.Addr, &node.Props); nil != err {
				return nil, err
			}
			return node, nil
		}
	}
	return nil, err
}

// GetNode 在数据库中查找节点
func (nm *NameStory) GetNodeInTx(conn *sql.Tx, id string) (*ifilestorage.TNode, error) {
	if len(id) == 0 {
		return nil, errors.New("id is empty")
	}
	sql := "select id, pid, flag, name, size, ctime, mtime, addr, props from " + nm.table + " where id=?"
	rows, err := conn.Query(sql, id)
	if nil == err {
		defer rows.Close()
		node := &ifilestorage.TNode{}
		for rows.Next() {
			if err := rows.Scan(&node.Id, &node.Pid, &node.Flag, &node.Name, &node.Size, &node.Ctime, &node.Mtime, &node.Addr, &node.Props); nil != err {
				return nil, err
			}
		}
		return node, nil
	}
	return nil, err
}

// GetNodeChilds 在数据库中查找子节点 - offset用到在加吧
func (nm *NameStory) GetNodeChilds(conn *sql.DB, pid string, limit int) ([]*ifilestorage.TNode, error) {
	if len(pid) == 0 {
		return nil, errors.New("pid is empty")
	}
	sql := "select id, pid, flag, name, size, ctime, mtime, addr, props from " + nm.table + " where pid=? limit ?"
	rows, err := conn.Query(sql, pid, limit)
	if nil != err {
		return nil, err
	}
	defer rows.Close()
	var nodes []*ifilestorage.TNode
	for rows.Next() {
		node := &ifilestorage.TNode{}
		if err := rows.Scan(&node.Id, &node.Pid, &node.Flag, &node.Name, &node.Size, &node.Ctime, &node.Mtime, &node.Addr, &node.Props); nil != err {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

// CountInTime 统计一段时间内的数量
func (nm *NameStory) CountInTime(conn *sql.DB, timeStart, timeEnd int64) (int64, error) {
	sql := "select count(id) from " + nm.table + " where ctime >? and ctime <?"
	rows, err := conn.Query(sql, timeStart, timeEnd)
	if nil != err {
		return -1, err
	}
	defer rows.Close()
	if rows.Next() {
		var count int64
		if err := rows.Scan(&count); nil != err {
			return -1, err
		}
		return count, nil
	}
	return 0, nil
}

// GetNodeAllChilds 在数据库中查找子节点
func (nm *NameStory) GetNodeAllChilds(conn *sql.DB, pid string) ([]*ifilestorage.TNode, error) {
	if len(pid) == 0 {
		return nil, errors.New("pid is empty")
	}
	sql := "with recursive t1 as "
	sql += "(select t0.id, t0.pid, t0.flag, t0.name, t0.size, t0.ctime, t0.mtime, t0.addr, t0.props from " + nm.table + " t0 where t0.pid=? "
	sql += "union all select t2.id, t2.pid, t2.flag, t2.name, t2.size, t2.ctime, t2.mtime, t2.addr, t2.props from " + nm.table + " t2, t1 where t2.pid = t1.id) "
	sql += "select id, pid, flag, name, size, ctime, mtime, addr, props from t1"
	rows, err := conn.Query(sql, pid)
	if nil != err {
		return nil, err
	}
	defer rows.Close()
	var nodes []*ifilestorage.TNode
	for rows.Next() {
		node := &ifilestorage.TNode{}
		if err := rows.Scan(&node.Id, &node.Pid, &node.Flag, &node.Name, &node.Size, &node.Ctime, &node.Mtime, &node.Addr, &node.Props); nil != err {
			return nil, err
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

// WalkNodeAllChilds 一次性查出所有查出某个节点所有子节点
func (nm *NameStory) WalkNodeAllChilds(conn *sql.DB, id string, walk func(node *ifilestorage.TNode) error) error {
	if len(id) == 0 {
		return errors.New("id is empty")
	}
	sql := "with recursive t1 as "
	sql += "(select t0.id, t0.pid, t0.flag, t0.name, t0.size, t0.ctime, t0.mtime, t0.addr, t0.props from " + nm.table + " t0 where t0.pid=? "
	sql += "union all select t2.id, t2.pid, t2.flag, t2.name, t2.size, t2.ctime, t2.mtime, t2.addr, t2.props from " + nm.table + " t2, t1 where t2.pid = t1.id) "
	sql += "select id, pid, flag, name, size, ctime, mtime, addr, props from t1"
	rows, err := conn.Query(sql, id)
	if nil != err {
		return err
	}
	defer rows.Close()
	for rows.Next() {
		node := &ifilestorage.TNode{}
		if err := rows.Scan(&node.Id, &node.Pid, &node.Flag, &node.Name, &node.Size, &node.Ctime, &node.Mtime, &node.Addr, &node.Props); nil != err {
			return err
		}
		if err = walk(node); nil != err {
			return err
		}
	}
	return nil
}

// CountNodeChild 统计一级子节点数量
func (nm *NameStory) CountNodeChild(conn *sql.DB, id string) (int64, error) {
	if len(id) == 0 {
		return -1, errors.New("id is empty")
	}
	sql := "select count(id) from " + nm.table + " where pid=?"

	if rows, err := conn.Query(sql, id); nil != err {
		return -1, err
	} else {
		defer rows.Close()
		var count int64
		if rows.Next() {
			rows.Scan(&count)
		}
		return count, nil
	}
}

// CountNodeChild 统计一级子节点数量
func (nm *NameStory) CountNodeChildTx(conn *sql.Tx, id string) (int64, error) {
	if len(id) == 0 {
		return -1, errors.New("id is empty")
	}
	sql := "select count(id) from " + nm.table + " where pid=?"

	if rows, err := conn.Query(sql, id); nil != err {
		return -1, err
	} else {
		defer rows.Close()
		var count int64
		if rows.Next() {
			rows.Scan(&count)
		}
		return count, nil
	}
}

// LoopNodeChilds 递归查找数据库中查找子节点
func (nm *NameStory) LoopNodeChilds(conn *sql.DB, id string, walk func(node *ifilestorage.TNode) error, mode TreeLoopMode) error {
	if len(id) == 0 {
		return errors.New("id is empty")
	}
	sql := "select id, pid, flag, name, size, ctime, mtime, addr, props from " + nm.table + " where pid=?"
	rows, err := conn.Query(sql, id)
	if nil != err {
		return err
	}
	defer rows.Close()
	dirs := make([]string, 0)
	dirObjs := make([]*ifilestorage.TNode, 0)
	for rows.Next() {
		node := &ifilestorage.TNode{}
		if err = rows.Scan(&node.Id, &node.Pid, &node.Flag, &node.Name, &node.Size, &node.Ctime, &node.Mtime, &node.Addr, &node.Props); nil != err {
			return err
		}
		// 文件夹优先模式
		if mode == TREELOOPMODE_DIRFIRST {
			if err = walk(node); nil != err {
				return err
			}
			if node.Flag == FLAG_DIR {
				dirs = append(dirs, node.Id)
			}

			// 文件优先模式
		} else if mode == TREELOOPMODE_FILEFIRST {
			if node.Flag == FLAG_FILE {
				if err = walk(node); nil != err {
					return err
				}
			} else {
				dirObjs = append(dirObjs, node)
			}
		}
	}

	// 文件夹优先模式
	if mode == TREELOOPMODE_DIRFIRST {
		if len(dirs) > 0 {
			for i := 0; i < len(dirs); i++ {
				if err = nm.LoopNodeChilds(conn, dirs[i], walk, mode); nil != err {
					return err
				}
			}
		}

		// 文件优先模式
	} else if mode == TREELOOPMODE_FILEFIRST {
		if len(dirObjs) > 0 {
			for i := 0; i < len(dirObjs); i++ {
				if err = nm.LoopNodeChilds(conn, dirObjs[i].Id, walk, mode); nil != err {
					return err
				}
				if err = walk(dirObjs[i]); nil != err {
					return err
				}
			}
		}
	}

	return nil
}

// CreateTables 初始化结构
func (nm *NameStory) CreateTables(conn *sql.Tx) (err error) {
	dropSql := "drop table  `" + nm.table + "`"
	tableSql := "create table if not exists `" + nm.table + "`  (`id` char(36)  not null,`pid` char(36)  not null,`flag` int not null,`name` varchar(1024)  not null,`size` bigint not null,`ctime` bigint not null,`mtime` bigint not null,`addr` char(64) default '',`props` text default '', primary key (`id`))"
	IndexSql := []string{
		"create index " + nm.table + "_index_id on " + nm.table + " (id)",
		"create index " + nm.table + "_index_pid on " + nm.table + " (pid)",
		"create index " + nm.table + "_index_flag on " + nm.table + " (flag)",
		"create index " + nm.table + "_index_addr on " + nm.table + " (addr)",
		"create index " + nm.table + "_index_ctime on " + nm.table + " (ctime)",
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
