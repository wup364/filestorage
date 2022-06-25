// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 用db存放数据

package user4rpc

import (
	"database/sql"
	"datanode/ifilestorage"
	"time"

	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

// UserStory 存储
type UserStory struct {
	db *sql.DB
}

// Initial 初始化配置
func (us *UserStory) Initial(st ifilestorage.DBSetting) (err error) {
	if nil == us.db {
		us.db, err = sql.Open(st.DriverName, st.DataSourceName)
		if nil == err {
			if st.DriverName == "sqlite3" {
				// database is locked
				// https://github.com/mattn/go-sqlite3/issues/209
				us.db.SetMaxOpenConns(1)
			} else {
				us.db.SetMaxIdleConns(250)
				us.db.SetConnMaxLifetime(time.Hour)
			}
		}
	}
	return err
}

// Install 初始化 users 表
func (us *UserStory) Install() (err error) {
	var tx *sql.Tx
	if tx, err = us.db.Begin(); err == nil {
		if _, err = tx.Exec(
			`create table if not exists users(
				userid varchar(64) primary key,
				userpwd varchar(255) default '',
				usertype varchar(255) default '',
				username varchar(255) default '',
				cttime date null
			);`); nil == err {
			//
			err = tx.Commit()
		} else {
			tx.Rollback()
		}
	}
	return err
}

// ListAllUsers 列出所有用户数据, 无分页
func (us *UserStory) ListAllUsers() ([]ifilestorage.UserInfo, error) {
	rows, err := us.db.Query("SELECT userid,username,usertype,cttime FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//
	res := make([]ifilestorage.UserInfo, 0)
	for rows.Next() {
		user := ifilestorage.UserInfo{}
		if err := rows.Scan(&user.UserID, &user.UserName, &user.UserType, &user.CtTime); err != nil {
			return nil, err
		}
		res = append(res, user)
	}
	return res, nil
}

// QueryUser 根据用户ID查询详细信息
func (us *UserStory) QueryUser(userID string) (*ifilestorage.UserInfo, error) {
	rows, err := us.db.Query("SELECT userid,username,usertype,cttime FROM users where userid='" + userID + "'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	//
	if rows.Next() {
		user := ifilestorage.UserInfo{}
		if err := rows.Scan(&user.UserID, &user.UserName, &user.UserType, &user.CtTime); err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, nil
}

// AddUser 添加用户
func (us *UserStory) AddUser(user *ifilestorage.CreateUserBo) (err error) {
	if len(user.UserID) == 0 {
		return ErrorUserIDIsNil
	}
	if len(user.UserName) == 0 {
		return ErrorUserNameIsNil
	}
	// 开启事务
	var ts *sql.Tx
	if ts, err = us.db.Begin(); err != nil {
		return err
	}
	//
	stmt, err := ts.Prepare("INSERT INTO users(userid,username,usertype,userpwd,cttime) values(?,?,?,?,?)")
	if err != nil {
		return err
	}
	//
	if _, err = stmt.Exec(user.UserID, user.UserName, user.UserType, strutil.GetMD5(user.UserPWD), time.Now()); err == nil {
		err = ts.Commit()
	} else {
		ts.Rollback()
	}
	return err
}

// UpdateUser 修改用户
func (us *UserStory) UpdateUser(user *ifilestorage.UserInfo) (err error) {
	if len(user.UserID) == 0 {
		return ErrorUserIDIsNil
	}
	// 开启事务
	var ts *sql.Tx
	if ts, err = us.db.Begin(); err != nil {
		return err
	}
	//
	var stmt *sql.Stmt
	if stmt, err = ts.Prepare("UPDATE users SET username=?, usertype=? WHERE userid=?"); err != nil {
		return err
	}
	//
	if _, err = stmt.Exec(user.UserName, user.UserType, user.UserID); err == nil {
		err = ts.Commit()
	} else {
		ts.Rollback()
	}
	return err
}

// DelUser 根据userId删除用户
func (us *UserStory) DelUser(userID string) (err error) {
	if len(userID) == 0 {
		return ErrorUserIDIsNil
	}
	// 开启事务
	var ts *sql.Tx
	if ts, err = us.db.Begin(); err != nil {
		return err
	}
	//
	var stmt *sql.Stmt
	if stmt, err = ts.Prepare("DELETE FROM users WHERE userid = ?"); err != nil {
		return err
	}
	//
	if _, err = stmt.Exec(userID); err == nil {
		err = ts.Commit()
	} else {
		ts.Rollback()
	}
	return err
}

// CheckPwd 校验密码是否一致
func (us *UserStory) CheckPwd(userID, pwd string) bool {
	if len(userID) > 0 {
		if rows, err := us.db.Query("SELECT userid FROM users where userid='" + userID + "' and userpwd='" + strutil.GetMD5(pwd) + "'"); nil == err {
			defer rows.Close()
			if rows.Next() {
				return true
			}
		} else {
			logs.Errorln(err)
		}
	}
	return false
}

// UpdateUser 修改用户密码
func (us *UserStory) UpdatePWD(user *ifilestorage.PwdUpdateBo) (err error) {
	if len(user.UserID) == 0 {
		return ErrorUserIDIsNil
	}
	// 开启事务
	var ts *sql.Tx
	if ts, err = us.db.Begin(); err != nil {
		return err
	}
	//
	var stmt *sql.Stmt
	if stmt, err = ts.Prepare("UPDATE users SET userpwd=? WHERE userid=?"); err != nil {
		return err
	}
	//
	if _, err = stmt.Exec(strutil.GetMD5(user.UserPWD), user.UserID); err == nil {
		err = ts.Commit()
	} else {
		ts.Rollback()
	}
	return err
}
