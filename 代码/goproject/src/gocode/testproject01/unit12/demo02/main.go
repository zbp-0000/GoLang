package main

import(
	"fmt"
	"time"
)

func main(){
	//匿名函数+外部变量 = 闭包
	for i := 1;i <= 5;i++ {
		//启动一个协程
		//使用匿名函数，直接调用匿名函数
		go func(n int){
			fmt.Println(n)
		}(i)
	}

	time.Sleep(time.Second * 2)
}