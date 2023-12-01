package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Account struct {
	Id int `gorm:"primary_key"`
	Name string
	Balance int
}
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

	//创建对应的表：
	//db.CreateTable(&Account{})

	//添加数据：
	//db.Create(&Account{
	//	Name:    "丽丽",
	//	Balance: 1000,
	//})
	//db.Create(&Account{
	//	Name:    "菲菲",
	//	Balance: 2000,
	//})


	//转账操作
	//查询丽丽和菲菲的数据：
	var lili_account Account
	var feifei_account Account
	db.Where("id = ?",1).Find(&lili_account)
	db.Where("id = ?",2).Find(&feifei_account)
	fmt.Println(lili_account)
	fmt.Println(feifei_account)

	//再更新：
	//开启事务：
	t := db.Begin()
	r1 := db.Model(&lili_account).Update("balance",lili_account.Balance - 200)
	if r1.RowsAffected == 0{
		t.Rollback()//回滚
		return
	}
	r2 := db.Model(&feifei_account).Update("balance",feifei_account.Balance + 200)
	if r2.RowsAffected == 0{
		t.Rollback()//回滚
		return
	}
	//事务提交：
	t.Commit()


}



//if r1.RowsAffected == 0 || r2.RowsAffected == 0 {
//t.Rollback()//回滚
//return
//}
