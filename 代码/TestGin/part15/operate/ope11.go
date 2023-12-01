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
	//开启打印日志：
	db.LogMode(true)
	//First:
	//var user demostruct.User
	//First:SELECT * FROM `users`  WHERE (`users`.`user_id` = 1) ORDER BY `users`.`user_id` ASC LIMIT 1
	//db.Debug().First(&user,1) -->默认情况查询的是主键
	//fmt.Println(user)
    //对应SQL：SELECT * FROM `users`  WHERE (user_id = 1) ORDER BY `users`.`user_id` ASC LIMIT 1
	//db.Debug().First(&user,"user_id = ?",1)
	//fmt.Println(user)
	//对应SQL：SELECT * FROM `users`  WHERE (user_id = 1) ORDER BY `users`.`user_id` ASC LIMIT 1 
	//db.Debug().Where("user_id = ?",1).First(&user)
	//fmt.Println(user)
	
	//FirstOrCreate
	//user2 := demostruct.User{//这里定义的结构体的实例的数值其实就是FirstOrCreate的查询条件
	//	UserId: 2,
	//	Age:    20,
	//	Name:   "菲菲",
	//	IID:    1,
	//}
	//如果有对应的数据，就查询出来，如果没有对应的数据，就会帮我们创建新的记录
	//db.FirstOrCreate(&user,user2)
	//fmt.Println(user)


	//Last:对应SQL：SELECT * FROM `users`  WHERE (`users`.`user_id` = 1) ORDER BY `users`.`user_id` DESC LIMIT 1
	//db.Debug().Last(&user,1)
	//fmt.Println(user)

	//Take:对应SQL： SELECT * FROM `users`  WHERE (`users`.`user_id` = 1) LIMIT 1
	//db.Debug().Take(&user,1)
	//fmt.Println(user)

	//Find:对应SQL：SELECT * FROM `users`  WHERE (`users`.`user_id` = 1)
	//user_id_arr := []int{1,2}
	//db.Debug().Find(&user,user_id_arr)//SELECT * FROM `users`  WHERE (`users`.`user_id` IN (1,2))
	//fmt.Println(user)

	//db.Debug().Where("user_id = ?",1).First(&user)
	//fmt.Println(user)

	//db.Debug().Where("user_id in (?)",[]int{1,2}).First(&user)
	//fmt.Println(user)

	//db.Debug().Select("name,age").Where("user_id = ?",1).First(&user)
	//fmt.Println(user)//{0 19 丽丽 0}


	//Create操作只可以插入一条技能，不能批量操作
	//user := demostruct.User{
	//	Age:    26,
	//	Name:   "小明",
	//	IID:    1,
	//}
	//
	//db.Create(&user)

	//user := demostruct.User{
	//	Age:    14,
	//	Name:   "莎莎",
	//	IID:    1,
	//}
	//
	//db.Save(&user)

	//更新：先查询再更新：
	//var user demostruct.User

	//（1）先查询，再通过Model进行操作,再Update操作：
	//db.Where("user_id = ?",1).First(&user)
	//db.Model(&user).Update("age",29)

	//（2）直接在查询之后进行操作：
	//db.Where("user_id = ?",1).First(&user).Update("name","露露")
	
	//（3）直接在查询之后进行操作，传入结构体示例，更新多个字段
	//db.Where("user_id = ?",1).First(&user).Update(demostruct.User{
	//	Age:    11,
	//	Name:   "小刚",
	//})

	//（4）直接在查询之后进行操作，传入map，更新多个字段

	//db.Where("user_id = ?",1).First(&user).Update(map[string]interface{}{
	//	"age" : 21,
	//	"name" : "小花",
	//})

	//Delete:删除数据:
	// (1)先查询再删除：
	//var user demostruct.User
	//db.Where("user_id = ?",1).First(&user)
	//db.Delete(&user)

	//（2）通过条件进行删除：
	//var user demostruct.User
	//db.Where("user_id = ?",2).Delete(&user)


	//Not:
	//var users []demostruct.User
	//db.Not("user_id = ?",1).Find(&users)
	//fmt.Println(users)
	//对应SQL：SELECT * FROM `users`  WHERE (`users`.`age` <> 18) AND (`users`.`name` <> '丽丽')
	//var users []demostruct.User
	//db.Debug().Not(demostruct.User{
	//	Age:    18,
	//	Name:   "丽丽",
	//}).Find(&users)
	//fmt.Println(users)

	//Or :
	//var users []demostruct.User
	//db.Where("user_id = ?",1).Or("user_id = ?",3).Find(&users)
	//fmt.Println(users)


	//Order:
	//var users []demostruct.User
	//db.Where("age = ?",14).Order("user_id asc").Find(&users)
	//fmt.Println(users)


	//Limit:
	//var users []demostruct.User
	//db.Limit(2).Find(&users)
	//fmt.Println(users)

	//Offset:
	//注意：Offset中设置的偏移数字为第几条记录，从0开始，0、1、2、、、、
	//注意：Offset必须和Limit结合使用
	//var users []demostruct.User
	//db.Offset(1).Limit(2).Find(&users)
	//fmt.Println(users)

	//Scan
	//type UserDemo struct {//你要扫描的结构体的字段的名字和User中的字段名字必须一致才可以扫描
	//	Name1 string
	//	Age int
	//}
	//var userdemo UserDemo
	//var user demostruct.User
	//db.Where("user_id=?",1).Find(&user).Scan(&userdemo)
	//fmt.Println(user)
	//fmt.Println(userdemo)

	//Count:
	//var users []demostruct.User
	//定义一个变量接收计数的数量：
	//var count int
	//db.Find(&users).Count(&count)
	//fmt.Println(users)
	//fmt.Println(count)

	//Group:
	//var users []demostruct.User
	//定义一个新的结构体：
	//type GroupData struct {
	//	Age int
	//	Count int
	//}
	//var group_date []GroupData
	//对应SQL：SELECT age,count(*) FROM `users`   GROUP BY age
	//对应SQL：SELECT age,count(*) as count FROM `users`   GROUP BY age HAVING (age > 18)
	//Having:在分组以后进行过滤
	//db.Debug().Select("age,count(*) as count").Group("age").Having("age > 18").Find(&users).Scan(&group_date)
	//fmt.Println(users)
	//fmt.Println(group_date)


	//Joins:
	//定义一个新的结构体用于Scan：
	type NewUserInfo struct {
		User_Id int
		Name string
		I_Id int
		Info_Id int
		Address string
	}
	var newUser []NewUserInfo
	var users []demostruct.User
	db.Select("users.user_id,users.name,users.i_id,user_infos.info_id,user_infos.address").Joins("left join user_infos on users.i_id = user_infos.info_id").Find(&users).Scan(&newUser)
	fmt.Println(users)
	fmt.Println(newUser)

}
