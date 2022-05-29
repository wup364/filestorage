package main

// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

import (
	"fmt"
	"opensdk"
	"os"
	"pakku/modules/appconfig/jsonconfig"
	"pakku/utils/strutil"
	"strings"
)

// NewShellSession 建立新会话
func NewShellSession(server, userName string) *ShellSession {
	session := &ShellSession{
		server:   server,
		userName: userName,
		conf:     new(jsonconfig.Config),
	}
	if err := session.conf.Init("shelltool"); nil != err {
		panic(err)
	}
	return session
}

// ShellSession 登录会话
type ShellSession struct {
	server     string
	userName   string
	currentdir string
	conf       *jsonconfig.Config
	openApi    opensdk.IOpenApi
	serverMG   opensdk.IServerMG
}

// GetShowInfo 获取命令行显示信息, username pathName
func (s *ShellSession) GetShowInfo() string {
	if len(s.currentdir) == 0 || s.currentdir == "/" {
		return s.userName + " /"
	}
	return s.userName + " " + strutil.GetPathName(s.currentdir)
}

// GetcurrentDir 获取当前路径
func (s *ShellSession) GetcurrentDir() string {
	if len(s.currentdir) == 0 || s.currentdir == "/" {
		return "/"
	}
	return strutil.Parse2UnixPath(s.currentdir)
}

// SetcurrentDir 获取登录名
func (s *ShellSession) SetcurrentDir(path string) {
	if len(path) == 0 {
		path = "/"
	}
	s.currentdir = strutil.Parse2UnixPath(path)
}

// GetOpenApi GetOpenApi
func (s *ShellSession) GetOpenApi() opensdk.IOpenApi {
	return s.openApi
}

// Login 登录
func (s *ShellSession) Login(passwd string) error {
	if !strings.Contains(s.server, ":") {
		s.server += ":5051"
	}
	// openapi
	s.openApi = opensdk.NewOpenApi(s.server, opensdk.User{
		User:   s.userName,
		Passwd: passwd,
	})
	if err := s.openApi.GetRPCConn().DoLogin(); nil != err {
		return err
	}
	// server manager
	s.serverMG = opensdk.NewSvrMngApi(s.server, opensdk.User{
		User:   s.userName,
		Passwd: passwd,
	})
	// datanode 路由
	if datanodes := s.conf.GetConfig("datanodes").ToStrMap(nil); len(datanodes) > 0 {
		dnodedns := make(map[string]string)
		for key, val := range datanodes {
			dnodedns[key] = val.(string)
		}
		s.openApi.SetDataNodeDNS(dnodedns)
		fmt.Fprintf(os.Stdout, "> Setting DataNode DNS(datanodes): %v \n", dnodedns)
	}
	// 输出信息
	fmt.Fprintf(os.Stdout, "> Login succeed: %s@%s \n", s.userName, s.server)
	return nil
}
