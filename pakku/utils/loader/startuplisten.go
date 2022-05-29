// Copyright (C) 2019 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 模块加载器 - 接口
// 依赖包: utils.utypes.Object utils.strutil.strutil

package loader

import (
	"pakku/utils/logs"
	"strconv"
)

// startupListen 默认监听处理
type startupListen struct{}

// bind 绑定事件监听
func (evt *startupListen) bind(l Loader) {
	l.OnModuleEvent("*", ModuleEventOnReady, evt.doReady)
	l.OnModuleEvent("*", ModuleEventOnSetup, evt.doSetup)
	l.OnModuleEvent("*", ModuleEventOnUpdate, evt.doUpdate)
	l.OnModuleEvent("*", ModuleEventOnInit, evt.doInit)
}

// doReady 模块准备
func (evt *startupListen) doReady(m interface{}, l Loader) {
	mt := m.(Module)
	if err := doAutoWired(mt, l); nil != err {
		logs.Panicln(err)
	}
	if nil != mt.AsModule().OnReady {
		logs.Infof("> On ready load \r\n")
		mt.AsModule().OnReady(l)
	}
}

// doSetup 模块安装
func (evt *startupListen) doSetup(m interface{}, l Loader) {
	mt := m.(Module)
	if nil != mt.AsModule().OnSetup {
		logs.Infof("> On setup module \r\n")
		mt.AsModule().OnSetup()
	}
}

// doUpdate 模块升级
func (evt *startupListen) doUpdate(m interface{}, l Loader) {
	mt := m.(Module)
	if nil != mt.AsModule().OnUpdate {
		logs.Infof("> On update version \r\n")
		if hv, err := strconv.ParseFloat(l.GetVersion(mt.AsModule().Name), 64); nil != err {
			logs.Panicln(err)
		} else {
			mt.AsModule().OnUpdate(hv)
		}
	}
}

// doInit 模块初始化
func (evt *startupListen) doInit(m interface{}, l Loader) {
	mt := m.(Module)
	if nil != mt.AsModule().OnInit {
		logs.Infof("> On init module \r\n")
		mt.AsModule().OnInit()
	}
}
