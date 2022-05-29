// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// namenode RPC服务

package rpcservice

import (
	"namenode/ifilestorage"
	"pakku/utils/fileutil"
)

// NameNode 命名空间节点
type NameNode struct {
	ua ifilestorage.UserAuth4Rpc `@autowired:"User4RPC"`
	p  ifilestorage.NameNode     `@autowired:"NameNode"`
}

func (n *NameNode) IsDir(req *RequestBody4RPC, res *bool) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	*res = n.p.IsDir(req.RequestBody.(string))
	return nil
}

func (n *NameNode) IsFile(req *RequestBody4RPC, res *bool) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	*res = n.p.IsFile(req.RequestBody.(string))
	return nil
}

func (n *NameNode) IsExist(req *RequestBody4RPC, res *bool) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	*res = n.p.IsExist(req.RequestBody.(string))
	return nil
}

func (n *NameNode) GetFileSize(req *RequestBody4RPC, res *int64) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	if node := n.p.GetNode(req.RequestBody.(string)); nil == node {
		return fileutil.PathNotExist("GetFileSize", req.RequestBody.(string))
	} else {
		*res = node.Size
	}
	return nil
}

func (n *NameNode) GetNode(req *RequestBody4RPC, res *ifilestorage.TNode) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	if node := n.p.GetNode(req.RequestBody.(string)); nil != node {
		*res = *node
	} else {
		return fileutil.PathNotExist("GetNode", req.RequestBody.(string))
	}
	return nil
}

func (n *NameNode) GetNodes(req *RequestBody4RPC, res *[]ifilestorage.TNode) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	paths := req.RequestBody.([]string)
	temp := make([]ifilestorage.TNode, len(paths))
	for i := 0; i < len(paths); i++ {
		if node := n.p.GetNode(paths[i]); nil == node {
			return fileutil.PathNotExist("GetNodes", paths[i])
		} else {
			temp[i] = *node
		}
	}
	*res = temp
	return nil
}

func (n *NameNode) GetDirNameList(req *RequestBody4RPC, res *DirNameListDto) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	query := req.RequestBody.(LimitQueryBo)
	res.Datas, res.Total = n.p.GetDirNameList(query.Query, query.Limit, query.Offset)
	return nil
}

func (n *NameNode) GetDirNodeList(req *RequestBody4RPC, res *DirNodeListDto) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	query := req.RequestBody.(LimitQueryBo)
	res.Datas, res.Total = n.p.GetDirNodeList(query.Query, query.Limit, query.Offset)
	return nil
}

func (n *NameNode) DoMkDir(req *RequestBody4RPC, res *string) (err error) {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi); nil != err {
		return err
	}
	*res, err = n.p.DoMkDir(req.RequestBody.(string))
	return err
}

func (n *NameNode) DoDelete(req *RequestBody4RPC, res *bool) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi); nil != err {
		return err
	}
	err := n.p.DoDelete(req.RequestBody.(string))
	*res = nil == err
	return err
}

func (n *NameNode) DoRename(req *RequestBody4RPC, res *bool) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi); nil != err {
		return err
	}
	co := req.RequestBody.(SrcAndDstBo)
	err := n.p.DoRename(co.SRC, co.DST)
	*res = nil == err
	return err
}

func (n *NameNode) DoCopy(req *RequestBody4RPC, res *string) (err error) {
	if err = req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi); nil == err {
		co := req.RequestBody.(CopyNodeBo)
		*res, err = n.p.DoCopy(co.SRC, co.DST, co.Override)
	}
	return err
}

func (n *NameNode) DoMove(req *RequestBody4RPC, res *bool) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi); nil != err {
		return err
	}
	co := req.RequestBody.(MoveNodeBo)
	err := n.p.DoMove(co.SRC, co.DST, co.Override)
	*res = nil == err
	return err
}

func (n *NameNode) DoQueryToken(req *RequestBody4RPC, res *ifilestorage.StreamToken) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	if temp, err := n.p.DoQueryToken(req.RequestBody.(string)); nil == err {
		*res = *temp
	} else {
		return err
	}
	return nil
}

func (n *NameNode) DoAskReadToken(req *RequestBody4RPC, res *ifilestorage.StreamToken) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	if temp, err := n.p.DoAskReadToken(req.RequestBody.(string)); nil == err {
		*res = *temp
	} else {
		return err
	}
	return nil
}

func (n *NameNode) DoAskWriteToken(req *RequestBody4RPC, res *ifilestorage.StreamToken) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	if temp, err := n.p.DoAskWriteToken(req.RequestBody.(string)); nil == err {
		*res = *temp
	} else {
		return err
	}
	return nil
}

func (n *NameNode) DoRefreshToken(req *RequestBody4RPC, res *ifilestorage.StreamToken) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	if temp, err := n.p.DoRefreshToken(req.RequestBody.(string)); nil == err {
		*res = *temp
	} else {
		return err
	}
	return nil
}

func (n *NameNode) DoSubmitWriteToken(req *RequestBody4RPC, res *ifilestorage.TNode) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_OpenApi+1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	wt := req.RequestBody.(WriteTokenBo)
	if temp, err := n.p.DoSubmitWriteToken(wt.Token, wt.Override); nil == err {
		*res = *temp
	} else {
		return err
	}
	return nil
}

func (n *NameNode) DoQueryDeletedDataAddr(req *RequestBody4RPC, res *[]ifilestorage.DeletedAddr) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	}
	co := req.RequestBody.(LimitQueryBo)
	if temp, err := n.p.DoQueryDeletedDataAddr(co.Query, co.Limit); nil == err {
		*res = temp
	} else {
		return err
	}
	return nil
}

func (n *NameNode) DoConfirmDeletedDataAddr(req *RequestBody4RPC, res *bool) error {
	if err := req.doValidation(n.ua, req.AccessKey, "", 1<<ifilestorage.UserRule_DataNode); nil != err {
		return err
	} else {
		if err = n.p.DoConfirmDeletedDataAddr(req.RequestBody.([]string)); nil == err {
			*res = true
		}
		return err
	}
}
