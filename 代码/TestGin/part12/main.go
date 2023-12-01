package main

import (
	"TestGin/part12/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//加载html页面：
	r.LoadHTMLGlob("part12/templates/**/*")
	//指定文件：
	r.Static("/s","part12/static")

	//指定总路由：
	router.Router(r)
	r.Run()
}
