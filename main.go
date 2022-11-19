package main

import (
	_ "encoding/json"
	"flag"
	"yqdk/secret"
	"yqdk/src/api"
	"yqdk/src/db"
	"yqdk/src/global"
	"yqdk/src/punch"
	"yqdk/src/punch/punchLogic"

	"github.com/golang/glog"
)

func init() {
	// global glog
	flag.Parse()
	secret.InitKey()
	// 读文件
	global.Init()
	// 加载db
	if global.Config.DataSource == 1 {
		db.Init()
	}
}

func main() {
	defer glog.Flush()
	if global.Config.DataSource == 1 {
		punchLogic.CornTask()
		api.Start()
	} else {
		punch.Start()
	}
}
