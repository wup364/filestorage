package versionup

import (
	"database/sql"
	"datanode/business/modules/datanode/repository"

	"github.com/wup364/pakku/ipakku"
)

// Updater_1_1 模块版本升级执行器
type Updater_1_1 struct {
	DHS *repository.DataHashRepo
}

// Version 要升级到的版本号
func (up Updater_1_1) Version() float64 {
	return 1.1
}

// Execute 执行升级
func (up Updater_1_1) Execute(mctx ipakku.Loader) (err error) {
	var sqlTx *sql.Tx
	table := up.DHS.GetTable()
	if sqlTx, err = up.DHS.GetSqlTx(); nil != err {
		return
	}

	sqlStr := "alter table " + table + " add  `scanmarker` int not null default 0; create index " + table + "_index_scanmarker on " + table + " (scanmarker);"
	if _, err = sqlTx.Exec(sqlStr); nil != err {
		sqlTx.Rollback()
	} else {
		err = sqlTx.Commit()
	}
	return
}
