// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package opensdk

import (
	"encoding/gob"
	"fmt"
	"io"
	"time"
)

func init() {
	gob.RegisterName("RequestBody.CopyNodeBo", CopyNodeBo{})
	gob.RegisterName("RequestBody.MoveNodeBo", MoveNodeBo{})
	gob.RegisterName("RequestBody.SrcAndDstBo", SrcAndDstBo{})
	gob.RegisterName("RequestBody.PwdUpdateBo", PwdUpdateBo{})
	gob.RegisterName("RequestBody.CreateUserBo", CreateUserBo{})
	gob.RegisterName("RequestBody.WriteTokenBo", WriteTokenBo{})
	gob.RegisterName("RequestBody.LimitQueryBo", LimitQueryBo{})
	gob.RegisterName("RequestBody.UserListDto", UserListDto{})
	gob.RegisterName("RequestBody.DirNameListDto", DirNameListDto{})
	gob.RegisterName("RequestBody.DirNodeListDto", DirNodeListDto{})
}

// StreamTokenType token类型
type StreamTokenType int

// IOpenApi 开放api
type IOpenApi interface {
	ConnTest(user User) error

	IsDir(src string) (bool, error)
	IsFile(src string) (bool, error)
	IsExist(src string) (bool, error)
	GetFileSize(src string) (int64, error)

	GetNode(src string) (*TNode, error)
	GetNodes(src []string) ([]TNode, error)
	GetDirNameList(src string, limit, offset int) (*DirNameListDto, error)
	GetDirNodeList(src string, limit, offset int) (*DirNodeListDto, error)

	DoMkDir(path string) (string, error)
	DoDelete(src string) error
	DoRename(src string, dst string) error
	DoCopy(src, dst string, override bool) (string, error)
	DoMove(src, dst string, override bool) error

	DoQueryToken(token string) (*StreamToken, error)
	DoAskReadToken(src string) (*StreamToken, error)
	DoAskWriteToken(src string) (*StreamToken, error)
	DoRefreshToken(token string) (*StreamToken, error)
	DoSubmitWriteToken(token string, override bool) (*TNode, error)
	SetDataNodeDNS(nodes map[string]string)
	GetReadStreamURL(nodeNo, token, endpoint string) (string, error)
	GetWriteStreamURL(nodeNo, token, endpoint string) (string, error)
	DoWriteToken(nodeNo, token, endpoint string, pieceNumber int, sha256 string, reader io.Reader) error
	DoReadToken(nodeNo, token, endpoint string, offset int64) (io.ReadCloser, error)
}

// IServerMG 管理类接口
type IServerMG interface {
	ConnTest(user User) error
	DoCreateUser(user CreateUserBo) (bool, error)
	DoUpdatePWD(user, pwd string) (bool, error)
	DoDeleteUser(userID string) (bool, error)
	DoListAllUsers() (*UserListDto, error)
}

// TNode 元数据内容(filenames)
type TNode struct {
	Id    string
	Pid   string
	Addr  string
	Flag  int
	Name  string
	Size  int64
	Ctime int64
	Mtime int64
	Props string
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
		st.Type = ua.Type
		return nil
	}
	return fmt.Errorf("can't support clone %T ", val)
}

// SrcAndDstBo src 和 dst
type SrcAndDstBo struct {
	SRC string
	DST string
}

// CopyNodeBo 拷贝
type CopyNodeBo struct {
	SRC      string
	DST      string
	Override bool
}

// MoveNodeBo 移动
type MoveNodeBo struct {
	SRC      string
	DST      string
	Override bool
}

// WriteTokenBo 文件上传递交token
type WriteTokenBo struct {
	Token    string
	Override bool
}

// LimitQueryBo 分页查询
type LimitQueryBo struct {
	Query  string
	Limit  int
	Offset int
}

// CreateUserBo 用户表存储的结构
type CreateUserBo struct {
	UserType int
	UserID   string
	UserName string
	UserPWD  string
}

// PwdUpdateBo 更改用户密码
type PwdUpdateBo struct {
	User string
	Pwd  string
}

// UserListDto 用户信息列表
type UserListDto struct {
	Datas []UserInfo
	Total int
}

// UserInfo 用户表存储的结构
type UserInfo struct {
	UserType int
	UserID   string
	UserName string
	CtTime   time.Time
}

// DirNameListDto 文件夹目录列表
type DirNameListDto struct {
	Datas []string
	Total int
}

// DirNodeListDto 文件夹目录列表
type DirNodeListDto struct {
	Datas []TNode
	Total int
}
