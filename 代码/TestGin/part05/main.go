package main

import (
	"TestGin/part05/myfunc"
	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	//写路由：
	//加载html页面：
	r.LoadHTMLGlob("part05/templates/**/*")
	//定义路由：
	r.GET("/userindex",myfunc.Hello1)
	r.POST("/savefile",myfunc.Hello3)
	r.Run()
}
