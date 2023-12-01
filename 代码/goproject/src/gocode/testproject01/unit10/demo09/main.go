package main
import "fmt"


type Student struct{
	Name string
}

//定义方法：
func (s Student) test01(){
	fmt.Println(s.Name)
}

func (s *Student) test02(){
	fmt.Println((*s).Name)
}
func main(){
	var s Student = Student{"丽丽"}
	s.test01()
	(&s).test01()//虽然用指针类型调用，但是传递还是按照值传递的形式

	(&s).test02()
	s.test02()//等价
}







//定义函数：
func method01(s Student){
	fmt.Println(s.Name)
}


func method02(s *Student){
	fmt.Println((*s).Name)
}



// func main(){
// 	var s Student = Student{"丽丽"}
// 	method01(s)
// 	//method01(&s)错误
// 	method02(&s)
// 	//method02(s)错误
// }


