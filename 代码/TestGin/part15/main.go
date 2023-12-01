package main

import (
	"TestGin/part15/demostruct"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main(){
	//连接数据库：
	db,err := gorm.Open("mysql","root:root@tcp(localhost:3306)/testgorm?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic(err) //如果出错，后续代码没有必要执行，想让程序中断，panic来执行即可
	}

	//数据库资源释放：
	defer db.Close()

	//创建表：通常情况下，数据库中新建的标的名字是结构体名字的复数形式，例如结构体User，表名 users
	db.CreateTable(&demostruct.User{})
	db.CreateTable(&demostruct.UserInfo{})

	//db.CreateTable(&demostruct.Author{})
	//db.CreateTable(&demostruct.Article{})

	//db.CreateTable(&demostruct.Student{})
	//db.CreateTable(&demostruct.Course{})

}
