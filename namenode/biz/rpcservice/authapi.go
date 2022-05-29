// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// RPC会话认证
package rpcservice

import (
	"namenode/ifilestorage"
	"pakku/ipakku"
)

// UserAccess access内容
type UserAccess struct {
	AccessKey string
	SecretKey string
}

// User 用户
type User struct {
	User     string
	Passwd   string
	PropJSON string
}

// AuthApi 认证api
type AuthApi struct {
	ev ipakku.AppSyncEvent       `@autowired:"AppEvent"`
	au ifilestorage.UserAuth4Rpc `@autowired:"User4RPC"`
}

func (n *AuthApi) DoLogin(req *User, res *UserAccess) error {
	if access, err := n.au.AskAccess(req.User, req.Passwd); nil != err {
		return err
	} else {
		res.AccessKey = access.AccessKey
		res.SecretKey = access.SecretKey
	}
	if nil == n.ev {
		return nil
	}
	return n.ev.PublishSyncEvent("authApi", "dologin.succeed", *req)
}
func (n *AuthApi) DoPing(req *string, res *bool) error {
	if err := n.au.RefreshAccessKey(*req); nil != err {
		return err
	}
	*res = true
	return nil
}
