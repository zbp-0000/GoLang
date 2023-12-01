package main

import (
	"TestGin/part10/myfunc"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//加载html页面：
	r.LoadHTMLGlob("part10/templates/**/*")
	//指定文件：
	r.Static("/s","part10/static")
	//定义路由：
	r.GET("/userindex",myfunc.Hello1)
	r.GET("/toFormBind",myfunc.Hello2)
	r.GET("/userindex2",myfunc.Hello3)

	r.GET("/userindex3",myfunc.Hello4)
	r.POST("/toajax",myfunc.Hello5)

	r.GET("/userindex4/:uname/:age",myfunc.Hello6)
	r.GET("/userindex4/丽丽/18",myfunc.Hello1)


	r.Run()
}
