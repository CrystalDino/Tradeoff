package core

import (
	"errors"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	DBEngine *xorm.Engine
	dsn      string
)

func init() {
	dsn = GetConfigString("dsn")
	if dsn == "" {
		log.Fatalln("db engine init error, dsn not set.")
		return
	}
	var err error
	if DBEngine, err = xorm.NewEngine("mysql", dsn); err != nil {
		log.Fatalln("db engine init error,", err)
		return
	}
}

func DBLinkTest() (err error) {
	if DBEngine == nil {
		return errors.New("link error")
	}
	return DBEngine.Ping()
}

func SyncModels(beans interface{}) (err error) {
	return DBEngine.Sync2(beans)
}
