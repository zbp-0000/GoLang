package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func MiddleWare01(context *gin.Context){
	fmt.Println("这是自定义的中间件1-开始")
	//context.Next()
	fmt.Println("这是自定义的中间件1-结束")
}

//gin.HandlerFunc 等价于  func(*Context)函数
//所以MiddleWare02就必须有个返回值，并且返回值是一个函数
func MiddleWare02() gin.HandlerFunc{
	return func(context *gin.Context){
		fmt.Println("这是自定义的中间件2-开始")
		fmt.Println("这是自定义的中间件2-结束")
	}
}


func MiddleWare03() gin.HandlerFunc{
	return func(context *gin.Context){
		fmt.Println("这是自定义的中间件3-开始")
		if 4 > 2 { //满足条件
			//终止链条：
			//context.Abort()
			return
		}
		context.Next()
		fmt.Println("这是自定义的中间件3-结束")
	}
}
