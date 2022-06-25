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
	"datanode/ifilestorage"
	"datanode/pakkusys"
	"net/http"
	"strings"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

// crateDefaultUsers 创建默认用户
func crateDefaultUsers() pakkusys.ModuleEvent {
	return pakkusys.ModuleEvent{
		Module: "User4RPC",
		Event:  ipakku.ModuleEventOnLoaded,
		Handler: func(module interface{}, loader ipakku.Loader) {
			var conf ipakku.AppConfig
			umg := module.(ifilestorage.UserManage)
			if err := loader.GetModuleByName("AppConfig", &conf); nil != err {
				logs.Panicln(err)
			}
			if err := umg.Clear(); nil != err {
				logs.Panicln(err)
			}
			nodeNo := conf.GetConfig("nodeno").ToString("DN101")
			if err := umg.AddUser(&ifilestorage.CreateUserBo{
				UserType: ifilestorage.UserRule_DataNode,
				UserID:   nodeNo,
				UserName: nodeNo,
				UserPWD:  strutil.GetSHA256(nodeNo + "@" + loader.GetInstanceID()),
			}); nil == err {
				loader.SetParam("nodeno", nodeNo)
			} else {
				logs.Panicln(err)
			}
		},
	}
}

// httpGlobalFilter 全局过滤器
func httpGlobalFilter() pakkusys.ModuleEvent {
	return pakkusys.ModuleEvent{
		Module: "AppService",
		Event:  ipakku.ModuleEventOnLoaded,
		Handler: func(module interface{}, loader ipakku.Loader) {
			service := module.(ipakku.AppService)
			// CORS过滤器
			service.Filter("/:*", func(w http.ResponseWriter, r *http.Request) bool {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				if acrhs := r.Header["Access-Control-Request-Headers"]; len(acrhs) > 0 {
					w.Header().Set("Access-Control-Allow-Headers", strings.Join(acrhs, ","))
				}
				if acrms := r.Header["Access-Control-Request-Method"]; len(acrms) > 0 {
					w.Header().Set("Access-Control-Allow-Methods", strings.Join(acrms, ","))
				}
				if strings.ToUpper(r.Method) == "OPTIONS" {
					w.WriteHeader(http.StatusOK)
					return false
				}
				return true
			})
		},
	}
}
