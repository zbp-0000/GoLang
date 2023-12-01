package main

import(
	"fmt"
	"reflect"
)

//利用一个函数，函数的参数定义为空接口：
func testReflect(i interface{}){//空接口没有任何方法,所以可以理解为所有类型都实现了空接口，也可以理解为我们可以把任何一个变量赋给空接口。
	//1.调用TypeOf函数，返回reflect.Type类型数据：
	reType := reflect.TypeOf(i)

	//2.调用ValueOf函数，返回reflect.Value类型数据：
	reValue :=reflect.ValueOf(i)

	//获取变量的类别：
	//（1）reType.Kind()
	k1 := reType.Kind()
	fmt.Println(k1)

	//（2）reValue.Kind()
	k2 := reValue.Kind()
	fmt.Println(k2)

	//获取变量的类型：
	//reValue转成空接口：
	i2 := reValue.Interface()
	//类型断言：
	n,flag := i2.(Student)
	if flag == true {//断言成功
		fmt.Printf("结构体的类型是：%T",n)
	}

}

//定义学生结构体：
type Student struct{
	Name string
	Age int
}

func main(){
	//对结构体类型进行反射：
	//定义结构体具体的实例：
	stu := Student{
		Name : "丽丽",
		Age : 18,	
	}
	testReflect(stu)
}