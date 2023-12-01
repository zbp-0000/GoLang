package myfunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func Hello1(context *gin.Context){
	//获取路径中的参数值：
	context.HTML(200,"demo01/hello.html",nil)
}

func Hello2(context *gin.Context){
	time.Sleep(time.Second * 10)
	//获取post请求的参数：
	//PostForm方法：作用：通过key得到value数据
	uname := context.PostForm("username")
	pwd := context.PostForm("pwd")
	fmt.Println(uname)
	fmt.Println(pwd)
}

//ajax的后端的处理
func Hello3(context *gin.Context){
	//获取post-ajax请求的数据，获取对应的参数：
	uname := context.PostForm("uname")
	fmt.Println(uname)
	fmt.Println(uname == "丽丽")
	//如果获取的数据和"丽丽"一样，那么就在前端响应-用户名录入重复：
	if uname == "丽丽" {
		//向浏览器返回数据，返回json格式数据：
		//mapdata := map[string]interface{}{
		//	"msg" : "用户名重复了！",
		//}
		//context.JSON(200,mapdata)

		context.JSON(200,gin.H{
			"msg" : "用户名重复了！",
		})
	}else {
		context.JSON(200,gin.H{
			"msg" : "",
		})
	}
}
