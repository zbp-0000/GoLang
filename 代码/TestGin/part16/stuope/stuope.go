package stuope

import (
	"TestGin/part16/dbope"
	"TestGin/part16/logs_ope"
	"TestGin/part16/models"
	"github.com/gin-gonic/gin"
)

func Hello1(context *gin.Context){
	//操作数据库：
	//添加一条记录：
	stu := models.Student{
		Name: "露露",
		Age:  21,
	}
	//添加操作
	dbope.Db.Create(&stu)
	//关闭资源：
	dbope.Db.Close()


	//日志记录操作：
	logs_ope.Logrus.Info("向数据库中增加了一条记录")

}
