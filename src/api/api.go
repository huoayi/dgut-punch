package api

import (
	"net/http"
	"yqdk/src/api/controller"
	"yqdk/src/punch"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

func Start() {
	r := gin.Default()
	r.Use(cors())
	r.POST("/insert", insertUser)
	r.GET("/run", run)
	r.GET("/read", read)
	r.POST("/delete", delete)
	err := r.Run(":4398")
	if err != nil {
		glog.Errorf("run gin error, msg:[%s]", err)
		return
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin") //请求头部
		if origin != "" {
			//接收客户端发送的origin （重要！）
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			//服务器支持的所有跨域请求的方法
			c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
			//允许跨域设置可以返回其他子段，可以自定义字段
			c.Header("Access-Control-Allow-Headers", "Authorization,Content-Type, Content-Length, X-CSRF-Token, Token,session")
			// 允许浏览器（客户端）可以解析的头部 （重要）
			c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers")
			//设置缓存时间
			c.Header("Access-Control-Max-Age", "172800")
			//允许客户端传递校验信息比如 cookie (重要)
			c.Header("Access-Control-Allow-Credentials", "true")
		}

		//允许类型校验
		if method == "OPTIONS" {
			c.JSON(http.StatusOK, "ok!")
		}
		defer func() {
			err := recover()
			if err != nil {
				glog.Errorf("cros function is error, msg: %s", err)
			}
		}()

		c.Next()
	}
}

func insertUser(c *gin.Context) {
	controller.InsertUser(c)
	glog.Flush()
}

func run(c *gin.Context) {
	punch.Start()
	c.JSON(http.StatusOK, gin.H{
		"message": "运行成功, 具体运行结果请查看日志",
	})
	glog.Flush()
}

func read(c *gin.Context) {
	controller.ReadLog(c)
	glog.Flush()
}

func delete(c *gin.Context) {
	controller.Delete(c)
	glog.Flush()
}
