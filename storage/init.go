package storage

import (
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"

	"dashboard/logger"
)

//定义orm引擎
var engine *xorm.Engine

func init() {
	var err error
	engine, err = xorm.NewEngine("sqlite3", "dashboard.db")
	if err != nil {
		logger.Error("init mysql:Error[" + err.Error() + "]")
		panic(err)
	}
	if err := engine.Sync(new(User)); err != nil {
		logger.Error("sync table user:Error[" + err.Error() + "]")
		panic(err)
	}
	if err := engine.Sync(new(Role)); err != nil {
		logger.Error("sync table role:Error[" + err.Error() + "]")
		panic(err)
	}
	logger.Info("init mysql success")
}
