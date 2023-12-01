package main

import (
	"TestGin/part13/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//加载html页面：
	r.LoadHTMLGlob("part13/templates/**/*")
	//指定文件：
	r.Static("/s","part13/static")

	//使用中间件：
	//r.Use(middleware.MiddleWare01)
	//方式2中参数中需要对函数进行调用
	//r.Use(middleware.MiddleWare03())
	//r.Use(middleware.MiddleWare02())

	//指定总路由：
	router.Router(r)

	r.Run()
}
