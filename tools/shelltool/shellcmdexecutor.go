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
	"errors"
	"fmt"
	"opensdk"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/strutil"
)

// NewShellCmdExecutor 各种命令行指令
func NewShellCmdExecutor() *ShellCmdExecutor {
	return &ShellCmdExecutor{
		helpInfo: map[string][]string{
			"help":       {"100", "help", "获取命令行支持列表"},
			"pwd":        {"101", "pwd", "获取当前所在目录"},
			"exit":       {"102", "exit", "退出终端"},
			"cls":        {"103", "cls", "清屏"},
			"_s_p_01":    {"199", "", "文件基础操作:"},
			"cd":         {"201", "cd", "进入某个目录: cd /usr/mydir"},
			"mkdir":      {"202", "mkdir", "新建文件夹, 在当前目录下创建: mkdir newdir 在其他目录创建文件夹: mkdir /usr/newdir"},
			"ls":         {"203", "ls", "列出当前目录下的文件和文件夹(基本信息), ls 远程目录 [limit, offset] 全部显示: ls 分页显示: ls 10 0"},
			"ll":         {"204", "ll", "列出当前目录下的文件和文件夹(详细信息), ll 远程目录 [limit, offset] 全部显示: ll 分页显示: ll 10 0"},
			"mv":         {"205", "mv", "移动或重命名文件或文件夹 重命名: mv /usr/mydir mydir1 移动: mv /usr/mydir /opts/mydir"},
			"cp":         {"206", "cp", "复制文件或文件夹, 如: cp /usr/mydir /opts/mydir"},
			"rm":         {"207", "rm", "删除文件或文件夹: rm /usr/mydir"},
			"put":        {"208", "put", "上传文件或文件夹到当前目录, put 本地目录 [-override(覆盖已存在)] 上传并覆盖已存在: put /home/user -override"},
			"get":        {"209", "get", "下载文件或文件夹到本地目录, get 远程目录 [本地目录, [-override(覆盖已存在)]] 下载并覆盖已存在: get /usr/mydir ./ -override"},
			"_s_p_02":    {"299", "", "服务管理操作(需要NAMENODE角色权限):"},
			"listusers":  {"301", "listusers", "列出所有用户"},
			"createuser": {"302", "createuser", "创建用户, createuser user name pwd type(0=nameNode,1=dataNode,2=openApi) 如: 创建admin账户: createuser admin 管理员 xxxxxx 0"},
			"updatepwd":  {"303", "updatepwd", "更新用户密码, updatepwd user pwd 如: 更新OPENAPI的密码: updatepwd OPENAPI xxxxxx"},
			"deleteuser": {"304", "deleteuser", "删除用户, deleteuser user 如: 删除admin账户: deleteuser admin"},
		},
	}
}

// ShellCmdExecutor 各种命令行指令
type ShellCmdExecutor struct {
	session  *ShellSession
	helpInfo map[string][]string
}

// SetSession 设置会话
func (s *ShellCmdExecutor) SetSession(session *ShellSession) {
	s.session = session
}

// GetHelpInfo 获取帮助列表
func (s *ShellCmdExecutor) GetHelpInfo() map[string][]string {
	return s.helpInfo
}

// GetHelpDesc 获取帮助提示
func (s *ShellCmdExecutor) GetHelpDesc(key string) string {
	return s.helpInfo[key][2]
}

// CLS 清屏
func (s *ShellCmdExecutor) CLS(cmd []string) error {
	Print(CleanerControl.EL3)
	return nil
}

// EXIT 退出终端
func (s *ShellCmdExecutor) EXIT(cmd []string) error {
	os.Exit(0)
	return nil
}

// HELP 获取命令行支持列表
func (s *ShellCmdExecutor) HELP(cmd []string) error {
	max := 6
	for key := range s.helpInfo {
		if max < len(key) {
			max = len(key)
		}
	}
	items := (&HelpInfoItemSort{}).Sort(s.helpInfo)
	for i := 0; i < len(items); i++ {
		if len(items[i][1]) > 0 {
			Printfln("  %-"+strconv.Itoa(max)+"s %s", items[i][1], items[i][2])
		} else {
			Printfln("  %s", items[i][2])
		}
	}
	return nil
}

// PWD 获取当前所在目录
func (s *ShellCmdExecutor) PWD(cmd []string) error {
	Println(s.session.GetcurrentDir())
	return nil
}

