package external

import (
	"TestGin/part12/bill"
	"TestGin/part13/middleware"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.RouterGroup){

	r.GET("/userindex3",middleware.MiddleWare01,Hello4)
	r.POST("/toajax",Hello5)
	r.GET("/userindex4/:uname/:age",Hello6)
	r.GET("/userindex4/丽丽/18",bill.Hello1)
}
