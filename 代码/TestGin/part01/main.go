package main

import (
	"TestGin/part01/myfunc"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//加载html文件：
	r.LoadHTMLGlob("part01/templates/**/*")
	//指定静态文件：css文件
	r.Static("/s","part01/static")
	//写路由：
	r.GET("/demo",myfunc.Hello)


	r.Run()
}
