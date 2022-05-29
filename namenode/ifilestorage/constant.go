// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package ifilestorage

import "fmt"

const (
	StreamTokenType_Read  = 0
	StreamTokenType_Write = 1
)

// StreamTokenType token类型
type StreamTokenType int

// DBSetting DBSetting
type DBSetting struct {
	DriverName     string
	DataSourceName string
}

// StreamToken 流操作Token
type StreamToken struct {
	Token    string
	NodeNo   string
	FileID   string
	FilePath string
	FileSize int64
	CTime    int64
	MTime    int64
	EndPoint string
	Type     StreamTokenType
}

// Clone Clone
func (ua *StreamToken) Clone(val interface{}) error {
	if st, ok := val.(*StreamToken); ok {
		st.Token = ua.Token
		st.NodeNo = ua.NodeNo
		st.FileID = ua.FileID
		st.FilePath = ua.FilePath
		st.FileSize = ua.FileSize
		st.CTime = ua.CTime
		st.MTime = ua.MTime
		st.EndPoint = ua.EndPoint
		st.Type = ua.Type
		return nil
	}
	return fmt.Errorf("can't support clone %T ", val)
}
