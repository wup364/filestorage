// Copyright (C) 2022 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 服务器管理接口

package opensdk

// NewSvrMngApi 服务器管理接口, 如: 用户管理
func NewSvrMngApi(addr string, user User) IServerMG {
	return &ServerMgSdkImpl{NewConn4RPC(addr, user)}
}

// ServerMgSdkImpl 开发api接口
type ServerMgSdkImpl struct {
	conn IConn4RPC
}

// ConnTest ConnTest
func (s *ServerMgSdkImpl) ConnTest(user User) error {
	if client, err := s.conn.GetClient(); nil != err {
		return err
	} else {
		defer client.Close()
		var reply bool
		if err = client.Call("AuthApi.DoCheckPwd", user, &reply); nil != err {
			return err
		}
	}
	return nil
}

func (s *ServerMgSdkImpl) DoCreateUser(user CreateUserBo) (bool, error) {
	var res bool
	if err := s.conn.DoRPC("NameNodeUser.DoCreateUser", user, &res); nil != err {
		return false, err
	}
	return res, nil
}

func (s *ServerMgSdkImpl) DoListAllUsers() (*UserListDto, error) {
	var res UserListDto
	if err := s.conn.DoRPC("NameNodeUser.DoListAllUsers", nil, &res); nil != err {
		return nil, err
	}
	return &res, nil
}

func (s *ServerMgSdkImpl) DoUpdatePWD(user, pwd string) (bool, error) {
	var res bool
	wt := PwdUpdateBo{User: user, Pwd: pwd}
	if err := s.conn.DoRPC("NameNodeUser.DoUpdatePWD", wt, &res); nil != err {
		return false, err
	}
	return res, nil
}

func (s *ServerMgSdkImpl) DoDeleteUser(userID string) (bool, error) {
	var res bool
	if err := s.conn.DoRPC("NameNodeUser.DoDeleteUser", userID, &res); nil != err {
		return false, err
	}
	return res, nil
}
