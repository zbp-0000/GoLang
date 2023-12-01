package myfunc

import "github.com/gin-gonic/gin"

func Hello(context *gin.Context) { //点一下，gin的包就会自动引入了
	name := "你好我是珊珊老师"
	context.HTML(200, "demo01/hello01.html", name)
}

type Student struct {
	Name string
	Age  int
}

func Hello6(context *gin.Context) { //点一下，gin的包就会自动引入了
	//创建Map：
	var a map[string]Student = make(map[string]Student, 2)
	//将键值对存入map：
	a["no1"] = Student{
		Name: "丽丽",
		Age:  18,
	}
	a["no2"] = Student{
		Name: "菲菲",
		Age:  21,
	}
	context.HTML(200, "demo01/hello01.html", a)
}
func Hello4(context *gin.Context) { //点一下，gin的包就会自动引入了
	//定义一个结构体类型的数组：
	var arr [3]Student
	arr[0] = Student{
		Name: "丽丽",
		Age:  18,
	}
	arr[1] = Student{
		Name: "菲菲",
		Age:  21,
	}
	arr[2] = Student{
		Name: "小明",
		Age:  27,
	}
	context.HTML(200, "demo01/hello01.html", arr)
}
func Hello5(context *gin.Context) { //点一下，gin的包就会自动引入了
	//创建Map：
	var a map[string]int = make(map[string]int, 3)
	//将键值对存入map：
	a["丽丽"] = 18
	a["菲菲"] = 16
	a["明明"] = 21
	context.HTML(200, "demo01/hello01.html", a)
}
func Hello7(context *gin.Context) { //点一下，gin的包就会自动引入了
	//创建切片：
	slice := []int{1, 2, 3, 4, 5, 6}
	context.HTML(200, "demo01/hello01.html", slice)
}
func Hello2(context *gin.Context) { //点一下，gin的包就会自动引入了
	//创建结构体示例：
	s := Student{
		Name: "丽丽",
		Age:  18,
	}
	context.HTML(200, "demo01/hello01.html", s)
}

func Hello3(context *gin.Context) { //点一下，gin的包就会自动引入了
	//定义一个数组：
	var arr [3]int
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
	context.HTML(200, "demo01/hello01.html", arr)
}
