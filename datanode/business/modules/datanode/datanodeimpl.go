// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 数据存储节点服务

package datanode

import (
	"database/sql"
	"datanode/ifilestorage"
	"errors"
	"io"

	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
)

// NewDataNodeImpl datanode 功能本机实现
func NewDataNodeImpl(ds *DataNodeStory, df *DataNodeFiles, tm *TokenManage) *DataNodeImpl {
	return &DataNodeImpl{
		dns: ds,
		df:  df,
		tm:  tm,
	}
}

// DataNodeImpl 数据存储节点
type DataNodeImpl struct {
	dns *DataNodeStory
	dhs *DataHashStory
	df  *DataNodeFiles
	tm  *TokenManage
}

// GetDataFileNode 根据归档ID获取文件信息
func (dn *DataNodeImpl) GetDataFileNode(id string) (node *ifilestorage.DNode) {
	var err error
	if node, err = dn.dns.GetNode(dn.dns.GetDB(), id); nil == err {
		return node
	}
	return nil
}

// DoQueryToken 查询token信息
func (dn *DataNodeImpl) DoQueryToken(token string) (*ifilestorage.StreamToken, error) {

	if val := dn.tm.GetToken(token); nil == val {
		return nil, ErrInvalidToken
	} else {
		return val, nil
	}
}

// DoRefreshToken 刷新token信息
func (dn *DataNodeImpl) DoRefreshToken(token string) (*ifilestorage.StreamToken, error) {
	return dn.tm.DoRefreshToken(token)
}

// DoWriteToken 通过token计算写入位置, 并记录写入信息
func (dn *DataNodeImpl) DoWriteToken(token string, pieceNumber int, sha256 string, reader io.Reader) (err error) {
	// 验证token是否有效
	if streamToken := dn.tm.GetToken(token); nil != streamToken {
		return dn.df.DoWriteStream(&ifilestorage.StreamWriteOpts{
			Token:       token,
			Sha256:      sha256,
			PieceNumber: pieceNumber,
		}, reader)
	} else {
		err = ErrInvalidToken
	}
	return err
}

// DoReadToken 通过token获取文件信息, 并读取
func (dn *DataNodeImpl) DoReadToken(token string, offset int64) (r io.Reader, err error) {
	if streamToken := dn.tm.GetToken(token); nil != streamToken {
		if node := dn.GetDataFileNode(streamToken.FileID); nil != node {
			r, err = dn.df.DoReadStream(&ifilestorage.StreamReadOpts{
				Token:  token,
				Offset: offset,
				Pieces: node.Pieces,
			}, offset)
		}
	} else {
		err = ErrInvalidToken
	}
	return r, err
}

// DoSubmitWriteToken 提交token, 表示写入完成
func (dn *DataNodeImpl) DoSubmitWriteToken(token string) (node *ifilestorage.DNode, err error) {
	if streamToken := dn.tm.GetToken(token); nil != streamToken {
		defer dn.df.DoDestroyToken(token)
		if dn.GetDataFileNode(streamToken.FileID) != nil {
			return nil, errors.New("token is submit, token: " + token)
		}
		if node, err = dn.df.DoArchiveToken(token); nil == err {
			node.Id = streamToken.FileID
			for {
				if len(node.Id) > 0 && dn.GetDataFileNode(node.Id) == nil {
					break
				}
				node.Id = strutil.GetUUID()
			}
			var conn4dns *sql.Tx
			if conn4dns, err = dn.dns.GetSqlTx(); nil == err {
				if err = dn.dns.InsertNode(conn4dns, *node); nil == err {
					var conn4dhs *sql.Tx
					if conn4dhs, err = dn.dhs.GetSqlTx(); nil == err {
						for i := 0; i < len(node.Pieces); i++ {
							if err = dn.dhs.InsertHash(conn4dhs, ifilestorage.HNode{Id: strutil.GetUUID(), Status: status_hashdata_enable, FId: node.Id, Hash: node.Pieces[i]}); nil != err {
								conn4dhs.Rollback()
								break
							}
						}
					}
					if nil == err {
						if err = conn4dns.Commit(); nil == err {
							if err = conn4dhs.Commit(); nil != err {
								var err1 error
								if conn4dns, err1 = dn.dns.GetSqlTx(); nil == err1 {
									if err1 = dn.dns.DeleteByID(conn4dns, node.Id); nil == err1 {
										err1 = conn4dns.Commit()
									} else {
										conn4dns.Rollback()
									}
								}
								if nil != err1 {
									logs.Errorln("datahash table commit failed and rollback datanode failed, node.id=" + node.Id)
									logs.Panicln(err)
								}
							}
						} else {
							conn4dhs.Rollback()
						}
					} else {
						conn4dns.Rollback()
					}
				} else {
					conn4dns.Rollback()
				}
			}
		}
	}
	if nil != err {
		node = nil
	}
	return node, err
}
