// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// namenode rpc服务

package rpcservice

import (
	"datanode/ifilestorage"
	"encoding/gob"
	"pakku/utils/strutil"
)

func init() {
	gob.RegisterName("RequestBody.SrcAndDstBo", SrcAndDstBo{})
	gob.RegisterName("RequestBody.WriteTokenBo", WriteTokenBo{})
	gob.RegisterName("RequestBody.LimitQueryBo", LimitQueryBo{})
	gob.RegisterName("RequestBody.DirNameListDto", DirNameListDto{})
}

// SrcAndDstBo src 和 dst
type SrcAndDstBo struct {
	SRC string
	DST string
}

// WriteTokenBo 文件上传递交token
type WriteTokenBo struct {
	Token    string
	Override bool
}

// LimitQueryBo 分页查询
type LimitQueryBo struct {
	Query  string
	Limit  int
	Offset int
}

// DirNameListDto 文件夹目录列表
type DirNameListDto struct {
	Datas []string
	Total int
}

// RequestBody4RPC RPC请求体
type RequestBody4RPC struct {
	AccessKey   string
	Signature   string
	RequestBody interface{}
}

// doValidation 验证auth
func (bd *RequestBody4RPC) doValidation(ua ifilestorage.UserAuth4Rpc, accessKey, playload string, userTypes int64) error {
	if access, err := ua.GetUserAccess(accessKey); nil == err {
		if err := bd.checkPermission(access.UserType, userTypes); nil != err {
			return err
		}
		// 校验参数完整性
		if sg := strutil.GetSHA256(bd.AccessKey + "/" + access.SecretKey + "/" + playload); sg != bd.Signature {
			return ErrorAuthentication
		}
		return nil
	} else {
		return err
	}
}

// checkPermission 是否包含某个角色
func (bd *RequestBody4RPC) checkPermission(userType int, userTypes int64) error {
	if 1<<userType == (userTypes & (1 << userType)) {
		return nil
	}
	return ErrorPermissionLess
}
