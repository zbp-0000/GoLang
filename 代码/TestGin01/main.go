package main

import (
	"TestGin01/myfunc"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//加载模板文件：
	//r.LoadHTMLFiles("templates/hello01.html", "templates/hello02.html")
	r.LoadHTMLGlob("templates/**/*")

	//指定静态文件：指定CSS文件：
	//r.Static("/s", "static")
	r.StaticFS("/s", http.Dir("static/css"))

	r.GET("/demo01", myfunc.Hello7)

	r.Run()
}
