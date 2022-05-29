// Copyright (C) 2019 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 模块加载器
// 依赖包: utils.utypes.Object utils.strutil.strutil

package loader

import (
	"fmt"
	"pakku/utils/logs"
	"pakku/utils/reflectutil"
	"pakku/utils/strutil"
	"pakku/utils/utypes"
	"reflect"
	"strconv"
	"time"
)

const (
	VAL_APPNAME = "app"
	KEY_APPNAME = "app.name"
)

// New 实例一个加载器对象
func New(appName string) Loader {
	loader := &MLoader{
		events:     utypes.NewSafeMap(),
		modules:    utypes.NewSafeMap(),
		mparams:    utypes.NewSafeMap(),
		instanceID: strutil.GetUUID(),
	}
	if len(appName) == 0 {
		appName = VAL_APPNAME
	}
	loader.SetParam(KEY_APPNAME, appName)
	loader.SetModuleInfoHandler(new(InfoRecorder))
	new(startupListen).bind(loader)
	return loader
}

// MLoader 模块加载器, 实例化后可实现统一管理模板
type MLoader struct {
	instanceID string          // 加载器实例ID
	events     *utypes.SafeMap // 模块生命周期事件
	modules    *utypes.SafeMap // 模块Map表
	mparams    *utypes.SafeMap // 保存在模块对象中共享的字段key-value
	mrecord    ModuleInfo      // 模块信息记录器
}

// Loads 初始化模块 - DO Setup -> Check Ver -> Do Init
func (loader *MLoader) Loads(mts ...Module) {
	for _, mt := range mts {
		loader.Load(mt)
	}
}

// Load 初始化模块 -- doready -> doSetup -> doCheckVersion -> doInit -> doEnd
func (loader *MLoader) Load(mt Module) {
	logs.Infof("> Loading %s(%s)[%p] start \r\n", mt.AsModule().Name, mt.AsModule().Description, mt)
	// doready 模块准备开始加载
	loader.doHandleModuleEvent(mt, ModuleEventOnReady)
	// doSetup 模块安装
	if len(loader.GetVersion(mt.AsModule().Name)) == 0 {
		loader.doHandleModuleEvent(mt, ModuleEventOnSetup)
		loader.setVersion(mt.AsModule())
		// doCheckVersion 模块升级
	} else if loader.GetVersion(mt.AsModule().Name) != strconv.FormatFloat(mt.AsModule().Version, 'f', 2, 64) {
		loader.doHandleModuleEvent(mt, ModuleEventOnUpdate)
		loader.setVersion(mt.AsModule())
	}
	// doInit 模块初始化
	loader.doHandleModuleEvent(mt, ModuleEventOnInit)
	// doEnd 模块加载结束
	loader.modules.Put(mt.AsModule().Name, mt)
	logs.Infof("> Loading %s complete \r\n", mt.AsModule().Name)
	loader.doHandleModuleEvent(mt, ModuleEventOnLoaded)
}

// Invoke 模块调用, 返回 []reflect.Value, 返回值暂时无法处理
func (loader *MLoader) Invoke(name string, method string, params ...interface{}) ([]reflect.Value, error) {
	if module, ok := loader.modules.Get(name); ok {
		val := reflect.ValueOf(module)
		fun := val.MethodByName(method)
		logs.Infof("> Invoke: "+name+"."+method+", %v, %+v \r\n", fun, &fun)
		args := make([]reflect.Value, len(params))
		for i, temp := range params {
			args[i] = reflect.ValueOf(temp)
		}
		return fun.Call(args), nil
	}
	return nil, fmt.Errorf(ErrModuleNotFoundStr, name)
}

// AutoWired 自动注入依赖对象
func (loader *MLoader) AutoWired(structobj interface{}) error {
	return doAutoWired(structobj, loader)
}

// GetInstanceID 获取实例的ID
func (loader *MLoader) GetInstanceID() string {
	return loader.instanceID
}

// SetParam 设置变量, 保存在模板加载器实例内部
func (loader *MLoader) SetParam(key string, val interface{}) {
	loader.mparams.Put(key, val)
}

// GetParam 模板加载器实例上的变量
func (loader *MLoader) GetParam(key string) utypes.Object {
	if val, ok := loader.mparams.Get(key); ok {
		return utypes.NewObject(val)
	}
	return utypes.Object{}
}

// OnModuleEvent 监听模块生命周期事件
func (loader *MLoader) OnModuleEvent(name string, event ModuleEvent, val OnModuleEvent) {
	var events []OnModuleEvent
	if val, ok := loader.events.Get("ModuleEvent." + name + "." + string(event)); ok {
		events = val.([]OnModuleEvent)
	} else {
		events = make([]OnModuleEvent, 0)
	}
	loader.events.Put("ModuleEvent."+name+"."+string(event), append(events, val))
}

// GetModuleByName 根据模块Name获取模块指针记录, 可以获取一个已经实例化的模块
func (loader *MLoader) GetModuleByName(name string, val interface{}) error {
	if tmp, ok := loader.modules.Get(name); ok {
		return reflectutil.SetInterfaceValueUnSafe(val, tmp)
	}
	return fmt.Errorf(ErrModuleNotFoundStr, name)
}

// SetModuleInfoHandler 设置模块信息记录器, 会自动调用init
func (loader *MLoader) SetModuleInfoHandler(mrecord ModuleInfo) {
	if nil != mrecord {
		loader.mrecord = mrecord
		err := loader.mrecord.Init(loader.GetParam(KEY_APPNAME).ToString(VAL_APPNAME))
		if nil != err {
			panic(err)
		}
	}
}

// GetVersion 获取模块版本号
func (loader *MLoader) GetVersion(name string) string {
	return loader.mrecord.GetValue(name + ".SetupVer")
}

// setVersion 设置模块版本号 - 模块保留小数两位
func (loader *MLoader) setVersion(opts Opts) {
	loader.mrecord.SetValue(opts.Name+".SetupVer", strconv.FormatFloat(opts.Version, 'f', 2, 64))
	loader.mrecord.SetValue(opts.Name+".SetupDate", strconv.FormatInt(time.Now().UnixNano(), 10))
}

// doHandleModuleEvent 执行监听模块生命周期事件
func (loader *MLoader) doHandleModuleEvent(mt Module, event ModuleEvent) {
	var events []OnModuleEvent
	if val, ok := loader.events.Get("ModuleEvent.*." + string(event)); ok {
		if funs, ok := val.([]OnModuleEvent); ok {
			events = funs
		}
	}
	if val, ok := loader.events.Get("ModuleEvent." + mt.AsModule().Name + "." + string(event)); ok {
		if funs, ok := val.([]OnModuleEvent); ok {
			events = append(events, funs...)
		}
	}
	if len(events) > 0 {
		for i := 0; i < len(events); i++ {
			events[i](mt, loader)
		}
	}
}
