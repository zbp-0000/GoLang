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


	//关联添加操作：关联关系在Student中，所以我们模型操作的也是Student:
	stu := demostruct.Student{
		SNo:    1001,
		Name:   "丽丽",
		Sex:    "女",
		Age:    18,
		Course: []demostruct.Course{
			{
				CName:"c++",
				TeacherName:"张三",
				Room:"s-103",
			},
			{
				CName:"高数",
				TeacherName:"李四",
				Room:"s-801",
			},
		},
	}

	db.Create(&stu)



}


