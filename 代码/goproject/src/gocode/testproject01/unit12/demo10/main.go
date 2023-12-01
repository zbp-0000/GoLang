package main

import(
	"fmt"
)

func main(){
	//默认情况下，管道是双向的--》可读可写：
	//var intChan1 chan int
	//声明为只写：
	var intChan2 chan<- int  // 管道具备<- 只写性质
	intChan2 = make(chan int,3)
	intChan2<- 20
	//num := <-intChan2 报错
	fmt.Println("intChan2:",intChan2)
	//声明为只读：
	var intChan3 <-chan int
	if intChan3 != nil {
		num1 := <-intChan3
		fmt.Println("num1:",num1)
	}

	//intChan3<- 30 报错
	
}