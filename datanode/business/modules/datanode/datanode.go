// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 数据存储节点

package datanode

import (
	"datanode/business/bizutils"
	"datanode/business/modules/datanode/namenode"
	"datanode/ifilestorage"
	"time"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/logs"
)

// DataNode 数据存储节点
type DataNode struct {
	DataNodeImpl
	nodeno string
	dhc    *HashDataCtrl
	nconn  *bizutils.Conn4RPC
	nrpc   ifilestorage.RPC4NameNode
	ch     ipakku.AppCache        `@autowired:"AppCache"`
	cf     ipakku.AppConfig       `@autowired:"AppConfig"`
	fd     ifilestorage.FileDatas `@autowired:"FileDatas"`
}

// AsModule 作为一个模块
func (dn *DataNode) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Name:        "DataNode",
		Version:     1.0,
		Description: "数据存储节点",
		OnReady: func(mctx ipakku.Loader) {
			if err := dn.ch.RegLib(CacheLib_StreamToken, CacheLib_TokenExpMilliSecond/1000); nil != err {
				logs.Panicln(err)
			}
			// hash文件写入、读取、删除控制
			dn.dhc = NewHashDataCtrl()
			// 节点编号
			dn.nodeno = mctx.GetParam("nodeno").ToString("")
			// 数据存储归档功能
			dn.df = NewDataNodeFiles(dn.fd, dn.dhc)
			// 元数据信息存储功能
			ds4datanode := "./.datas/" + mctx.GetParam(ipakku.PARAMKEY_APPNAME).ToString("app") + "#datanode?cache=shared"
			dn.dns = NewDataNodeStory("datanodes", ifilestorage.DBSetting{
				DriverName:     "sqlite3",
				DataSourceName: dn.cf.GetConfig("store.datanode.datasource").ToString(ds4datanode),
			})
			// 文件Hash记录索引
			ds4datahash := "./.datas/" + mctx.GetParam(ipakku.PARAMKEY_APPNAME).ToString("app") + "#datahash?cache=shared"
			dn.dhs = NewDataHashStory("datahashs", ifilestorage.DBSetting{
				DriverName:     "sqlite3",
				DataSourceName: dn.cf.GetConfig("store.datahash.datasource").ToString(ds4datahash),
			})
			// 远程调用namenode
			dn.nconn = bizutils.NewConn4RPC(
				dn.cf.GetConfig("namenode.rpc.address").ToString("127.0.0.1:5051"),
				bizutils.User{
					User:   dn.cf.GetConfig("namenode.rpc.user").ToString("DATANODE"),
					Passwd: dn.cf.GetConfig("namenode.rpc.pwd").ToString(""),
					PropJSON: PropJSON4DataNode{
						NodePid:     mctx.GetInstanceID(),
						NodeNo:      dn.nodeno,
						RPCAddress:  dn.cf.GetConfig("listen.rpc.endpoint").ToString(dn.cf.GetConfig("listen.rpc.address").ToString("127.0.0.1:5061")),
						HTTPAddress: dn.cf.GetConfig("listen.http.endpoint").ToString(dn.cf.GetConfig("listen.http.address").ToString("http://127.0.0.1:5062")),
					}.ToJSON(),
				})
			// 尝试登录
			if err := dn.nconn.DoLogin(); nil != err {
				logs.Panicln(err)
			}
			dn.nrpc = namenode.NewRPC4NameNode(dn.nconn)
			// 传输Token功能
			dn.tm = NewTokenManage(dn.nrpc, dn.ch)
		},
		OnSetup: func() {
			if conn, err := dn.dns.GetSqlTx(); nil == err {
				if err := dn.dns.CreateTables(conn); nil != err {
					logs.Panicln(err)
				} else {
					if err := conn.Commit(); nil != err {
						logs.Panicln(err)
					}
				}
			} else {
				logs.Panicln(err)
			}
			if conn, err := dn.dhs.GetSqlTx(); nil == err {
				if err := dn.dhs.CreateTables(conn); nil != err {
					logs.Panicln(err)
				} else {
					if err := conn.Commit(); nil != err {
						logs.Panicln(err)
					}
				}
			} else {
				logs.Panicln(err)
			}
		},
		OnInit: func() {
			go dn.checkConn()
			go dn.startHashDataCleaner()
			go dn.df.startUploadTempCleaner()
		},
	}
}
func (dn *DataNode) GetNodeNo() string {
	return dn.nconn.GetUser()
}
func (dn *DataNode) checkConn() {
	for {
		if err := dn.nconn.DoPing(); nil != err {
			logs.Errorln("Ping: ", err)
			for {
				if err = dn.nconn.DoLogin(); nil != err {
					logs.Errorln("ReLogin: ", err)
					time.Sleep(time.Second * 10)
					continue
				} else {
					logs.Infoln("ReLogin succeed")
					break
				}
			}
		}
		time.Sleep(time.Minute)
	}
}

// startHashDataCleaner 删除冗余的hash文件和数据
func (dn *DataNode) startHashDataCleaner() {
	dcl := &HashDataClear{
		nodeno:  dn.nodeno,
		dhc:     dn.dhc,
		dns:     dn.dns,
		dhs:     dn.dhs,
		fds:     dn.fd,
		nrpc:    dn.nrpc,
		delFile: dn.cf.GetConfig("clearsetting.deletefile").ToBool(false),
	}
	// Tset
	dcl.StartCleaner()
}
