package main

import "fmt"

//接口的定义：定义规则、定义规范，定义某种能力：
type SayHello interface{
	//声明没有实现的方法：
	sayHello()
}


//接口的实现：定义一个结构体：
//中国人：
type Chinese struct{
	name string
}

//实现接口的方法---》具体的实现：
func (person Chinese) sayHello(){
	fmt.Println("你好")
}
//中国人特有的方法
func (person Chinese) niuYangGe(){
	fmt.Println("东北文化-扭秧歌")
}



//接口的实现：定义一个结构体：
//美国人：
type American struct{
	name string
}
//实现接口的方法---》具体的实现：
func (person American) sayHello(){
	fmt.Println("hi")
}

func (person American) disco(){
	fmt.Println("野狼disco")
}


//定义一个函数：专门用来各国人打招呼的函数，接收具备SayHello接口的能力的变量：
func greet(s SayHello){
	s.sayHello()
	//断言：
	// ch,flag := s.(Chinese)//看s是否能转成Chinese类型并且赋给ch变量,flag是判断是否转成功
	// if flag == true{
	// 	ch.niuYangGe()
	// }else{
	// 	fmt.Println("美国人不会扭秧歌")
	// }

	
	// if ch,flag := s.(Chinese);flag{
	// 	ch.niuYangGe()
	// }else{
	// 	fmt.Println("美国人不会扭秧歌")
	// }

	switch s.(type){//type属于go中的一个关键字，固定写法
		case Chinese:
			ch := s.(Chinese)
			ch.niuYangGe()
		case American:
			am := s.(American)
			am.disco()
	}

	fmt.Println("打招呼。。。")
}


func main(){
	
	//创建一个中国人：
	c := Chinese{}
	//创建一个美国人：
	//a := American{}

	//美国人打招呼：
	//greet(a)
	//中国人打招呼：
	greet(c)

}
