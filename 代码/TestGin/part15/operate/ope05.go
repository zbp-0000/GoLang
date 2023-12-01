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



	//关联删除
	//先查询
	var userinfo demostruct.UserInfo
	db.Preload("User").Find(&userinfo,"info_id = ?",1)
	fmt.Println(userinfo)
	//再删除：借助userinfo模型删除User记录
	db.Delete(&userinfo.User) //UserInfo中信息没有被删除，删除的是关联的User表中的记录
	db.Delete(&userinfo)
}
