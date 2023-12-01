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



	//关联查询操作：（关联关系在UserInfo表中，所以从UserInfo入手）
	var userinfo demostruct.UserInfo
	db.First(&userinfo,"info_id = ?",1)
	fmt.Println(userinfo)

	var user demostruct.User
	//通过userinfo模型查出来的User字段的信息放入新的容器user中：
	db.Model(&userinfo).Related(&user,"User")
	fmt.Println(user)
	fmt.Println(userinfo)

}
