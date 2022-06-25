// Copyright (C) 2021 WuPeng <wup364@outlook.com>.
// Use of this source code is governed by an MIT-style.
// Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction,
// including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software,
// and to permit persons to whom the Software is furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
// IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

// 文件命名空间管理, 文件操作(新建、删除、移动、复制等)

package namenode

import (
	"database/sql"
	"namenode/business/modules/namenode/datanodes"
	"namenode/ifilestorage"

	"github.com/wup364/pakku/ipakku"
	"github.com/wup364/pakku/utils/fileutil"
	"github.com/wup364/pakku/utils/logs"
)

// NameNode 元数据节点
type NameNode struct {
	NameNodeImpl
	cf ipakku.AppConfig `@autowired:"AppConfig"`
}

// AsModule 作为一个模块
func (n *NameNode) AsModule() ipakku.Opts {
	return ipakku.Opts{
		Name:        "NameNode",
		Version:     1.0,
		Description: "元数据节点",
		OnReady: func(mctx ipakku.Loader) {
			if err := mctx.AutoWired(&n.NameNodeImpl); nil != err {
				logs.Panicln(err)
			} else {
				n.da = &Addr4Deleted{}
				n.d = datanodes.NewConn4DataNodes(n.e)
				mctx.SetParam("Conn4DataNodes", n.d) // 用于依赖注入
			}
			//
			deftDataSource := "./.datas/" + mctx.GetParam(ipakku.PARAMKEY_APPNAME).ToString("app") + "#deleted?cache=shared"
			confDataSource := n.cf.GetConfig("store.deletedaddr.datasource").ToString(deftDataSource)
			if deftDataSource == confDataSource {
				n.mkSqliteDIR()
			}
			n.da.Initial("deletedaddr", ifilestorage.DBSetting{
				DriverName:     n.cf.GetConfig("store.deletedaddr.driver").ToString("sqlite3"),
				DataSourceName: confDataSource,
			})
			// token缓存库
			if err := n.c.RegLib(CacheLib_StreamToken, CacheLib_StreamTokenExpSecond); nil != err {
				logs.Panicln(err)
			}
			// 注册节点删除事件, 用于删除datanode冗余数据
			if err := n.e.ConsumerSyncEvent("filenames", "deletenode", func(v interface{}) (err error) {
				tnode := v.(ifilestorage.TNode)
				var conn *sql.Tx
				if conn, err = n.da.GetSqlTx(); nil == err {
					if err = n.da.InsertItem(conn, tnode.Addr); nil == err {
						err = conn.Commit()
					} else {
						conn.Rollback()
					}
				}
				return err
			}); nil != err {
				logs.Panicln(err)
			}
		},
		OnSetup: func() {
			var err error
			var conn *sql.Tx
			if conn, err = n.da.GetSqlTx(); nil == err {
				if err = n.da.CreateTables(conn); nil == err {
					err = conn.Commit()
				} else {
					conn.Rollback()
				}
			}
			if nil != err {
				logs.Panicln(err)
			}
		},
		OnInit: func() {
		},
	}
}

// mkSqliteDIR 创建sqlite文件存放目录
func (n *NameNode) mkSqliteDIR() {
	if !fileutil.IsExist("./.datas") {
		if err := fileutil.MkdirAll("./.datas"); nil != err {
			logs.Panicln(err)
		}
	}
}
