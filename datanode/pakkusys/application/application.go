// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 实例化一个pakku运行环境

package application

import (
	"datanode/pakkusys/pakkuconf"
	"errors"
	"path/filepath"
	"reflect"
	"sync"

	"github.com/wup364/pakku"
	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/logs"
)

//pakkuBoot pakkuBoot instance
var pakkuBoot *PakkuBoot

// ApplicationBoot 启动单例应用
func ApplicationBoot(name string) *PakkuBoot {
	if nil == pakkuBoot {
		pakkuBoot = &PakkuBoot{lock: new(sync.Mutex), appBoot: pakku.NewApplication(name)}
	}
	return pakkuBoot
}

// PakkuBoot PakkuBoot
type PakkuBoot struct {
	booted  bool
	lock    *sync.Mutex
	appBoot ipakku.ApplicationBoot
}

// EnableCoreModule 启用默认的核心模块
func (boot *PakkuBoot) EnableCoreModule() *PakkuBoot {
	boot.appBoot.EnableCoreModule()
	return boot
}

// EnableNetModule 启用默认的网络服务模块
func (boot *PakkuBoot) EnableNetModule() *PakkuBoot {
	boot.appBoot.EnableNetModule()
	return boot
}

// SetLogger 初始化默认的本地日志
func (boot *PakkuBoot) SetLogger(logger, logdir, loglevel string) *PakkuBoot {
	if logger == "file" {
		logPath := logdir
		if len(logPath) == 0 {
			logPath = "./logs"
		}
		if !filepath.IsAbs(logPath) {
			if path, err := filepath.Abs(logPath); nil != err {
				logs.Panicln(err)
			} else {
				logPath = path
			}
		}
		if !fileutil.IsExist(logPath) {
			if err := fileutil.MkdirAll(logPath); nil != err {
				logs.Panicln(err)
			}
		}
		logPath = filepath.Clean(logPath + "/" + boot.GetApplication().GetParam(ipakku.PARAMKEY_APPNAME).ToString(boot.GetApplication().GetInstanceID()) + ".log")
		file, err := fileutil.GetWriter(logPath)
		if nil != err {
			logs.Panicln(err)
		}
		logs.Infoln("Logger output file, path: ", logPath)
		boot.appBoot.SetLoggerOutput(file)
	}
	switch loglevel {
	case "none":
		boot.appBoot.SetLoggerLevel(logs.NONE)
	case "error":
		boot.appBoot.SetLoggerLevel(logs.ERROR)
	case "info":
		boot.appBoot.SetLoggerLevel(logs.INFO)
	default:
		boot.appBoot.SetLoggerLevel(logs.DEBUG)
	}
	return boot
}

// GetApplication 获取实例
func (boot *PakkuBoot) GetApplication() ipakku.Application {
	return boot.appBoot.GetApplication()
}

// GetModuleByName 获取模块
func (boot *PakkuBoot) GetModuleByName(name string, val interface{}) {
	if nil == val {
		logs.Panicln(errors.New("value is nil"))
	}
	if nil == pakkuBoot {
		logs.Panicln(errors.New("application is not initialized"))
	}
	if err := boot.GetApplication().GetModuleByName(name, val); nil != err {
		logs.Panicln(err)
	}
}

// GetModule 获取模块, 模块名字和接口名字一样
func (boot *PakkuBoot) GetModule(val ...interface{}) {
	for _, v := range val {
		name := reflect.TypeOf(v).Elem().Name()
		boot.GetModuleByName(name, v)
	}
}

// BootStart 加载&启动程序
func (boot *PakkuBoot) BootStart() ipakku.Application {
	boot.lock.Lock()
	defer boot.lock.Unlock()
	if boot.booted {
		return boot.GetApplication()
	}
	boot.booted = true
	// 加载自定义模块
	if modules := pakkuconf.RegisterModules(); len(modules) > 0 {
		boot.appBoot.AddModules(modules...)
	}
	// 注册覆盖模块
	if overrides := pakkuconf.RegisterOverride(); len(overrides) > 0 {
		for i := 0; i < len(overrides); i++ {
			if overrides[i].Instance != nil {
				ipakku.Override.RegisterInterfaceImpl(overrides[i].Instance, overrides[i].Interface, overrides[i].Implement)
			}
			ipakku.Override.SetInterfaceDefaultImpl(boot.GetApplication(), overrides[i].Interface, overrides[i].Implement)
		}
	}
	// 注册模块加载事件
	if events := pakkuconf.RegisterModuleEvent(); len(events) > 0 {
		for i := 0; i < len(events); i++ {
			boot.GetApplication().OnModuleEvent(events[i].Module, events[i].Event, events[i].Handler)
		}
	}
	return boot.appBoot.BootStart()
}

// BootStartWeb 以web服务方式启动, https证书放在conf目录下自动加载
func (boot *PakkuBoot) BootStartWeb(debug bool) {
	// bootstart
	boot.EnableCoreModule().EnableNetModule().BootStart()
	boot.GetApplication().SetParam("AppService.debug", debug)
	// 获取appservice接口
	var config ipakku.AppConfig
	var service ipakku.AppService
	boot.GetModule(&config, &service)
	address := config.GetConfig("listen.http.address").ToString("127.0.0.1:5062")
	boot.GetApplication().SetParam("listen.http.address", address)
	// 注册controller
	for _, v := range pakkuconf.RegisterController() {
		if err := service.AsController(v); nil != err {
			logs.Panicln(err)
		}
	}
	// 启动服务
	certFile, keyFile := getCertFile()
	service.StartHTTP(ipakku.HTTPServiceConfig{
		ListenAddr: address,
		CertFile:   certFile,
		KeyFile:    keyFile,
	})
}

// BootStartRpc 以Rpc服务方式启动
func (boot *PakkuBoot) BootStartRpc() {
	// bootstart
	boot.EnableCoreModule().EnableNetModule().BootStart()
	// 获取appservice接口
	var config ipakku.AppConfig
	var service ipakku.AppService
	boot.GetModule(&config, &service)
	address := config.GetConfig("listen.rpc.address").ToString("127.0.0.1:5061")
	boot.GetApplication().SetParam("listen.rpc.address", address)
	// 注册rpcservice
	for _, v := range pakkuconf.RegisterRPCService() {
		if err := service.RegisteRPC(v); nil != err {
			logs.Panicln(err)
		}
	}
	service.StartRPC(ipakku.RPCServiceConfig{ListenAddr: address})
}

// GetApplication 获取实例
func GetApplication() ipakku.Application {
	return pakkuBoot.GetApplication()
}

// GetModuleByName 获取模块
func GetModuleByName(name string, val interface{}) {
	pakkuBoot.GetModuleByName(name, val)
}

// GetModule 获取模块, 模块名字和接口名字一样
func GetModule(val ...interface{}) {
	pakkuBoot.GetModule(val...)
}

// BootStartWeb 以web服务方式启动, https证书放在conf目录下自动加载
func BootStartWeb(debug bool) {
	pakkuBoot.BootStartWeb(debug)
}

// BootStartRpc 以Rpc服务方式启动
func BootStartRpc() {
	pakkuBoot.BootStartRpc()
}

// getCertFile 获取https证书
func getCertFile() (certFile string, keyFile string) {
	if fileutil.IsFile("./.conf/key.pem") {
		certFile = "./.conf/key.pem"
	}
	if fileutil.IsFile("./.conf/cert.pem") {
		certFile = "./.conf/cert.pem"
	}
	return
}
