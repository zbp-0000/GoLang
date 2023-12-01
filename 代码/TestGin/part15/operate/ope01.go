package main

import (
	"TestGin/part15/demostruct"
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
	//db.CreateTable(&demostruct.User{})
	//db.CreateTable(&demostruct.UserInfo{})

	//关联添加数据：
	userinfo := demostruct.UserInfo{
		Pic:     "/upload/1.jpg",
		Address: "北京海淀区",
		Email:   "124234@126.com",
		User:    demostruct.User{
			Age : 19,
			Name : "丽丽",
		},
	}

	db.Create(&userinfo)

	//关联查询操作：（关联关系在UserInfo表中，所以从UserInfo入手）
	//var userinfo demostruct.UserInfo
	////如果只是执行下面这步操作，那么关联的User信息是查询不到的：
	//db.Debug().First(&userinfo,"info_id = ?",1)
	////fmt.Println(userinfo)//{1 /upload/1.jpg 北京海淀区 124234@126.com {0 0  0}}
	//
	////如果想要查询到User相关内容，必须执行如下操作：
	////Model参数：要查询的表数据，Association参数：关联到的具体的模型：模型名字User（字段名字）
	////Find参数：查询的数据要放在什么字段中&userinfo.User
	//db.Debug().Model(&userinfo).Association("User").Find(&userinfo.User)
	//fmt.Println(userinfo)//{1 /upload/1.jpg 北京海淀区 124234@126.com {1 19 丽丽 1}}

}
