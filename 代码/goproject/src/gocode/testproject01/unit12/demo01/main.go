package main

import(
	"fmt"
	"strconv"
	"time"
)

func test(){
	for i := 1;i <= 1000;i++ {
		fmt.Println("hello golang + " + strconv.Itoa(i))
		//阻塞一秒：
		time.Sleep(time.Second)
	}
}

func main(){//主线程
	go test() //开启一个协程

	for i := 1;i <= 10;i++ {
		fmt.Println("hello msb + " + strconv.Itoa(i))
		//阻塞一秒：
		time.Sleep(time.Second)
	}
}