// CD 进入某个目录, 如: cd /usr/mydir
func (s *ShellCmdExecutor) CD(cmd []string) error {
	if len(cmd) == 0 {
		return fmt.Errorf("参数不能为空, %s", s.GetHelpDesc("cd"))
	}
	path := s.formatRemotePath(cmd[0])
	// 查询节点是否存在
	if node, err := s.session.openApi.GetNode(path); err != nil {
		return err
	} else if nil == node {
		return fmt.Errorf("目录不存在, %s", path)
	} else if node.Flag == 1 {
		return fmt.Errorf("%s 不是一个目录", path)
	}
	s.session.SetcurrentDir(path)
	return nil
}

// LL 列出当前目录下的文件和文件夹, 显示详细信息
func (s *ShellCmdExecutor) LL(cmd []string) error {
	limit := -1
	offset := -1
	path := s.session.GetcurrentDir()
	if len(cmd) == 2 {
		if val, err := strconv.ParseInt(cmd[0], 10, 32); nil == err {
			limit = int(val)
		} else {
			return fmt.Errorf("参数错误, %s", s.GetHelpDesc("11"))
		}
		if val, err := strconv.ParseInt(cmd[1], 10, 32); nil == err {
			offset = int(val)
		} else {
			return fmt.Errorf("参数错误, %s", s.GetHelpDesc("11"))
		}
	}
	if res, err := s.session.openApi.GetDirNodeList(path, limit, offset); nil != err {
		return err
	} else if res == nil || res.Total == 0 {
		Println("- 空 -")
	} else {
		Println("[ ID\tCtime\tMtime\tSize\tName\tProps ]")
		for i := 0; i < len(res.Datas); i++ {
			node := res.Datas[i]
			if node.Flag == 0 {
				node.Name = BuildColorText(TextColor.Blue, node.Name)
			} else if node.Flag == 1 {
				node.Name = BuildColorText(TextColor.Green, node.Name)
			}
			ctime := ""
			if node.Ctime > 0 {
				ctime = time.UnixMilli(node.Ctime).Format("2006-01-02 15:04:05")
			}
			mtime := ""
			if node.Mtime > 0 {
				mtime = time.UnixMilli(node.Mtime).Format("2006-01-02 15:04:05")
			}
			Println(fmt.Sprintf("%s %s %s %8s %s %s", node.Id, ctime, mtime, fileutil.FormatFileSize(node.Size), node.Name, node.Props))
		}
		Println(fmt.Sprintf("[ Total=%d  Limit=%d Offset=%d ]", res.Total, limit, offset))
	}
	return nil
}

// LS 列出当前目录下的文件和文件夹, 显示基本信息
func (s *ShellCmdExecutor) LS(cmd []string) error {
	limit := -1
	offset := -1
	path := s.session.GetcurrentDir()
	if len(cmd) == 2 {
		if val, err := strconv.ParseInt(cmd[0], 10, 32); nil == err {
			limit = int(val)
		} else {
			return fmt.Errorf("参数错误, %s", s.GetHelpDesc("ls"))
		}
		if val, err := strconv.ParseInt(cmd[1], 10, 32); nil == err {
			offset = int(val)
		} else {
			return fmt.Errorf("参数错误, %s", s.GetHelpDesc("ls"))
		}
	}
	if res, err := s.session.openApi.GetDirNodeList(path, limit, offset); nil != err {
		return err
	} else if res == nil || res.Total == 0 {
		Println("- 空 -")
	} else {
		Println("[ Mtime\tSize\tName ]")
		for i := 0; i < len(res.Datas); i++ {
			node := res.Datas[i]
			if node.Flag == 0 {
				node.Name = BuildColorText(TextColor.Blue, node.Name)
			} else if node.Flag == 1 {
				node.Name = BuildColorText(TextColor.Green, node.Name)
			}
			mtime := ""
			if node.Mtime > 0 {
				mtime = time.UnixMilli(node.Mtime).Format("2006-01-02 15:04:05")
			}
			Println(fmt.Sprintf("%s %8s %s", mtime, fileutil.FormatFileSize(node.Size), node.Name))
		}
		Println(fmt.Sprintf("[ Total=%d  Limit=%d Offset=%d ]", res.Total, limit, offset))
	}
	return nil
}

// MKDIR 在当前目录下新建文件夹
func (s *ShellCmdExecutor) MKDIR(cmd []string) error {
	if len(cmd) < 1 {
		return fmt.Errorf("参数不正确, %s", s.GetHelpDesc("mkdir"))
	}
	for i := 0; i < len(cmd); i++ {
		src := s.formatRemotePath(cmd[i])
		if _, err := s.session.openApi.DoMkDir(src); nil != err {
			return err
		}
	}
	return nil
}

