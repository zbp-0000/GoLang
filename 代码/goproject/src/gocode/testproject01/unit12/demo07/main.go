package main

import(
	"fmt"
)

func main(){
	//定义管道 、 声明管道
	var intChan chan int
	//通过make初始化：管道可以存放3个int类型的数据
	intChan = make(chan int,3)
	//在管道中存放数据：
	intChan<- 10
	intChan<- 20

	//关闭管道：
	close(intChan)

	//再次写入数据：--->报错
	//intChan<- 30

	//当管道关闭后，读取数据是可以的：
	num := <- intChan
	fmt.Println(num)

}