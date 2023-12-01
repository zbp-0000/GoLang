package main

import(
	"fmt"
	"sync"
)
var wg sync.WaitGroup //只定义无需赋值
func main(){
	wg.Add(5)
	//启动五个协程
	for i := 1 ;i <= 5;i++ {
		go func(n int){
			defer wg.Done()
			fmt.Println(n)		
		}(i)
	}
	//主线程一直在阻塞，什么时候wg减为0了，就停止
	wg.Wait()
}