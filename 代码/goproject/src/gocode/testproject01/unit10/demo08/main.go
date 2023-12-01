package main
import "fmt"


type Student struct{
	Name string
}

//定义方法：
func (s Student) test01(){
	fmt.Println(s.Name)
}

//定义函数：
func method01(s Student){
	fmt.Println(s.Name)
}

func main(){
	//调用函数：
	var s Student = Student{"丽丽"}
	method01(s)

	//方法调用：
	s.test01()
}
