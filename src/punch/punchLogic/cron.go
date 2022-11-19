package punchLogic

import (
	"github.com/golang/glog"
	"github.com/robfig/cron"
	"net/http"
)

func CornTask() {
	c := cron.New()
	c.AddFunc("0 0 1,2 * * ?", func() {
		glog.Info("corn Task run...")
		http.Get("http://127.0.0.1:4398/run")
	})
	c.Start()
}
