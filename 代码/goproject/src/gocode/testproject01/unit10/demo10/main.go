package main
import "fmt"

type Student struct{
	Name string
	Age int
}


func main(){
	//方式1：按照顺序赋值操作
	var s1 Student = Student{"小李",19}
	fmt.Println(s1)

	//方式2：按照指定类型
	var s2 Student = Student{
		Name : "丽丽",
		Age : 20,
	}
	fmt.Println(s2)
	//方式3：想要返回结构体的指针类型
	var s3 *Student = &Student{"明明",26}
	fmt.Println(*s3)
	var s4 *Student = &Student{
		Name : "娜娜",
		Age : 29,
	}
	fmt.Println(*s4)
}