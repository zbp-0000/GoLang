package myfunc

import "github.com/gin-gonic/gin"

func Hello1(context *gin.Context){
	//获取路径中的参数值：
	id := context.Param("id")
	context.String(200,"获取路径上拼接的参数值,%s",id)
}
func Hello2(context *gin.Context){
	//获取路径中的参数值：
	id := context.Param("id")
	context.String(200,"获取路径上拼接的参数值,%s",id)
}
func Hello3(context *gin.Context){
	//获取路径中的参数值：通过key获取对应的value
	id := context.Query("id")
	name := context.Query("name")
	context.String(200,"获取路径上拼接的参数值,%s,%s",id,name)
}

func Hello4(context *gin.Context){
	//获取路径中的参数值：通过key获取对应的value
	id := context.DefaultQuery("id","123")
	name := context.DefaultQuery("name","丽丽")
	context.String(200,"获取路径上拼接的参数值,%s,%s",id,name)
}

func Hello5(context *gin.Context){
	//获取路径中的参数值：通过key获取对应的value的多个参数：
	idvalues := context.QueryArray("id")
	context.String(200,"获取路径上拼接的参数值,%s",idvalues)
}
func Hello6(context *gin.Context){
	//获取路径中的参数值：通过key获取对应的value的多个参数：
	user_map := context.QueryMap("user")
	context.String(200,"获取路径上拼接的参数值,%s",user_map)
}