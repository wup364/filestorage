// Copyright (C) 2020 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package user4rpc

import (
	"errors"
)

const (
	Clib_Access = "User4RPC.AccessToken"
)

// ErrorConnIsNil 空连接
var ErrorConnIsNil = errors.New("the data source is empty")

// ErrorUserIDIsNil ErrorUserIDIsNil
var ErrorUserIDIsNil = errors.New("the userID is empty")

// ErrorUserNotExist ErrorUserNotExist
var ErrorUserNotExist = errors.New("user does not exist")

// ErrorUserNameIsNil ErrorUserNameIsNil
var ErrorUserNameIsNil = errors.New("the userName is empty")

// ErrorAuthentication 认证失败
var ErrorAuthentication = errors.New("authentication failed")

// ErrCreateFailed101 创建失败, 失败原因: 账户已存在
var ErrCreateFailed101 = errors.New("account already exists")
