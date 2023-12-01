package main

import (
	"TestGin/part03/myfunc"
	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	//写路由：
	//加载html页面：
	r.LoadHTMLGlob("part03/templates/**/*")
	//定义路由：
	r.GET("/userindex",myfunc.Hello1)
	r.POST("/getUserInfo",myfunc.Hello2)
	r.Run()
}
