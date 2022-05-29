// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 数据节点RPC客户端

package datanodes

import (
	"namenode/biz/bizutils"
	"namenode/ifilestorage"
)

// NewRPC4DataNode NewRPC4DataNode
func NewRPC4DataNode(conn *bizutils.Conn4RPC, props *ifilestorage.Prop4DataNodeConn) ifilestorage.RPC4DataNode {
	return &RPC4DataNode{conn: conn, props: props}
}

// RPC4DataNode 文件数据块节点
type RPC4DataNode struct {
	props *ifilestorage.Prop4DataNodeConn
	conn  *bizutils.Conn4RPC
}

func (s *RPC4DataNode) GetNodeNo() string {
	return s.conn.GetUser()
}

func (s *RPC4DataNode) GetStreamEndpoint() string {
	if nil == s.props {
		return ""
	}
	return s.props.HTTPAddress
}

func (s *RPC4DataNode) DoSubmitWriteToken(token string) (node *ifilestorage.DNode, err error) {
	var res *ifilestorage.DNode
	if err := s.conn.DoRPC("DataNode.DoSubmitWriteToken", token, &res); nil != err {
		return nil, err
	}
	return res, nil
}
