// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package ifilestorage

import (
	"encoding/json"
	"fmt"
	"time"
)

const (
	UserRule_NameNode = 0
	UserRule_DataNode = 1
	UserRule_OpenApi  = 2
)

// User4RPC 用户管理接口
type User4RPC interface {
	UserManage
	UserAuth4Rpc
}

// UserManage access接口
type UserManage interface {
	Clear() error
	AddUser(user *CreateUserBo) error             // 添加用户
	DelUser(userID string) error                  // 根据userID删除用户
	CheckPwd(userID, pwd string) bool             // 校验密码是否一致
	UpdatePWD(userID, pwd string) error           // 修改用户密码
	ListAllUsers() ([]UserInfo, error)            // 列出所有用户数据, 无分页
	QueryUser(userID string) (*UserInfo, error)   // 根据用户ID查询详细信息
	UpdateUserName(userID, userName string) error // 修改用户名字
}

// UserAuth4Rpc access接口
type UserAuth4Rpc interface {
	// AskAccess 获取access
	AskAccess(userID, pwd string) (*UserAccess, error)
	// GetSecretKey 获取 userAccess
	GetUserAccess(accessKey string) (*UserAccess, error)
	// RefreshAccessKey 刷新 access
	RefreshAccessKey(accessKey string) error
}

// CreateUserBo 创建用户
type CreateUserBo struct {
	UserType int
	UserID   string
	UserName string
	UserPWD  string
}

// PwdUpdateBo 更改用户密码
type PwdUpdateBo struct {
	UserID  string
	UserPWD string
}

// UserInfo 用户表存储的结构
type UserInfo struct {
	UserType int
	UserID   string
	UserName string
	UserPWD  string
	CtTime   time.Time
}

// UserAccess access内容
type UserAccess struct {
	UserID    string
	UserType  int
	UserName  string
	AccessKey string
	SecretKey string
}

// Clone Clone
func (ua *UserAccess) Clone(val interface{}) error {
	if uat, ok := val.(*UserAccess); ok {
		uat.UserID = ua.UserID
		uat.UserType = ua.UserType
		uat.UserName = ua.UserName
		uat.AccessKey = ua.AccessKey
		uat.SecretKey = ua.SecretKey
		return nil
	}
	return fmt.Errorf("can't support clone %T ", val)
}

// ToJSON ToJSON
func (ua *UserAccess) ToJSON() string {
	if bt, err := json.Marshal(ua); nil == err {
		return string(bt)
	}
	return ""
}