// MV 移动或重命名文件或文件夹, 如:\n 重命名: mv /usr/mydir mydir1 \n 移动: mv /usr/mydir /opts/mydir
func (s *ShellCmdExecutor) MV(cmd []string) error {
	if len(cmd) < 2 {
		return fmt.Errorf("参数不正确, %s", s.GetHelpDesc("mv"))
	}
	src := s.formatRemotePath(cmd[0])
	dest := strutil.Parse2UnixPath(cmd[1])
	if isRename := !strings.Contains(dest, "/"); isRename {
		return s.session.openApi.DoRename(src, dest)
	} else {
		return s.session.openApi.DoMove(src, dest, false)
	}
}

// CP 复制文件或文件夹, 如: cp /usr/mydir /opts/mydir
func (s *ShellCmdExecutor) CP(cmd []string) error {
	if len(cmd) < 2 {
		return fmt.Errorf("参数不正确, %s", s.GetHelpDesc("mv"))
	}
	isMulti := len(cmd) > 2
	dest := s.formatRemotePath(cmd[len(cmd)-1])
	for i := 0; i < len(cmd)-1; i++ {
		desttmp := dest
		src := s.formatRemotePath(cmd[i])
		if len(src) == 0 || src == "/" {
			return errors.New("不能拷贝根目录")
		}
		if isMulti {
			desttmp = dest + "/" + strutil.GetPathName(cmd[i])
		}
		if err := copyNodes(src, desttmp, s.session.openApi); nil != err {
			return err
		}
	}
	return nil
}

// RM 删除文件或文件夹, 如: rm /usr/mydir
func (s *ShellCmdExecutor) RM(cmd []string) error {
	if len(cmd) < 1 {
		return fmt.Errorf("参数不正确, %s", s.GetHelpDesc("rm"))
	}
	for i := 0; i < len(cmd); i++ {
		src := s.formatRemotePath(cmd[i])
		if len(src) == 0 || src == "/" {
			return errors.New("不能删除或移动根目录")
		}
		if err := s.session.openApi.DoDelete(src); nil != err {
			return err
		}
	}
	return nil
}

// PUT 上传文件或文件夹到当前目录, put 本地目录 [-override(覆盖已存在)] 如:\n 上传并覆盖已存在: put /home/user -override
func (s *ShellCmdExecutor) PUT(cmd []string) error {
	if len(cmd) < 1 {
		return fmt.Errorf("参数不正确, %s", s.GetHelpDesc("put"))
	}
	override := false
	if len(cmd) > 1 {
		override = cmd[1] == "-override"
	}
	src := s.formatLoclPath(cmd[0])
	Printfln("upload: src=%s, dest=%s, override=%v", src, s.session.GetcurrentDir(), override)
	return doUploadDir(s.session.GetcurrentDir(), src, src, override, s.session.openApi)
}

// GET 下载文件或文件夹到当前目录, get 远程目录 [本地目录, [-override(覆盖已存在)]] 如:\n 下载并覆盖已存在: get /usr/mydir ./ -override
func (s *ShellCmdExecutor) GET(cmd []string) error {
	if len(cmd) < 1 {
		return fmt.Errorf("参数不正确, %s", s.GetHelpDesc("get"))
	}
	dest := "./"
	override := false
	if len(cmd) > 1 {
		if len(cmd) < 3 {
			dest = cmd[1]
		} else {
			dest = cmd[1]
			override = cmd[2] == "-override"
		}
	}
	src := s.formatRemotePath(cmd[0])
	if node, err := s.session.openApi.GetNode(src); nil != err {
		return err
	} else if node.Flag == 0 {
		dest = s.formatLoclPath(dest)
		Printfln("download: src=%s, dest=%s, override=%v", src, dest, override)
		return doDownloadDir(dest, src, src, override, s.session.openApi)
	} else {
		if dest == "./" {
			dest = s.formatLoclPath(dest + "/" + node.Name)
		}
		Printfln("download: src=%s, dest=%s, override=%v", src, dest, override)
		return doDownloadFile(dest, src, override, s.session.openApi)
	}
}

// LISTUSERS 列出所有用户(需要NAMENODE角色权限)
func (s *ShellCmdExecutor) LISTUSERS(cmd []string) error {
	if res, err := s.session.serverMG.DoListAllUsers(); nil != err {
		return err
	} else if res == nil || res.Total == 0 {
		Println("- 空 -")
	} else {
		Println("[ Type\tUser\tName\tCtime ]")
		for i := 0; i < len(res.Datas); i++ {
			user := res.Datas[i]
			if user.UserType == 0 {
				user.UserID = BuildColorText(TextColor.Magenta, user.UserID)
			} else if user.UserType == 1 {
				user.UserID = BuildColorText(TextColor.Blue, user.UserID)
			} else if user.UserType == 2 {
				user.UserID = BuildColorText(TextColor.Green, user.UserID)
			}
			ctime := user.CtTime.Format("2006-01-02 15:04:05")
			Println(fmt.Sprintf("%d %s %s %s", user.UserType, user.UserID, user.UserName, ctime))
		}
		Println(fmt.Sprintf("[ Total=%d  Limit=%s Offset=%d ]", res.Total, "max", 0))
	}
	return nil
}

