package main

import(
	"fmt"
	"time"
)

func main(){
	//定义一个int管道：
	intChan := make(chan int,1)
	go func(){
		time.Sleep(time.Second * 15)
		intChan<- 10
	}()
	//定义一个string管道：
	stringChan := make(chan string,1)
	go func(){
		time.Sleep(time.Second * 12)
		stringChan<- "msbgolang"
	}()

	//fmt.Println(<-intChan)//本身取数据就是阻塞的
	select{
		case v := <-intChan :
			fmt.Println("intChan:",v)
		case v := <-stringChan :
			fmt.Println("stringChan:",v)
		default:
			fmt.Println("防止select被阻塞")
	}
}