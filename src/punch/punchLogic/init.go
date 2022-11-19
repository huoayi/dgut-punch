package punchLogic

import (
	"yqdk/base"
	"yqdk/secret"
	db2 "yqdk/src/db"

	"github.com/golang/glog"
)

func LoadDataBaseConfig() []base.UserAccount {
	db := db2.DB
	var dal []base.UserAccount
	db.DataBase.Table(db.Table).Find(&dal)
	glog.Infof("load database success")
	var u []base.UserAccount
	for _, row := range dal {
		var v = base.UserAccount{
			Username: row.Username,
			Password: string(secret.Decrypt(row.Password)),
		}
		// fmt.Println(v.Password)
		u = append(u, v)
	}
	return u
}

func BuildRequestData(user base.UserAccount) map[string]string {
	execution := GetExecuation()
	return base.SetForm(user, execution)

}
