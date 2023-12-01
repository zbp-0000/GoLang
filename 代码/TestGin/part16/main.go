package main

import (
	_ "TestGin/part16/dbope"
	_ "TestGin/part16/logs_ope"
	"TestGin/part16/router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//加载html页面：
	r.LoadHTMLGlob("part16/templates/**/*")
	//指定文件：
	r.Static("/s","part16/static")

	//指定总路由：
	router.Router(r)

	r.Run()
}
