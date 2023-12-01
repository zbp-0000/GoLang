package myfunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Red1(context *gin.Context){
	fmt.Println("这是Red1")
	//发送一个重定向的请求：
	context.Redirect(http.StatusFound,"/red2")
}


func Red2(context *gin.Context){
	fmt.Println("这是Red2")
	//在浏览器响应字符串：
	context.String(http.StatusOK,"重定向成功-这里是Red2")
}
