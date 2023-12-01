package main

import (
	"TestGin/part04/myfunc"
	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	//写路由：
	//加载html页面：
	r.LoadHTMLGlob("part04/templates/**/*")
	//指定js文件：
	r.Static("/s","part04/static")
	//定义路由：
	r.GET("/userindex",myfunc.Hello1)
	r.POST("/getUserInfo",myfunc.Hello2)
	r.POST("/ajaxpost",myfunc.Hello3)
	r.Run()
}
