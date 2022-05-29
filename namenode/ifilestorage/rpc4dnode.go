// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package ifilestorage

import "encoding/json"

// DNode 数据存储节点文件信息
type DNode struct {
	Id     string
	Size   int64
	Hash   string
	Pieces []string
}

// RPC4DataNode RPC4DataNode
type RPC4DataNode interface {
	GetNodeNo() string
	GetStreamEndpoint() string
	DoSubmitWriteToken(token string) (*DNode, error)
}

// Prop4DataNodeConn Prop4DataNodeConn
type Prop4DataNodeConn struct {
	NodeNo      string
	NodePid     string
	RPCAddress  string
	HTTPAddress string
}

// ToJSON ToJSON
func (prop Prop4DataNodeConn) ToJSON() string {
	if bt, err := json.Marshal(prop); nil == err {
		return string(bt)
	}
	return ""
}

// ParseJSON ParseJSON
func (prop *Prop4DataNodeConn) ParseJSON(jsonStr string) error {
	return json.Unmarshal([]byte(jsonStr), &prop)
}

// Conn4RPC Conn4RPC
type Conn4RPC struct {
	RPCAddr string // rpc调用地址
	User    *User
}

// User 授权信息
type User struct {
	User   string
	Passwd string
}

// Conn4DataNodes datanode conns
type Conn4DataNodes interface {
	UnRegisterConn(user string)
	GetRandomDNode() (RPC4DataNode, error)
	GetDNode(no string) (RPC4DataNode, error)
	RegisterConn(conn Conn4RPC, props Prop4DataNodeConn)
}
