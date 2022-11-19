package controller

import (
	"io/ioutil"
	"net/http"
	"yqdk/base"
	"yqdk/src/api/service"
	e "yqdk/src/global/error_handler/base"
	"yqdk/src/punch/punchLogic"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func InsertUser(c *gin.Context) {
	glog.Infof("inser user...")
	var user base.UserAccount
	if err := c.ShouldBindJSON(&user); err != nil {
		glog.Errorf("insert user bind json error, msg:[%s]", err.Error())
		return
	}
	if ok, err := check(user); !ok {
		glog.Errorf("insert user error, msg: [%s]", err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	var u []base.UserAccount
	u = append(u, user)
	if ok, err := service.Insert(u); !ok {
		err = err.(e.GinError)
		glog.Errorf("insert user error, msg: [%s]", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "添加成功",
	})
}

func ReadLog(c *gin.Context) {
	out, err := ioutil.ReadFile("./out.out")
	if err != nil {
		glog.Errorf("read out.out error, msg:[%s]", err.Error())
		c.JSON(500, gin.Error{
			Err: err,
		})
	}
	c.JSON(200, gin.H{
		"message": string(out),
	})
}

func Delete(c *gin.Context) {
	glog.Info("delete user...")
	var user base.UserAccount
	if err := c.ShouldBindJSON(&user); err != nil {
		glog.Errorf("insert user bind json error, msg:[%s]", err.Error())
		return
	}
	// 检查账号密码是否正确, 错误的话不允许删除
	if ok, err := check(user); !ok {
		glog.Errorf("delete user error msg: [%s]", err)
		c.JSON(http.StatusOK, gin.H{
			"message": err,
		})
		return
	}
	if ok, err := service.Delete(user); !ok {
		err = err.(e.GinError)
		glog.Errorf("delete user error msg: [%s]", err.Error())
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "删除成功",
	})
}

func check(user base.UserAccount) (bool, string) {
	if user.Username == "" {
		return false, "用户名不能为空"
	}

	if user.Password == "" {
		return false, "密码不能为空"
	}
	// 校验密码是否正确
	data := punchLogic.BuildRequestData(user)
	var dc = base.DataForm{Data: data}
	s1 := punchLogic.GetCode(dc.Data)
	if s1 == "" {
		glog.Error("登录错误")
		return false, "[" + user.Username + "] 用户名或密码错误"
	}

	return true, ""
}
