// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// namenode RPC服务

package rpcservice

import (
	"errors"
	"namenode/ifilestorage"
)

// NameNodeUser 命名空间节点管理
type NameNodeUser struct {
	ua ifilestorage.UserAuth4Rpc `@autowired:"User4RPC"`
	um ifilestorage.UserManage   `@autowired:"User4RPC"`
}

func (n *NameNodeUser) DoCreateUser(req *RequestBody4RPC, res *bool) (err error) {
	if err = req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_NameNode); nil != err {
		return err
	}
	user := req.RequestBody.(CreateUserBo)
	if err = n.um.AddUser(&ifilestorage.CreateUserBo{
		UserID:   user.UserID,
		UserPWD:  user.UserPWD,
		UserType: user.UserType,
		UserName: user.UserName,
	}); nil == err {
		*res = true
	}
	return err
}

func (n *NameNodeUser) DoListAllUsers(req *RequestBody4RPC, res *UserListDto) (err error) {
	if err = req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_NameNode); nil != err {
		return err
	}
	if u, err := n.um.ListAllUsers(); err != nil {
		return err
	} else if res.Total = len(u); res.Total > 0 {
		res.Datas = u
	}
	return err
}

func (n *NameNodeUser) DoUpdatePWD(req *RequestBody4RPC, res *bool) (err error) {
	if err = req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_NameNode); nil != err {
		return err
	}
	user := req.RequestBody.(PwdUpdateBo)
	if err = n.um.UpdatePWD(user.User, user.Pwd); nil == err {
		*res = true
	}
	return err
}

func (n *NameNodeUser) DoDeleteUser(req *RequestBody4RPC, res *bool) (err error) {
	if err = req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_NameNode); nil != err {
		return err
	}
	userID := req.RequestBody.(string)
	if ack, err := n.ua.GetUserAccess(req.AccessKey); nil != err {
		return err
	} else if ack.UserID == userID {
		return errors.New("cannot delete current account")
	}
	if err = n.um.DelUser(userID); nil == err {
		*res = true
	}
	return err
}
