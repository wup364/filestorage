// Copyright (C) 2019 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 模块加载器 - Test
// 示例加载了 DemoModule, HelloModule两个模块, DemoModule 自动注入依赖的HelloModule模块
package loader

import (
	"pakku/utils/logs"
	"pakku/utils/strutil"
	"testing"
)

type SayHello interface {
	Say()
}

// HelloModule 示例模块
type HelloModule struct {
}

// AsModule 作为一个模块加载
func (t *HelloModule) AsModule() Opts {
	return Opts{
		Name:        "HelloModule",
		Version:     1.0,
		Description: "Hello模板",
		OnReady: func(mctx Loader) {
			mctx.SetParam("HelloModule.Hello_Text", "onReady func set")
		},
		OnSetup: func() {

		},
		OnUpdate: func(cv float64) {

		},
		OnInit: func() {

		},
	}
}

// Hello Hello
func (t *HelloModule) Say() {
	logs.Infoln("HelloModule -> Hello")
}

// DemoModule 示例模块
type DemoModule struct {
	hello SayHello `@autowired:"HelloModule"`
}

// AsModule 作为一个模块加载
func (t *DemoModule) AsModule() Opts {
	return Opts{
		Name:        "DemoModule",
		Version:     1.0,
		Description: "示例模板",
		OnReady: func(mctx Loader) {
			logs.Infoln(mctx.GetParam("HelloModule.Hello_Text").ToString(""))
		},
		OnSetup: func() {

		},
		OnUpdate: func(cv float64) {

		},
		OnInit: func() {

		},
	}
}

// Hello Hello
func (t *DemoModule) SayHello() {
	t.hello.Say()
}

// 在 mian 中调用
func TestLoader(t *testing.T) {
	loader := New("Test")
	loader.OnModuleEvent(new(HelloModule).AsModule().Name, ModuleEventOnReady, func(_ interface{}, loader Loader) {
		logs.Infoln("OnEvent: old=", loader.GetParam("HelloModule.Hello_Text").ToString(""))
		loader.SetParam("HelloModule.Hello_Text", strutil.GetRandom(6))
	})
	loader.Loads(new(HelloModule), new(DemoModule))
	loader.Invoke("DemoModule", "SayHello")
}
