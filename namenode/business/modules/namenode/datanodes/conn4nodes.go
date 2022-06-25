// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// datanode 连接管理
package datanodes

import (
	"errors"
	"math/rand"
	"namenode/business/bizutils"
	"namenode/business/rpcservice"
	"namenode/ifilestorage"
	"time"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/logs"
	"github.com/wup364/pakku/utils/strutil"
	"github.com/wup364/pakku/utils/utypes"
)

var ErrNoAliveDNode = errors.New("not have alive datanode")

// NewConn4DataNodes NewConn4DataNodes
func NewConn4DataNodes(syncEvent ipakku.AppSyncEvent) ifilestorage.Conn4DataNodes {
	conn := &Conn4DataNodes{
		conns: utypes.NewSafeMap(),
		ev:    syncEvent,
	}
	conn.ListenLoginSucceedEvent()
	return conn
}

// Conn4DataNodes 记录连接好的datanode
type Conn4DataNodes struct {
	conns *utypes.SafeMap
	ev    ipakku.AppSyncEvent
}

// ListenLoginSucceedEvent 监听登录成功事件
func (s *Conn4DataNodes) ListenLoginSucceedEvent() {
	s.ev.ConsumerSyncEvent("authApi", "dologin.succeed", func(v interface{}) error {
		if user, ok := v.(rpcservice.User); ok && len(user.PropJSON) > 0 {
			props := &ifilestorage.Prop4DataNodeConn{}
			if err := props.ParseJSON(user.PropJSON); nil != err {
				logs.Errorln(err)
				return nil
			}
			s.RegisterConn(ifilestorage.Conn4RPC{
				RPCAddr: props.RPCAddress,
				User: &ifilestorage.User{
					User:   props.NodeNo,
					Passwd: strutil.GetSHA256(props.NodeNo + "@" + props.NodePid),
				},
			}, *props)
			logs.Infof("datanode login succeed, user: %s, instanceId: %s, Props: %s", user.User, props.NodePid, user.PropJSON)
		}
		return nil
	})
}

// RegisterConn 注册datanode连接
func (s *Conn4DataNodes) RegisterConn(conn ifilestorage.Conn4RPC, props ifilestorage.Prop4DataNodeConn) {
	conn4rpc := bizutils.NewConn4RPC(conn.RPCAddr, bizutils.User{
		User:   conn.User.User,
		Passwd: conn.User.Passwd,
	})
	s.conns.Put(conn.User.User, NewRPC4DataNode(conn4rpc, &props))
}

// UnRegisterConn 注销datanode连接
func (s *Conn4DataNodes) UnRegisterConn(user string) {
	s.conns.Delete(user)
}

// GetDNode 根据编号获取在线的datanode
func (s *Conn4DataNodes) GetDNode(no string) (ifilestorage.RPC4DataNode, error) {
	if val, ok := s.conns.Get(no); ok {
		if res, ok := val.(ifilestorage.RPC4DataNode); ok {
			return res, nil
		}
	}
	return nil, errors.New(no + " is not alive")
}

// GetRandomDNode 随机获取一个在线的datanode
func (s *Conn4DataNodes) GetRandomDNode() (ifilestorage.RPC4DataNode, error) {
	nos := s.conns.Keys()
	if len(nos) == 0 {
		return nil, ErrNoAliveDNode
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	if val, ok := s.conns.Get(nos[r.Intn(len(nos))]); ok {
		return val.(ifilestorage.RPC4DataNode), nil
	} else {
		return s.GetRandomDNode()
	}
}
