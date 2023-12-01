package myfunc

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func Hello1(context *gin.Context){
	//获取路径中的参数值：
	context.HTML(200,"demo01/hello.html",nil)
}

func Hello2(context *gin.Context){
	//获取前端传入的文件：
	file,_ := context.FormFile("myfile")
	fmt.Println(file.Filename)

	//加入一个时间戳：
	time_int := time.Now().Unix()
	time_str := strconv.FormatInt(time_int,10) //10:十进制
	//保存在我的本地：
	context.SaveUploadedFile(file,"e://" + time_str + file.Filename)

	//响应一个字符串：
	context.String(200,"文件上传成功")

}


func Hello3(context *gin.Context){
	//先获取form表单
	form,_ := context.MultipartForm()
	//在form表单中获取name相同的文件：
	files := form.File["myfile"]  //File是个Map，通过key获取value部分

	//files就是name相同的多个文件：挨个处理---遍历处理：
	for _,file := range files{
		//加入一个时间戳：
		time_int := time.Now().Unix()
		time_str := strconv.FormatInt(time_int,10) //10:十进制
		//保存在我的本地：
		context.SaveUploadedFile(file,"e://" + time_str + file.Filename)
	}

	//响应一个字符串：
	context.String(200,"文件上传成功")

}
