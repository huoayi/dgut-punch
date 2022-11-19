package global

import (
	"github.com/golang/glog"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"yqdk/base"
)

const CONFIG_FILE = "config.yaml"

var Config *base.ConfigApplication

func Init() {
	readFile()
}

func readFile() {
	glog.Infof("read file [%s]", CONFIG_FILE)
	bytes, err := ioutil.ReadFile("./" + CONFIG_FILE)
	var ca base.ConfigApplication
	if err != nil {
		glog.Errorf("read file [%s] error, msg:[%s],", CONFIG_FILE, err.Error())
		return
	}
	err = yaml.Unmarshal(bytes, &ca)
	if err != nil {
		glog.Errorf("Unmarshal user account error: %s", err.Error())
		return
	}
	Config = &base.ConfigApplication{
		DB:          ca.DB,
		UserAccount: ca.UserAccount,
		FuncMAX:     ca.FuncMAX,
		DataSource:  ca.DataSource,
	}
}
