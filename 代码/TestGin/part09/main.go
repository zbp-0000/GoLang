package main

import (
	"TestGin/part09/myfunc"
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	r := gin.Default()
	//注册函数：FuncMap是html/FuncMap
	r.SetFuncMap(template.FuncMap{
		//键值对的作用：key指定前端调用的名字，value指定的是后端对应的函数
		"add" : myfunc.Add,
	})
	//写路由：
	//加载html页面：
	r.LoadHTMLGlob("part09/templates/**/*")
	//定义路由：
	r.GET("/userindex",myfunc.Hello1)
	r.Run()
}
