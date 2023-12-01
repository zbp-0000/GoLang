package main
import "fmt"

//定义老师结构体，将老师中的各个属性  统一放入结构体中管理：
type Teacher struct{
	//变量名字大写外界可以访问这个属性
	Name string
	Age int
	School string
}
func main(){
	//创建老师结构体的实例、对象、变量：
	var t *Teacher = &Teacher{"马士兵",46,"清华大学"}
	// (*t).Name = "马士兵"
	// (*t).Age = 45
	// t.School = "清华大学"
	fmt.Println(*t)
}


// func main(){
// 	//创建老师结构体的实例、对象、变量：
// 	var t *Teacher = new(Teacher) 
// 	//t是指针，t其实指向的就是地址，应该给这个地址的指向的对象的字段赋值：
// 	(*t).Name = "马士兵"
// 	(*t).Age = 45  //*的作用：根据地址取值
// 	//为了符合程序员的编程习惯，go提供了简化的赋值方式：
// 	t.School = "清华大学" //go编译器底层对t.School转化 (*t).School = "清华大学"
// 	fmt.Println(*t)
// }




// func main(){
// 	//创建老师结构体的实例、对象、变量：
// 	var t Teacher = Teacher{"赵珊珊",31,"黑龙江大学"}
// 	fmt.Println(t)
// 	// t.Name = "赵珊珊"
// 	// t.Age = 31
// 	// t.School = "黑龙江大学"
// 	fmt.Println(t)
// }

// func main(){
// 	//创建老师结构体的实例、对象、变量：
// 	var t1 Teacher // var a int
// 	fmt.Println(t1) //在未赋值时默认值：{ 0 }
// 	t1.Name = "马士兵"
// 	t1.Age = 45
// 	t1.School = "清华大学"
// 	fmt.Println(t1)
// 	fmt.Println(t1.Age + 10)
// }

