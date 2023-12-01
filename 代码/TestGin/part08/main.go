package main

import (
	"TestGin/part08/myfunc"
	"github.com/gin-gonic/gin"

)

func main() {
	r := gin.Default()
	//写路由：
	//定义路由：
	r.GET("/red1",myfunc.Red1)
	r.GET("/red2",myfunc.Red2)
	r.Run()
}
