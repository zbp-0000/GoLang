package main

import(
	"fmt"
	"sync"
	"time"
)
var wg sync.WaitGroup //只定义无需赋值
//加入读写锁：
var lock sync.RWMutex


func read(){
	defer wg.Done()
	lock.RLock()//如果只是读数据，那么这个锁不产生影响，但是读写同时发生的时候，就会有影响
	fmt.Println("开始读取数据")
	time.Sleep(time.Second)
	fmt.Println("读取数据成功")
	lock.RUnlock()
}

func write(){
	defer wg.Done()
	lock.Lock()
	fmt.Println("开始修改数据")
	time.Sleep(time.Second * 10)
	fmt.Println("修改数据成功")
	lock.Unlock()
}

func main(){
	wg.Add(6)
	//启动协程 ---> 场合：读多写少
	for i := 0;i < 5;i++ {
		go read()
	}

	go write()

	wg.Wait()

}