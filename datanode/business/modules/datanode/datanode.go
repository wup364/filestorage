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
	"datanode/business/modules/datanode/cleaner"
	"datanode/business/modules/datanode/fio"
	"datanode/business/modules/datanode/remote"
	"datanode/business/modules/datanode/repository"
	"datanode/business/modules/datanode/scaner"
	"datanode/business/modules/datanode/versionup"
	"datanode/ifilestorage"
	"time"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/logs"
)

// DataNode 数据存储节点
type DataNode struct {
	DataNodeImpl
	nodeno string
	dhc    *fio.HashDataCtrl
	nconn  *bizutils.Conn4RPC
	nrpc   ifilestorage.RPC4NameNode
	ch     ipakku.AppCache        `@autowired:"AppCache"`
	fd     ifilestorage.FileDatas `@autowired:"FileDatas"`
	config *ConfigBean            `@autoConfig:""`
}

// AsModule 作为一个模块
func (dn *DataNode) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Name:        "DataNode",
		Version:     1.1,
		Description: "数据存储节点",
		Updaters: func(mctx ipakku.Loader) ipakku.Updaters {
			return []ipakku.Updater{
				&versionup.Updater_1_1{DHS: dn.dhs},
			}
		},
		OnReady: dn.onReady,
		OnSetup: dn.onSetup,
		OnInit:  dn.onInit,
	}
}

func (dn *DataNode) onReady(mctx ipakku.Loader) {
	if err := dn.ch.RegLib(fio.CacheLib_StreamToken, fio.CacheLib_TokenExpMilliSecond/1000); nil != err {
		logs.Panicln(err)
	}
	// hash文件写入、读取、删除控制
	dn.dhc = fio.NewHashDataCtrl()
	// 节点编号
	dn.nodeno = mctx.GetParam("nodeno").ToString("")
	// 数据存储归档功能
	dn.df = fio.NewDataNodeFiles(dn.fd, dn.dhc)
	// 元数据信息存储功能
	ds4datanode := "./.datas/" + mctx.GetParam(ipakku.PARAMKEY_APPNAME).ToString("app") + "#datanode?cache=shared"
	dn.dns = repository.NewDataNodeRepo("datanodes", ifilestorage.DBSetting{
		DriverName:     "sqlite3",
		DataSourceName: dn.config.datanodeDatasource.ToString(ds4datanode),
	})
	// 文件Hash记录索引
	ds4datahash := "./.datas/" + mctx.GetParam(ipakku.PARAMKEY_APPNAME).ToString("app") + "#datahash?cache=shared"
	dn.dhs = repository.NewDataHashRepo("datahashs", ifilestorage.DBSetting{
		DriverName:     "sqlite3",
		DataSourceName: dn.config.datahashDatasource.ToString(ds4datahash),
	})
	// 远程调用namenode
	dn.nconn = bizutils.NewConn4RPC(
		dn.config.namenode.rpcAddr,
		bizutils.User{
			User:   dn.config.namenode.rpcUser,
			Passwd: dn.config.namenode.rpcPwd,
			PropJSON: PropJSON4DataNode{
				NodePid:     mctx.GetInstanceID(),
				NodeNo:      dn.nodeno,
				RPCAddress:  dn.config.listen.rpcEndpoint.ToString(dn.config.listen.rpcAddress),
				HTTPAddress: dn.config.listen.httpEndpoint.ToString(dn.config.listen.httpAddress),
			}.ToJSON(),
		})
	// 尝试登录
	if err := dn.nconn.DoLogin(); nil != err {
		logs.Panicln(err)
	}
	dn.nrpc = remote.NewRPC4NameNode(dn.nconn)
	// 传输Token功能
	dn.tm = fio.NewTokenManage(dn.nrpc, dn.ch)
}

func (dn *DataNode) onSetup() {
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
}

func (dn *DataNode) onInit() {
	go dn.checkConn()

	// 文件hash块清理程序
	go cleaner.NewHashDataCleaner(
		dn.nodeno,
		dn.dhc,
		dn.dns,
		dn.dhs,
		dn.fd,
		dn.nrpc,
		dn.config.clearsetting.deletefile).StartCleaner()

	// 上传缓存清理程序
	go cleaner.NewUploadTempCleaner(dn.fd).StartCleaner()

	// 从数据库中读取hash, 并验证文件是否还存在
	go scaner.NewDataHashCanerForRepo(dn.fd, dn.dhs).StartScaner()
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

func (dn *DataNode) GetNodeNo() string {
	return dn.nconn.GetUser()
}
