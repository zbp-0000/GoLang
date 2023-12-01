package external

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Hello4(context *gin.Context){
	//获取路径中的参数值：
	context.HTML(200,"demo01/hello3.html",nil)
}
type User2 struct {
	Uname string `json:"uname" uri:"uname" form:"uname"`
	Age int `json:"age" uri:"age" form:"age"`
}
func Hello5(context *gin.Context){
	//创建结构体示例：
	var user User2
	//数据绑定：
	err := context.ShouldBind(&user)
	//打印结构体对象的内容：
	fmt.Println(user)
	if(err != nil ){
		context.JSON(404,gin.H{
			"msg" : "绑定失败",
		})
	}else{
		context.JSON(200,gin.H{

			"msg" : "绑定成功",
		})
	}
}

func Hello6(context *gin.Context){
	//创建结构体示例：
	var user User2
	//数据绑定：
	err := context.ShouldBindUri(&user)
	//打印结构体对象的内容：
	fmt.Println(user)
	if(err != nil ){
		context.String(404,"绑定失败")
	}else{
		context.String(200,"绑定成功")
	}
}