// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 名称节点客户端

package namenode

import (
	"datanode/biz/bizutils"
	"datanode/biz/rpcservice"
	"datanode/ifilestorage"
	"pakku/utils/logs"
)

// NewRPC4NameNode NewRPC4NameNode
func NewRPC4NameNode(conn *bizutils.Conn4RPC) ifilestorage.RPC4NameNode {
	return &RPC4NameNode{conn}
}

// RPC4NameNode 命名空间节点
type RPC4NameNode struct {
	conn *bizutils.Conn4RPC
}

func (s *RPC4NameNode) IsDir(src string) bool {
	var res bool
	if err := s.conn.DoRPC("NameNode.IsDir", src, &res); nil != err {
		logs.Panicln(err)
		return false
	}
	return res
}
func (s *RPC4NameNode) IsFile(src string) bool {
	var res bool
	if err := s.conn.DoRPC("NameNode.IsFile", src, &res); nil != err {
		logs.Panicln(err)
		return false
	}
	return res
}
func (s *RPC4NameNode) IsExist(src string) bool {
	var res bool
	if err := s.conn.DoRPC("NameNode.IsExist", src, &res); nil != err {
		logs.Panicln(err)
		return false
	}
	return res
}
func (s *RPC4NameNode) GetFileSize(src string) int64 {
	var res int64
	if err := s.conn.DoRPC("NameNode.GetFileSize", src, &res); nil != err {
		logs.Panicln(err)
		return -1
	}
	return res
}
func (s *RPC4NameNode) DoQueryToken(token string) (*ifilestorage.StreamToken, error) {
	var res *ifilestorage.StreamToken
	if err := s.conn.DoRPC("NameNode.DoQueryToken", token, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *RPC4NameNode) DoAskReadToken(src string) (*ifilestorage.StreamToken, error) {
	var res *ifilestorage.StreamToken
	if err := s.conn.DoRPC("NameNode.DoAskReadToken", src, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *RPC4NameNode) DoAskWriteToken(src string) (*ifilestorage.StreamToken, error) {
	var res *ifilestorage.StreamToken
	if err := s.conn.DoRPC("NameNode.DoAskWriteToken", src, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *RPC4NameNode) DoRefreshToken(token string) (*ifilestorage.StreamToken, error) {
	var res *ifilestorage.StreamToken
	if err := s.conn.DoRPC("NameNode.DoRefreshToken", token, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *RPC4NameNode) DoQueryDeletedDataAddr(nodeno string, limit int) ([]ifilestorage.DeletedAddr, error) {
	var res []ifilestorage.DeletedAddr
	if err := s.conn.DoRPC("NameNode.DoQueryDeletedDataAddr", rpcservice.LimitQueryBo{Query: nodeno, Limit: limit}, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *RPC4NameNode) DoConfirmDeletedDataAddr(nodenos []string) error {
	var res bool
	if err := s.conn.DoRPC("NameNode.DoConfirmDeletedDataAddr", nodenos, &res); nil != err {
		return err
	}
	return nil
}
