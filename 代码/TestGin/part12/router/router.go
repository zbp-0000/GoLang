package router

import (
	"TestGin/part12/bill"
	"TestGin/part12/external"
	"github.com/gin-gonic/gin"
)

func Router(r *gin.Engine){
	b := r.Group("/bill") //支票模块的分组
	e := r.Group("/external") //第三方工具的分组

	//模块分组：
	bill.Router(b)
	external.Router(e)
}
