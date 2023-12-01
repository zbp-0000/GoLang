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



	//关联添加：一对多关系，一个作者对应多个文章，关联关系在作者中，所以我们操作的模型是作者的模型：
	author := demostruct.Author{
		Name:    "张三",
		Age:     30,
		Sex:     "男",
		Article: []demostruct.Article{
			{
				Title:"HTML入门",
				Content:"HTML******",
				Desc:"好的不得了",
			},
			{
				Title:"CSS入门",
				Content:"CSS******",
				Desc:"好的不得了2",
			},
		},
	}

	db.Create(&author)
}
