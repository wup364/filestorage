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
	"namenode/biz/modules/filenames"
	"namenode/biz/modules/namenode"
	"namenode/biz/modules/user4rpc"
	"namenode/biz/rpcservice"
	"namenode/pakkusys"

	"github.com/wup364/pakku/ipakku"
)

// RegisterModules 注册需要加载的模块
func RegisterModules() []ipakku.Module {
	return []ipakku.Module{
		new(user4rpc.User4RPC),
		new(filenames.FileNames),
		new(namenode.NameNode),
	}
}

// RegisterController 注册需要加载的controller
func RegisterController() []ipakku.Controller {
	return []ipakku.Controller{}
}

// RegisterRPCService 注册需要加载的gorpc
func RegisterRPCService() []interface{} {
	return []interface{}{
		new(rpcservice.AuthApi),
		new(rpcservice.NameNode),
		new(rpcservice.NameNodeUser),
	}
}

// RegisterModuleEvent 注册模块加载事件
func RegisterModuleEvent() []pakkusys.ModuleEvent {
	return []pakkusys.ModuleEvent{
		crateDefaultUsers(),
	}
}

// RegisterOverride 注册复写的模块
func RegisterOverride() []pakkusys.OverrideModule {
	return []pakkusys.OverrideModule{}
}
