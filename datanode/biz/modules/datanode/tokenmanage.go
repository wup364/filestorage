// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package datanode

import (
	"datanode/ifilestorage"
	"pakku/ipakku"
	"pakku/utils/logs"
	"time"
)

const (
	CacheLib_StreamToken         = "datanode.streamtoken"
	CacheLib_TokenExpMilliSecond = 1000 * 60 * 5  // 每5分钟重新拉取一次, namenode是15分钟
	TokenExpMilliSecond_NameNode = 1000 * 60 * 15 // 15分钟namenode的token就过期了
)

// NewTokenManage NewTokenManage
func NewTokenManage(nn ifilestorage.RPC4NameNode, ch ipakku.AppCache) *TokenManage {
	return &TokenManage{
		pn: nn,
		ch: ch,
	}
}

// TokenManage TokenManage
type TokenManage struct {
	pn ifilestorage.RPC4NameNode
	ch ipakku.AppCache
}

// getTokenFromRPC 获取token信息
func (t *TokenManage) getTokenFromRPC(token string) *ifilestorage.StreamToken {
	if st, err := t.pn.DoQueryToken(token); nil == err {
		if nil == st {
			return nil
		}
		if st.MTime+TokenExpMilliSecond_NameNode-time.Now().UnixMilli() <= CacheLib_TokenExpMilliSecond {
			if st, err = t.pn.DoRefreshToken(token); nil == err {
				logs.Infof("refresh token, token=%s, ctime=%d, mtime=%d \r\n", token, st.CTime, st.MTime)
			}
		}
		if nil == err {
			err = t.setTokenCache(token, st)
		}
		if nil == err {
			return st
		} else {
			logs.Errorln(err)
		}
	} else {
		logs.Errorln(err)
	}
	return nil
}

// getTokenFromCache 获取token信息
func (t *TokenManage) getTokenFromCache(token string) *ifilestorage.StreamToken {
	var val ifilestorage.StreamToken
	if err := t.ch.Get(CacheLib_StreamToken, token, &val); nil == err {
		return &val
	}
	return nil
}

// GetToken 将token保存在本地
func (t *TokenManage) setTokenCache(token string, tbody *ifilestorage.StreamToken) error {
	return t.ch.Set(CacheLib_StreamToken, token, tbody)
}

func (t *TokenManage) DoRefreshToken(token string) (*ifilestorage.StreamToken, error) {
	if st, err := t.pn.DoRefreshToken(token); nil == err {
		if err := t.setTokenCache(token, st); nil != err {
			logs.Errorln(err)
		}
		return st, nil
	} else {
		return nil, err
	}
}

// GetToken 获取token信息
func (t *TokenManage) GetToken(token string) *ifilestorage.StreamToken {
	if len(token) > 0 {
		var res *ifilestorage.StreamToken
		if res = t.getTokenFromCache(token); nil == res {
			res = t.getTokenFromRPC(token)
		}
		return res
	}
	return nil
}
