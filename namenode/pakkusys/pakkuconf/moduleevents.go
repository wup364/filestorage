// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 程序启动前的相关配置
package pakkuconf

import (
	"namenode/ifilestorage"
	"namenode/pakkusys"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/logs"
)

// crateDefaultUsers 创建默认用户
func crateDefaultUsers() pakkusys.ModuleEvent {
	return pakkusys.ModuleEvent{
		Module: "User4RPC",
		Event:  ipakku.ModuleEventOnSetup,
		Handler: func(module interface{}, _ ipakku.Loader) {
			umg := module.(ifilestorage.UserManage)
			if err := umg.Clear(); nil == err {
				err = umg.AddUser(&ifilestorage.CreateUserBo{
					UserType: ifilestorage.UserRule_NameNode,
					UserID:   "NAMENODE",
					UserName: "NameNode operator",
					UserPWD:  "",
				})
				if nil == err {
					err = umg.AddUser(&ifilestorage.CreateUserBo{
						UserType: ifilestorage.UserRule_DataNode,
						UserID:   "DATANODE",
						UserName: "DataNode operator",
						UserPWD:  "",
					})
				}
				if nil == err {
					err = umg.AddUser(&ifilestorage.CreateUserBo{
						UserType: ifilestorage.UserRule_OpenApi,
						UserID:   "OPENAPI",
						UserName: "OpenApi operator",
						UserPWD:  "",
					})
				}
				if nil != err {
					logs.Panicln(err)
				}
			} else {
				logs.Panicln(err)
			}

		},
	}
}
