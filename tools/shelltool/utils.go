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
	"os"
	"strings"

	"github.com/mattn/go-colorable"
)

// 输出流
var sdtout = colorable.NewColorable(os.Stdout)
var sdterr = colorable.NewColorable(os.Stderr)

// BuildColorText 构建一段带颜色的字符
func BuildColorText(color TMCode, text string) string {
	return string(color) + text + string(TextColor.Reset)
}

// PrintErrln 打印错误信息
func PrintErrln(text ...interface{}) {
	if len(text) > 0 {
		swap := make([]interface{}, 0)
		swap = append(append(swap, TextColor.Red), text...)
		swap = append(swap, TextStyle.Reset, TextColor.Reset, BackgroundColor.Reset, CursorControl.Visible)
		fmt.Fprintln(sdterr, swap...)
	}
}

// PrintErrAndExit 打印错误信息, 并退出程序
func PrintErrAndExit(text ...interface{}) {
	PrintErrln(text...)
	os.Exit(0)
}

// ResetConsoleColor 重置控制台颜色
func ResetConsoleColor() {
	Print(TextStyle.Reset, TextColor.Reset, BackgroundColor.Reset, CursorControl.Visible)
}

// Print 打印普通信息
func Printf(text string, args ...interface{}) {
	fmt.Fprintf(sdtout, text, args...)
}

// Println 打印普通信息
func Printfln(text string, args ...interface{}) {
	if !strings.HasSuffix(text, "\n") {
		text += "\n"
	}
	fmt.Fprintf(sdtout, text, args...)
}

// Print 打印普通信息
func Print(text ...interface{}) {
	fmt.Fprint(sdtout, text...)
}

// Println 打印普通信息
func Println(text ...interface{}) {
	fmt.Fprintln(sdtout, text...)
}
