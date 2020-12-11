package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/huahuayu/address-generator/common/config"
	log "github.com/sirupsen/logrus"
	"time"
	"xorm.io/xorm"
)

const (
	dsn       = "%s:%s@tcp(%s:%s)/%s?%s"
	dsnParams = "charset=utf8mb4&parseTime=true&timeout=5s"
)

var (
	DB *xorm.Engine
)

func Init() {
	dbConfig := config.App.Db
	var err error
	DB, err = xorm.NewEngine("mysql", fmt.Sprintf(dsn, dbConfig.User, dbConfig.Pass, dbConfig.Host, dbConfig.Port, dbConfig.Name, dsnParams))
	if err != nil {
		log.Panic("connect to db error:: ", err)
	}

	DB.TZLocation, _ = time.LoadLocation(config.App.Server.TimezoneLoc)
	DB.SetMaxOpenConns(config.App.Db.MaxConnect)
	DB.SetMaxOpenConns(config.App.Db.MaxIdle)

	err = DB.Ping()
	if err != nil {
		log.Panic("ping db err: ", err)
	}

	if config.App.Db.ShowSQL {
		DB.ShowSQL()
	}
}
