// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 用户管理模块

package user4rpc

import (
	"namenode/ifilestorage"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

// User4RPC 用户管理模块
type User4RPC struct {
	us *UserRepo
	c  ipakku.AppConfig `@autowired:"AppConfig"`
	ch ipakku.AppCache  `@autowired:"AppCache"`
}

// AsModule 作为一个模块
func (umg *User4RPC) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Name:        "User4RPC",
		Version:     1.0,
		Description: "RPC用户鉴权",
		OnReady: func(mctx ipakku.Loader) {
			umg.us = &UserRepo{}
			deftDataSource := "./.datas/" + mctx.GetParam(ipakku.PARAMKEY_APPNAME).ToString("app") + "#user?cache=shared"
			confDataSource := umg.c.GetConfig("store.user4rpc.datasource").ToString(deftDataSource)
			if confDataSource == deftDataSource {
				umg.mkSqliteDIR() // 创建sqlite文件存放目录
			}
			if err := umg.us.Initial(ifilestorage.DBSetting{
				DriverName:     umg.c.GetConfig("store.user4rpc.driver").ToString("sqlite3"),
				DataSourceName: confDataSource,
			}); nil != err {
				logs.Panicln(err)
			}
			// 注册 accesstoken 缓存库, x分钟过期
			if err := umg.ch.RegLib(Clib_Access, 60*60*24); nil != err {
				logs.Panicln(err)
			}
		},
		OnSetup: func() {
			if err := umg.us.Install(); nil != err {
				logs.Panicln(err)
			}
		},
		OnInit: func() {
		},
	}
}

// ListAllUsers 列出所有用户数据, 无分页
func (umg *User4RPC) ListAllUsers() ([]ifilestorage.UserInfo, error) {
	return umg.us.ListAllUsers()
}

// QueryUser 根据用户ID查询详细信息
func (umg *User4RPC) QueryUser(userID string) (*ifilestorage.UserInfo, error) {
	return umg.us.QueryUser(userID)
}

// Clear 清空用户
func (umg *User4RPC) Clear() error {
	if users, err := umg.us.ListAllUsers(); nil == err {
		if len(users) > 0 {
			for _, val := range users {
				if err = umg.us.DelUser(val.UserID); nil != err {
					return err
				}
			}
		}
	} else {
		return err
	}
	return nil
}

// AddUser 添加用户
func (umg *User4RPC) AddUser(user *ifilestorage.CreateUserBo) error {
	if u, _ := umg.QueryUser(user.UserID); nil != u {
		return ErrCreateFailed101
	}
	return umg.us.AddUser(user)
}

// UpdatePWD 修改用户密码
func (umg *User4RPC) UpdatePWD(userID, pwd string) error {
	if userOld, err := umg.QueryUser(userID); nil != err {
		return err
	} else if nil == userOld {
		return ErrorUserNotExist
	}
	return umg.us.UpdatePWD(&ifilestorage.PwdUpdateBo{
		UserID:  userID,
		UserPWD: pwd,
	})
}

// UpdateUserName 修改用户昵称
func (umg *User4RPC) UpdateUserName(userID, userName string) error {
	if len(userName) == 0 {
		return ErrorUserNameIsNil
	}
	userOld, err := umg.QueryUser(userID)
	if nil != err {
		return err
	}
	if nil == userOld {
		return ErrorUserNotExist
	}
	userOld.UserName = userName
	return umg.us.UpdateUser(userOld)
}

// DelUser 根据userID删除用户
func (umg *User4RPC) DelUser(userID string) error {
	return umg.us.DelUser(userID)
}

// CheckPwd 校验密码是否一致
func (umg *User4RPC) CheckPwd(userID, pwd string) bool {
	return umg.us.CheckPwd(userID, pwd)
}

// AskAccess 获取access
func (umg *User4RPC) AskAccess(userID, pwd string) (*ifilestorage.UserAccess, error) {
	if !umg.CheckPwd(userID, pwd) {
		return nil, ErrorAuthentication
	}
	if user, err := umg.QueryUser(userID); nil != err {
		return nil, err
	} else {
		access := &ifilestorage.UserAccess{
			UserID:    user.UserID,
			UserType:  user.UserType,
			UserName:  user.UserName,
			AccessKey: strutil.GetUUID(),
			SecretKey: strutil.GetUUID(),
		}
		if err := umg.ch.Set(Clib_Access, access.AccessKey, access); nil != err {
			return nil, err
		}
		return access, nil
	}
}

// GetUserAccess 获取 useraccess
func (umg *User4RPC) GetUserAccess(accessKey string) (*ifilestorage.UserAccess, error) {
	var val ifilestorage.UserAccess
	if err := umg.ch.Get(Clib_Access, accessKey, &val); nil == err {
		return &val, nil
	} else if err != ipakku.ErrNoCacheHit {
		return nil, err
	}
	return nil, ErrorAuthentication
}

// RefreshAccessKey 刷新 access
func (umg *User4RPC) RefreshAccessKey(accessKey string) error {
	var val ifilestorage.UserAccess
	if err := umg.ch.Get(Clib_Access, accessKey, &val); nil == err {
		if err := umg.ch.Set(Clib_Access, val.AccessKey, &val); nil != err {
			return err
		}
		return nil
	}
	return ErrorAuthentication
}

// mkSqliteDIR 创建sqlite文件存放目录
func (umg *User4RPC) mkSqliteDIR() {
	if !fileutil.IsExist("./.datas") {
		if err := fileutil.MkdirAll("./.datas"); nil != err {
			logs.Panicln(err)
		}
	}
}
