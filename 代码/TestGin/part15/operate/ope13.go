package main

import (
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
	//开启打印日志：
	db.LogMode(true)

	//查询操作：Raw
	//var users []demostruct.User
	//db.Raw("select * from users where age = ?",14).Find(&users)
	//fmt.Println(users)

	//增加、删除、修改 ：Exec
	//db.Exec("insert into users (age,name) values (?,?)",33,"莹莹")
	//db.Exec("delete from users where user_id = ?",1)
	//db.Exec("update users set name = ? where user_id = ?","明明",3)
}