// CREATEUSER 创建用户(需要NAMENODE角色权限), createuser user name pwd type(0=nameNode,1=dataNode,2=openApi) 如: 创建admin账户: createuser admin 管理员 xxxxxx 0
func (s *ShellCmdExecutor) CREATEUSER(cmd []string) error {
	if len(cmd) < 4 {
		return fmt.Errorf("参数不正确, %s", s.GetHelpDesc("createuser"))
	}
	userType := strutil.String2Int(cmd[3], -1)
	if userType == -1 {
		return fmt.Errorf("参数不正确, %s", s.GetHelpDesc("createuser"))
	}

	if ok, err := s.session.serverMG.DoCreateUser(opensdk.CreateUserBo{
		UserType: userType,
		UserID:   cmd[0],
		UserName: cmd[1],
		UserPWD:  cmd[2],
	}); nil != err {
		return err
	} else if !ok {
		return errors.New("操作失败")
	}
	return nil
}

// UPDATEPWD 更新用户密码(需要NAMENODE角色权限), updatewpd user pwd 如:\n 跟新OPENAPI的密码: updatepwd OPENAPI xxxxxx
func (s *ShellCmdExecutor) UPDATEPWD(cmd []string) error {
	if len(cmd) < 1 {
		return fmt.Errorf("参数不正确, %s", s.GetHelpDesc("updatepwd"))
	}
	userAccount := cmd[0]
	if len(userAccount) == 0 {
		return fmt.Errorf("参数不正确, %s", s.GetHelpDesc("updatepwd"))
	}
	inputPwd := ""
	if len(cmd) > 1 {
		inputPwd = cmd[1]
	}
	if ok, err := s.session.serverMG.DoUpdatePWD(userAccount, inputPwd); ok {
		Printfln("用户 %s 密码已修改", userAccount)
	} else {
		return err
	}
	return nil
}

// DELETEUSER 删除用户(需要NAMENODE角色权限), deleteuser user 如: 删除admin账户: deleteuser admin
func (s *ShellCmdExecutor) DELETEUSER(cmd []string) error {
	if len(cmd) < 1 {
		return fmt.Errorf("参数不正确, %s", s.GetHelpDesc("deleteuser"))
	}
	if ok, err := s.session.serverMG.DoDeleteUser(cmd[0]); ok {
		Printfln("用户 %s 已删除", cmd[0])
	} else {
		return err
	}
	return nil
}

// formatRemotePath 转换服务器路径参数
func (s *ShellCmdExecutor) formatRemotePath(src string) string {
	src = strutil.Parse2UnixPath(src)
	if !strings.Contains(src, "/") || strings.HasPrefix(src, "./") || strings.HasPrefix(src, "../") {
		src = strutil.Parse2UnixPath(s.session.GetcurrentDir() + "/" + src)
	}
	return path.Clean(src)
}

// formatLoclPath 转换本地路径参数
func (s *ShellCmdExecutor) formatLoclPath(src string) string {
	if src = filepath.Clean(src); !filepath.IsAbs(src) {
		if abs, err := filepath.Abs(src); nil == err {
			src = abs
		}
	}
	return filepath.Clean(src)
}

type HelpInfoItemSort struct {
	items [][]string
}

// Sort 排序
func (s *HelpInfoItemSort) Sort(helpInfo map[string][]string) [][]string {
	s.items = make([][]string, 0)
	if len(helpInfo) > 0 {
		for _, v := range helpInfo {
			s.items = append(s.items, v)
		}
	}
	if len(s.items) > 0 {
		sort.Sort(s)
	}
	return s.items
}

// 实现sort.Interface接口取元素数量方法
func (s *HelpInfoItemSort) Len() int {
	return len(s.items)
}

// 实现sort.Interface接口比较元素方法
func (s *HelpInfoItemSort) Less(i, j int) bool {
	si := strutil.String2Int(s.items[i][0], 0)
	sj := strutil.String2Int(s.items[j][0], 0)
	return si < sj
}

// 实现sort.Interface接口交换元素方法
func (s *HelpInfoItemSort) Swap(i, j int) {
	s.items[i], s.items[j] = s.items[j], s.items[i]
}
