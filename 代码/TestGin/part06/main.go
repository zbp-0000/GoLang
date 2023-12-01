package main

import (
	"TestGin/part06/myfunc"
	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	//写路由：
	//加载html页面：
	r.LoadHTMLGlob("part06/templates/**/*")
	//指定js文件：
	r.Static("/s","part06/static")
	//定义路由：
	r.GET("/userindex",myfunc.Hello1)
	r.POST("/savefile",myfunc.Hello4)
	r.Run()
}
