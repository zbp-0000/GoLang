package main
import (
	"fmt"
)


//定义动物结构体：
type Animal struct{
	Age int
	weight float32
}

//给Animal绑定方法：喊叫：
func (an *Animal) Shout(){
	fmt.Println("我可以大声喊叫")
}

////给Animal绑定方法：自我展示：
func (an *Animal) showInfo(){
	fmt.Printf("动物的年龄是：%v,动物的体重是：%v",an.Age,an.weight)
}

//定义结构体：Cat
type Cat struct{
	//为了复用性，体现继承思维，嵌入匿名结构体：——》将Animal中的字段和方法都达到复用
	Animal
	Age int
}

func (c *Cat) showInfo(){
	fmt.Printf("~~~~~~~~动物的年龄是：%v,动物的体重是：%v",c.Age,c.weight)
}

//对Cat绑定特有的方法：
func (c *Cat) scratch(){
	fmt.Println("我是小猫，我可以挠人")
}


func main(){
	//创建Cat结构体示例：
	// cat := &Cat{}
	// cat.Age = 3
	// cat.weight = 10.6
	// cat.Shout()
	// cat.showInfo()
	// cat.scratch()

	cat := &Cat{}
	cat.weight = 9.4
	cat.Age = 10 //就近原则
	cat.Animal.Age = 20
	cat.showInfo()//就近原则
	cat.Animal.showInfo()
}

