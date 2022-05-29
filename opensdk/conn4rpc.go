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
	"errors"
	"net/rpc"
	"opensdk/utils"
)

var ErrorAuthentication = errors.New("authentication failed")

// NewConn4RPC RPC客户端
func NewConn4RPC(addr string, user User) *Conn4RPC {
	return &Conn4RPC{
		addr:    addr,
		user:    &user,
		session: &Session{},
		clients: utils.NewSafeMap(),
	}
}

// Conn4RPC Conn4RPC
type Conn4RPC struct {
	addr    string
	user    *User
	session *Session
	clients *utils.SafeMap
}

// Session 授权信息
type Session struct {
	AccessKey string
	SecretKey string
}

// User 登录凭据
type User struct {
	User     string
	Passwd   string
	PropJSON string
}

// RequestBody RPC请求体
type RequestBody struct {
	AccessKey   string
	Signature   string
	RequestBody interface{}
}

// GetUser 获取用户ID
func (c *Conn4RPC) GetUser() string {
	if nil != c.user {
		return c.user.User
	}
	return ""
}

// GetAccessKey 获取accessKey, 为空则自动登录一下
func (c *Conn4RPC) GetAccessKey() (string, error) {
	if nil == c.session || len(c.session.AccessKey) == 0 {
		if err := c.DoLogin(); nil != err {
			return "", err
		}
	}
	return c.session.AccessKey, nil
}

// GetSecretKey 获取 secretKey
func (c *Conn4RPC) GetSecretKey() (string, error) {
	if nil == c.session || len(c.session.SecretKey) == 0 {
		if err := c.DoLogin(); nil != err {
			return "", err
		}
	}
	return c.session.SecretKey, nil
}

// DoLogin 登录一下, 获取 Data={"AccessKey":"e194244b-1683-cb4c-98a9-a79e43fe1576","SecretKey":"e194244b-1683-cb4c-98a9-a05f6ef420d2"}
func (c *Conn4RPC) DoLogin() error {
	if nil != c.session {
		c.session.AccessKey = ""
		c.session.SecretKey = ""
	}
	{
		if client, err := c.getClient(); nil != err {
			return err
		} else {
			defer client.Close()
			reply := &Session{}
			if err = client.Call("AuthApi.DoLogin", c.user, reply); nil != err {
				return err
			}
			c.session = reply
		}
	}
	return nil
}

// DoRPC RPC
func (c *Conn4RPC) DoRPC(method string, args interface{}, reply interface{}) (err error) {
	if err = c.doRPC(method, args, reply); nil != err {
		// 可能是重启了, 重连一下
		if err.Error() == ErrorAuthentication.Error() {
			if err := c.DoLogin(); nil == err {
				return c.doRPC(method, args, reply)
			}
		}
	}
	return err
}

// doRPC RPC
func (c *Conn4RPC) doRPC(method string, args interface{}, reply interface{}) (err error) {
	var ak string
	if ak, err = c.GetAccessKey(); nil == err {
		var sg string
		if sg, err = c.doSignature(""); nil == err {
			var client *RPCClient
			if client, err = c.getClient(); nil == err {
				defer client.Close()
				err = client.Call(method, RequestBody{RequestBody: args, AccessKey: ak, Signature: sg}, reply)
			}
		}
	}
	return err
}

// doSignature 请求签名一下
func (c *Conn4RPC) doSignature(playload string) (string, error) {
	if nil == c.session || len(c.session.SecretKey) == 0 {
		if err := c.DoLogin(); nil != err {
			return "", err
		}
	}
	return utils.GetStringSHA256(c.session.AccessKey + "/" + c.session.SecretKey + "/" + playload), nil
}

// getClient 获取一个空闲的客户端
func (c *Conn4RPC) getClient() (*RPCClient, error) {
	if nil == c.clients {
		return c.createClient(true)
	}
	if client, ok := c.clients.CutR(); ok {
		return client.(*RPCClient), nil
	} else {
		return c.createClient(true)
	}
}

// createClient 创建一个空闲的客户端
func (c *Conn4RPC) createClient(save bool) (*RPCClient, error) {
	if client, err := rpc.DialHTTP("tcp", c.addr); nil != err {
		return nil, err
	} else {
		for i := 0; i < 65535; i++ {
			if id := utils.GetRandom(16); c.clients.ContainsKey(id) {
				continue
			} else {
				rpcclient := &RPCClient{
					ClientID: id,
					Client:   client,
					conn4rpc: c,
				}
				if !save {
					return rpcclient, nil
				} else {
					if c.clients.PutX(id, rpcclient) != nil {
						continue
					}
					return rpcclient, nil
				}
			}
		}
		return nil, errors.New("failed to create client")
	}
}

// RPCClient RPCClient
type RPCClient struct {
	conn4rpc *Conn4RPC
	ClientID string
	shutdown bool
	*rpc.Client
}

// Call override Call
func (client *RPCClient) Call(serviceMethod string, args interface{}, reply interface{}) error {
	if err := client.Client.Call(serviceMethod, args, reply); nil != err {
		// 连接已经关闭了, 重试一下(重新创建新的连接, 旧的连接可能都坏了..)
		if err == rpc.ErrShutdown {
			client.shutdown = true
			if c, err := client.conn4rpc.createClient(false); nil != err {
				return err
			} else {
				defer c.Close()
				return c.Client.Call(serviceMethod, args, reply)
			}
		}
		return err
	}
	return nil
}

// Close override Close
func (client *RPCClient) Close() error {
	if !client.shutdown {
		client.conn4rpc.clients.PutX(client.ClientID, client)
	}
	return nil
}
