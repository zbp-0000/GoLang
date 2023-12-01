package main

import(
	"fmt"
	"reflect"
)

//利用一个函数，函数的参数定义为空接口：
func testReflect(i interface{}){//空接口没有任何方法,所以可以理解为所有类型都实现了空接口，也可以理解为我们可以把任何一个变量赋给空接口。
	reValue :=reflect.ValueOf(i)

	//通过SetInt()来改变值：
	reValue.Elem().SetInt(40)
	
}

func main(){
	//对基本数据类型进行反射：
	//定义一个基本数据类型：
	var num int = 100
	testReflect(&num) //传入指针地址

	fmt.Println(num)
}