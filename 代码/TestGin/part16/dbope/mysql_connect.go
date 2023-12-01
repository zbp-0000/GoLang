package dbope

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
//提取DB，ERR：
var Db *gorm.DB
var Err error
//init函数：初始化操作：
func init() {
	//连接数据库：
	Db,Err = gorm.Open("mysql","root:root@tcp(localhost:3306)/testgorm?charset=utf8&parseTime=True&loc=Local")

	if Err != nil {
		panic(Err) //如果出错，后续代码没有必要执行，想让程序中断，panic来执行即可
	}

	//创建表：
	//Db.CreateTable(&models.Student{})

}

