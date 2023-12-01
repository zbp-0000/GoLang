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
	//DefaultPostForm方法:作用：当页面中未定义表单元素进行提交给出默认值，如果页面定义了元素但是提交没有提交数据，那么不会有默认值，会认为是没有提交数据
	age := context.DefaultPostForm("age","18")
	//PostFormArray方法：作用：如果前端value数据过多可以用数组接收：
	lovelang := context.PostFormArray("lovelang")
	//PostFormMap方法:作用：获取map的数据,参数需要注意：传入的是整个map（而不是具体的key）
	usermap := context.PostFormMap("user")
	fmt.Println(uname)
	fmt.Println(pwd)
	fmt.Println(age)
	fmt.Println(lovelang)
	fmt.Println(usermap)
}
