package main

import(
	"fmt"
	"reflect"
)

//定义一个结构体：
type Student struct{
	Name string
	Age int
}

//给结构体绑定方法：
func (s Student) CPrint(){
	fmt.Println("调用了Print()方法")
	fmt.Println("学生的名字是：",s.Name)
}

func (s Student) AGetSum(n1,n2 int) int{
	fmt.Println("调用了AGetSum方法")
	return n1 + n2
}

func (s Student) BSet(name string,age int){
	s.Name = name
	s.Age = age
}

//定义函数操作结构体进行反射操作：
func TestStudentStruct(a interface{}){
	//a转成reflect.Value类型：
	val := reflect.ValueOf(a)
	fmt.Println(val)

	//通过reflect.Value类型操作结构体内部的字段：
	n1 := val.NumField()
	fmt.Println(n1)
	//遍历-获取具体的字段：
	for i := 0; i < n1;i++{
		fmt.Printf("第%d个字段的值是：%v",i,val.Field(i))
	}
	fmt.Println()

	//通过reflect.Value类型操作结构体内部的方法：
	n2 := val.NumMethod()
	fmt.Println(n2)

	//调用CPrint()方法：
	//调用方法，方法的首字母必须大写才能有对应的反射的访问权限
	//方法的顺序按照ASCII的顺序排列的，a,b,c,,,,,,索引：0,1,2，，，，，
	val.Method(2).Call(nil)

	//调用AGetSum方法：
	//定义Value的切片：
	var params []reflect.Value
	params = append(params,reflect.ValueOf(10))
	params = append(params,reflect.ValueOf(20))
	result := val.Method(0).Call(params)
	fmt.Println("AGetSum方法的返回值为：",result[0].Int())
}


func main(){
	//定义结构体具体的实例：
	s := Student{
		Name : "丽丽",
		Age : 18,
	}

	//调用TestStudentStruct：
	TestStudentStruct(s)
}


