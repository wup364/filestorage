package main

// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 终端|控制台特殊指令集

// TMCode 终端特殊指令类型
type TMCode string

// TextColor 文字颜色
var TextColor = _TextColor{
	Black:   "\033[30m", // 黑色
	Red:     "\033[31m", // 红色
	Green:   "\033[32m", // 绿色
	Yellow:  "\033[33m", // 黄色
	Blue:    "\033[34m", // 蓝色
	Magenta: "\033[35m", // 品红
	Cyan:    "\033[36m", // 青色
	White:   "\033[37m", // 白色
	Reset:   "\033[39m", // 重置默认
}

// BackgroundColor 背景颜色
var BackgroundColor = _BgColor{
	Black:   "\033[40m", // 黑色
	Red:     "\033[41m", // 红色
	Green:   "\033[42m", // 绿色
	Yellow:  "\033[43m", // 黄色
	Blue:    "\033[44m", // 蓝色
	Magenta: "\033[45m", // 品红
	Cyan:    "\033[46m", // 青色
	White:   "\033[47m", // 白色
	Reset:   "\033[49m", // 重置默认
}

// cursorControl 光标控制
var CursorControl = _CursorControl{
	Home:    "\033[H",    // 移动光标到原点
	SC:      "\0337",     // 保存当前光标位置
	RC:      "\0338",     // 恢复已保存的光标位置
	Cub1:    "\033?",     // 向左移动一个位置 (退格)
	Hidden:  "\033[?25l", // 光标不可见
	Visible: "\033[?25h", // 光标可见
}

// CleanerControl  删除文本
var CleanerControl = _CleanerControl{
	EL:  "\033[0K", // 清除从当前位置到本行结尾的字符
	EL1: "\033[1K", // 清除从本行开始到当前位置的字符
	EL2: "\033[2K", // 清除本行所有字符 (光标位置不变)
	EL3: "\033[2J", // 清屏 (光标位置不变)
}

// TextStyle 文本属性
var TextStyle = _TextStyle{
	Reset:      "\033[0m", // 重置所有属性
	Bold:       "\033[1m", // 粗体
	Deepen:     "\033[2m", // 加深
	Protrude:   "\033[3m", // 突出
	Underscore: "\033[4m", // 下划线
	Flashing:   "\033[5m", // 闪烁
	Reverse:    "\033[7m", // 倒序
	Hidden:     "\033[8m", // 隐藏
}

// _CleanerControl  删除文本
type _CleanerControl struct {
	EL  TMCode // 清除从当前位置到本行结尾的字符
	EL1 TMCode // 清除从本行开始到当前位置的字符
	EL2 TMCode // 清除本行所有字符 (光标位置不变)
	EL3 TMCode // 清屏 (光标位置不变)
}

// _CursorControl 光标控制
type _CursorControl struct {
	Home    TMCode // 移动光标到原点
	SC      TMCode // 保存当前光标位置
	RC      TMCode // 恢复已保存的光标位置
	Cub1    TMCode // 向左移动一个位置 (退格)
	Hidden  TMCode // 光标不可见
	Visible TMCode // 光标可见
}

// _TextStyle 文本属性
type _TextStyle struct {
	Reset      TMCode // 重置所有属性
	Bold       TMCode // 粗体
	Deepen     TMCode // 加深
	Protrude   TMCode // 突出
	Underscore TMCode // 下划线
	Flashing   TMCode // 闪烁
	Reverse    TMCode // 倒序
	Hidden     TMCode // 隐藏
}

// _TextColor 文字颜色
type _TextColor struct {
	Black   TMCode // 黑色
	Red     TMCode // 红色
	Green   TMCode // 绿色
	Yellow  TMCode // 黄色
	Blue    TMCode // 蓝色
	Magenta TMCode // 品红
	Cyan    TMCode // 青色
	White   TMCode // 白色
	Reset   TMCode // 重置默认
}

// _BgColor 背景颜色
type _BgColor struct {
	Black   TMCode // 黑色
	Red     TMCode // 红色
	Green   TMCode // 绿色
	Yellow  TMCode // 黄色
	Blue    TMCode // 蓝色
	Magenta TMCode // 品红
	Cyan    TMCode // 青色
	White   TMCode // 白色
	Reset   TMCode // 重置默认
}
