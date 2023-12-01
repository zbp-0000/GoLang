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



	//关联更新
	//先查询
	var userinfo demostruct.UserInfo
	db.Preload("User").Find(&userinfo,"info_id = ?",1)
	fmt.Println(userinfo)
	//再更新：注意：Update的参数age可以用结构体中字段Age也可以用数据库age字段
	db.Model(&userinfo.User).Update("age",31)
	fmt.Println(userinfo)
}
