package main

import (
	"TestGin/part15/demostruct"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
func main() {
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

	//关联添加数据：
	//userinfo := demostruct.UserInfo{
	//	Pic:     "/upload/1.jpg",
	//	Address: "北京海淀区",
	//	Email:   "124234@126.com",
	//	User:    demostruct.User{
	//		Age : 19,
	//		Name : "丽丽",
	//	},
	//}
	//
	//db.Create(&userinfo)

	//关联查询操作：（关联关系在UserInfo表中，所以从UserInfo入手）
	var userinfo demostruct.UserInfo
	db.Debug().Preload("User").Find(&userinfo,"info_id = ?",1)
	fmt.Println(userinfo)


}
