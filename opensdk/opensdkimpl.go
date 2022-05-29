// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 开放api接口

package opensdk

import (
	"io"
)

// NewOpenApi NewOpenApi
func NewOpenApi(addr string, user User) IOpenApi {
	return &OpenSdkImpl{NewConn4RPC(addr, user), NewDataStream()}
}

// OpenSdkImpl 开放api接口
type OpenSdkImpl struct {
	conn *Conn4RPC
	dsm  *DataStream
}

// GetRPCConn GetRPCConn
func (s *OpenSdkImpl) GetRPCConn() *Conn4RPC {
	return s.conn
}

func (s *OpenSdkImpl) IsDir(src string) (bool, error) {
	var res bool
	if err := s.conn.DoRPC("NameNode.IsDir", src, &res); nil != err {
		return false, err
	}
	return res, nil
}
func (s *OpenSdkImpl) IsFile(src string) (bool, error) {
	var res bool
	if err := s.conn.DoRPC("NameNode.IsFile", src, &res); nil != err {
		return false, err
	}
	return res, nil
}
func (s *OpenSdkImpl) IsExist(src string) (bool, error) {
	var res bool
	if err := s.conn.DoRPC("NameNode.IsExist", src, &res); nil != err {
		return false, err
	}
	return res, nil
}
func (s *OpenSdkImpl) GetFileSize(src string) (int64, error) {
	var res int64
	if err := s.conn.DoRPC("NameNode.GetFileSize", src, &res); nil != err {
		return 0, err
	}
	return res, nil
}
func (s *OpenSdkImpl) GetNode(src string) (*TNode, error) {
	var res *TNode
	if err := s.conn.DoRPC("NameNode.GetNode", src, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *OpenSdkImpl) GetNodes(src []string) ([]TNode, error) {
	var res []TNode
	if err := s.conn.DoRPC("NameNode.GetNodes", src, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *OpenSdkImpl) GetDirNameList(src string, limit, offset int) (*DirNameListDto, error) {
	var res DirNameListDto
	if err := s.conn.DoRPC("NameNode.GetDirNameList", LimitQueryBo{Query: src, Limit: limit, Offset: offset}, &res); nil != err {
		return nil, err
	}
	return &res, nil
}
func (s *OpenSdkImpl) GetDirNodeList(src string, limit, offset int) (*DirNodeListDto, error) {
	var res DirNodeListDto
	if err := s.conn.DoRPC("NameNode.GetDirNodeList", LimitQueryBo{Query: src, Limit: limit, Offset: offset}, &res); nil != err {
		return nil, err
	}
	return &res, nil
}

func (s *OpenSdkImpl) DoMkDir(path string) (string, error) {
	var res string
	if err := s.conn.DoRPC("NameNode.DoMkDir", path, &res); nil != err {
		return "", err
	}
	return res, nil
}
func (s *OpenSdkImpl) DoDelete(src string) error {
	var res bool
	if err := s.conn.DoRPC("NameNode.DoDelete", src, &res); nil != err {
		return err
	}
	return nil
}
func (s *OpenSdkImpl) DoRename(src string, dst string) error {
	var res bool
	sd := SrcAndDstBo{
		SRC: src,
		DST: dst,
	}
	if err := s.conn.DoRPC("NameNode.DoRename", sd, &res); nil != err {
		return err
	}
	return nil
}
func (s *OpenSdkImpl) DoCopy(src, dst string, override bool) (string, error) {
	var res string
	sd := CopyNodeBo{
		SRC:      src,
		DST:      dst,
		Override: override,
	}
	if err := s.conn.DoRPC("NameNode.DoCopy", sd, &res); nil != err {
		return "", err
	}
	return res, nil
}
func (s *OpenSdkImpl) DoMove(src, dst string, override bool) error {
	var res bool
	sd := MoveNodeBo{
		SRC:      src,
		DST:      dst,
		Override: override,
	}
	if err := s.conn.DoRPC("NameNode.DoMove", sd, &res); nil != err {
		return err
	}
	return nil
}

func (s *OpenSdkImpl) DoQueryToken(token string) (*StreamToken, error) {
	var res *StreamToken
	if err := s.conn.DoRPC("NameNode.DoQueryToken", token, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *OpenSdkImpl) DoAskReadToken(src string) (*StreamToken, error) {
	var res *StreamToken
	if err := s.conn.DoRPC("NameNode.DoAskReadToken", src, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *OpenSdkImpl) DoAskWriteToken(src string) (*StreamToken, error) {
	var res *StreamToken
	if err := s.conn.DoRPC("NameNode.DoAskWriteToken", src, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *OpenSdkImpl) DoRefreshToken(token string) (*StreamToken, error) {
	var res *StreamToken
	if err := s.conn.DoRPC("NameNode.DoRefreshToken", token, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *OpenSdkImpl) DoSubmitWriteToken(token string, override bool) (node *TNode, err error) {
	var res *TNode
	wt := WriteTokenBo{Token: token, Override: override}
	if err := s.conn.DoRPC("NameNode.DoSubmitWriteToken", wt, &res); nil != err {
		return nil, err
	}
	return res, nil
}
func (s *OpenSdkImpl) SetDataNodeDNS(nodes map[string]string) {
	s.dsm.SetNodeStreamAddr(nodes)
}
func (s *OpenSdkImpl) GetReadStreamURL(nodeNo, token, endpoint string) (string, error) {
	return s.dsm.GetReadStreamURL(nodeNo, token, endpoint)
}
func (s *OpenSdkImpl) GetWriteStreamURL(nodeNo, token, endpoint string) (string, error) {
	return s.dsm.GetWriteStreamURL(nodeNo, token, endpoint)
}
func (s *OpenSdkImpl) DoWriteToken(nodeNo, token, endpoint string, pieceNumber int, sha256 string, reader io.Reader) (err error) {
	return s.dsm.DoWriteToken(nodeNo, token, endpoint, pieceNumber, sha256, reader)
}
func (s *OpenSdkImpl) DoReadToken(nodeNo, token, endpoint string, offset int64) (io.ReadCloser, error) {
	return s.dsm.DoReadToken(nodeNo, token, endpoint, offset)
}
