// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package ifilestorage

import (
	"io"
)

// DNode 数据存储节点文件信息
type DNode struct {
	Id     string
	Size   int64
	Hash   string
	Pieces []string
}

// HNode Hash使用索引
type HNode struct {
	Id     string
	Status int
	FId    string
	Hash   string
}

// StreamWriteOpts StreamWriteOpts
type StreamWriteOpts struct {
	Token       string
	Sha256      string
	PieceNumber int
}

// StreamReadOpts StreamReadOpts
type StreamReadOpts struct {
	Token  string
	Offset int64
	Pieces []string
}

// DataNode 文件数据块节点
type DataNode interface {
	GetNodeNo() string
	GetDataFileNode(id string) (node *DNode)
	DoQueryToken(token string) (*StreamToken, error)
	DoRefreshToken(token string) (*StreamToken, error)
	DoSubmitWriteToken(token string) (node *DNode, err error)
	DoReadToken(token string, offset int64) (r io.Reader, err error)
	DoWriteToken(token string, pieceNumber int, sha256 string, reader io.Reader) (err error)
}

// DataNodePrivateRPC DataNodePrivateRPC
type DataNodePrivateRPC interface {
	GetNodeNo() string
	DoSubmitWriteToken(token string) (*DNode, error)
}
