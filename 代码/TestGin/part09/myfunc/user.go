package myfunc

import (
	"github.com/gin-gonic/gin"
	"time"
)
type Student struct {
	Age int
	Name string
}
func Hello1(context *gin.Context){
	//定义数据：
	age := 19
	arr := []int{33,66,99}
	flag := true
	username := "丽丽"
	//创建结构体实例：
	stu := Student{
		Age : 18,
		Name : "丽丽",
	}
	now_time := time.Now()
	//将age 和arr放入map中：
	map_data := map[string]interface{}{
		"age" : age,
		"arr" : arr,
		"flag" : flag,
		"username" : username,
		"stu" : stu,
		"now_time" : now_time,
	}
	//获取路径中的参数值：
	context.HTML(200,"demo01/hello.html",map_data)
}


//定义一个函数：
func Add(num1 int,num2 int) int{
	return num1+num2
}
