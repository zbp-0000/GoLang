package myfunc

import "github.com/gin-gonic/gin"

func Hello(context *gin.Context){
	context.HTML(200,"demo01/hello.html",nil)
}
