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



	//关联更新：
	//先查询
	//Preload方式查询:
	var author demostruct.Author
	//查询a_id=1的数据放入author中，并关联查询到Article字段对应的数据
	db.Preload("Article").Find(&author,"a_id = ?",1)
	fmt.Println(author)//{1 张三 30 男 [{1 HTML入门 HTML****** 好的不得了 1} {2 CSS入门 CSS****** 好的不得了2 1}]}

	//再更新：
	//如果直接Update操作那么关联的文章的记录就会被全部更改
	//db.Model(&author.Article).Update("title","JS入门")
	//所以你要改动指定的记录必须加入限定条件：
	db.Model(&author.Article).Where("ar_id = ?",1).Update("title","JS入门")


}

