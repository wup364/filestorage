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
	"bufio"
	"fmt"
	"io"
	"os"
	"pakku/utils/fileutil"
	"pakku/utils/logs"
	"pakku/utils/reflectutil"
	"path/filepath"
	"strings"
)

func main() {
	defer func() {
		ResetConsoleColor()
		if err := recover(); nil != err {
			logs.Errorln(err)
			PrintErrAndExit("程序运行异常:", err)
		}
	}()
	// 基本设置 & 登录
	settings()
	session := doLogin()
	cmdexec := NewShellCmdExecutor()
	// 就绪
	cmdexec.SetSession(session)
	reader := bufio.NewReader(os.Stdin)
	ResetConsoleColor()
	Printf(BuildColorText(TextColor.Reset, "[%s]$ "), session.GetShowInfo())
	for {
		if command, err := reader.ReadString('\n'); err == nil && len(command) > 0 {
			canprint := false
			if canprint, err = runCommand(cmdexec, command); err != nil {
				PrintErrln(err.Error())
			}
			if canprint {
				ResetConsoleColor()
				Printf(BuildColorText(TextColor.Reset, "[%s]$ "), cmdexec.session.GetShowInfo())
			}
		} else if nil != err && err != io.EOF {
			logs.Errorln(err)
			PrintErrAndExit(err.Error())
		}
	}
}

// 基础设置
func settings() {
	Println("VERSION: 1.0.0")
	if userHome, err := os.UserHomeDir(); nil != err {
		logs.Panicln(err)
	} else {
		logPath := filepath.Clean(userHome + "/shelltools.log")
		if w, err := fileutil.GetWriter(logPath); nil == err {
			Println("LOGPATH: ", logPath)
			logs.SetOutput(w)
		} else {
			PrintErrAndExit(err.Error())
		}
	}
}

// 登录
func doLogin() *ShellSession {
	if len(os.Args) < 2 {
		PrintErrAndExit("请输入连接地址, 如: shelltool OPENAPI@192.168.2.201")
	}
	//
	logs.Infoln("doLogin address:", os.Args[1])
	conn := strings.Split(os.Args[1], "@")
	if len(conn) != 2 {
		PrintErrAndExit("请输入正确的连接地址, 如: shelltool OPENAPI@192.168.2.201")
	}
	//
	if len(os.Args) > 2 && os.Args[2] == "-debug" {
		logs.SetLoggerLevel(logs.DEBUG)
	} else {
		logs.SetLoggerLevel(logs.INFO)
	}
	//
	session := NewShellSession(conn[1], conn[0])
	for {
		// 输入登录密码
		ResetConsoleColor()
		Print("Passwd: ", CursorControl.Hidden, TextStyle.Hidden)
		reader := bufio.NewReader(os.Stdin)
		if command, err := reader.ReadString('\n'); err == nil {
			ResetConsoleColor()
			Print(CleanerControl.EL3, "> Login please wait...\n")

			command = strings.TrimSuffix(command, "\n")
			commands := strings.Fields(command)
			passwd := ""
			if len(commands) > 0 {
				passwd = commands[0]
			}
			if err := session.Login(passwd); nil != err {
				PrintErrln(err.Error())
			} else {
				break
			}
		} else if nil != err {
			ResetConsoleColor()
			logs.Error(err)
			return nil
		}
	}
	return session
}

// runCommand 执行命令行
func runCommand(cmdexec *ShellCmdExecutor, command string) (bool, error) {
	logs.Debugln("command input:", command)
	command = strings.TrimSuffix(command, "\n")
	commands := strings.Fields(command)
	if len(commands) == 0 {
		return true, nil
	}
	if _, ok := cmdexec.GetHelpInfo()[commands[0]]; !ok {
		return true, fmt.Errorf("不支持的命令 '%s' 输入'help'以获取支持列表", commands[0])
	}
	// exec
	go func() {
		if ires := reflectutil.Invoke(cmdexec, strings.ToUpper(commands[0]), commands[1:]); len(ires) > 0 && !ires[0].IsNil() {
			if err := ires[0].Interface().(error); nil != err {
				PrintErrln(err.Error())
			}
		}
		ResetConsoleColor()
		Printf("\033[39m[%s]$ ", cmdexec.session.GetShowInfo())
	}()
	return false, nil
}
