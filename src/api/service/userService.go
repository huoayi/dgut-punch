package service

import (
	"net/http"
	"yqdk/base"
	"yqdk/secret"
	db2 "yqdk/src/db"
	e "yqdk/src/global/error_handler/base"

	"github.com/golang/glog"
)

func Insert(us []base.UserAccount) (bool, error) {
	var info = db2.DB
	var db = info.DataBase
	glog.Infof("insert user into database...")

	var u []base.UserAccount
	// 加密
	for _, row := range us {
		var tmpu base.UserAccount
		// 当err为空时, 说明db在数据库中找到了数据, 说明数据库中已经有这条数据了, 应该取消
		if err := db.Table(info.Table).Where("username = ?", row.Username).First(&tmpu).Error; err == nil {
			return false, e.GinError{
				Status:  http.StatusOK,
				Code:    1001,
				Message: "该账号[" + row.Username + "]已存在, 无需重复录入",
			}
		}
		glog.Infof("insert data username: [%s]...", row.Username)
		encrypt := secret.Encrypt(row.Password)
		var v = base.UserAccount{
			Username: row.Username,
			Password: encrypt,
		}
		u = append(u, v)
	}
	db.Table(info.Table).Create(&u)
	return true, nil
}

func Delete(us base.UserAccount) (bool, error) {
	var info = db2.DB
	var db = info.DataBase
	var tmpu base.UserAccount
	db_ops := db.Table(info.Table).Where("username = ?", us.Username)
	if err := db_ops.First(&tmpu).Error; err != nil {
		return false, e.GinError{
			Status:  http.StatusOK,
			Code:    1002,
			Message: "该账号[" + us.Username + "]不存在, 请检查是否输入错误",
		}
	}

	glog.Infof("delete data username: [%s]...", us.Username)
	// 这里不做数据库正确性的校验了, 只要能过中央认证就行
	db_ops.Delete(&tmpu)
	return true, nil
}
