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



	//关联查询：
	//Association方式查询：因为关联关系在Author中，所以我们操作的是Author模型
	//var author demostruct.Author
	//如果只是执行下面这步操作，那么关联的Article信息是查询不到的：
	//db.First(&author,"a_id = ?",1)
	//fmt.Println(author)//{1 张三 30 男 []}

	//如果想要查询到Article相关内容，必须执行如下操作：
	//Model参数：要查询的表数据，Association参数：关联到的具体的模型：模型名字Article（字段名字）
	//Find参数：查询的数据要放在什么字段中&author.Article
	//db.Model(&author).Association("Article").Find(&author.Article)
	//fmt.Println(author)//{1 张三 30 男 [{1 HTML入门 HTML****** 好的不得了 1} {2 CSS入门 CSS****** 好的不得了2 1}]}



	//Preload方式查询:
	//var author demostruct.Author
	//查询a_id=1的数据放入author中，并关联查询到Article字段对应的数据
	//db.Preload("Article").Find(&author,"a_id = ?",1)
	//fmt.Println(author)//{1 张三 30 男 [{1 HTML入门 HTML****** 好的不得了 1} {2 CSS入门 CSS****** 好的不得了2 1}]}



	//Related方式查询：
	var author demostruct.Author
	db.First(&author,"a_id = ?",1)
	fmt.Println(author)//{1 张三 30 男 []}

	var as []demostruct.Article
	//通过author模型查出来的Article字段的信息放入新的容器as中：
	db.Model(&author).Related(&as,"Article")
	fmt.Println(as)//[{1 HTML入门 HTML****** 好的不得了 1} {2 CSS入门 CSS****** 好的不得了2 1}]
	fmt.Println(author)//{1 张三 30 男 []}

}

