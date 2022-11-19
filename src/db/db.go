package db

import (
	"github.com/golang/glog"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	it "yqdk/src/global"
)

type DBConfig struct {
	DataBase *gorm.DB
	Table    string
}

var DB *DBConfig

func Init() {
	initDB()
}

func initDB() {
	dbInfo := it.Config.DB
	dsn := dbInfo.DBUsername + ":" + dbInfo.DBPassword + "@tcp(" + dbInfo.DBHost + ":" + dbInfo.DBPort + ")/" + dbInfo.DBName + "?" + "charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		glog.Errorf("connect db error, msg:[%s]", err.Error())
	}
	DB = &DBConfig{
		DataBase: db,
		Table:    dbInfo.DBTable,
	}
}
