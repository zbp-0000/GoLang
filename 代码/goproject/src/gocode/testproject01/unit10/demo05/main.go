package main
import "fmt"
//定义Person结构体
type Person struct{
	Name string
}


//给Person结构体绑定方法test
func (p *Person) test(){//参数名字随便起
	p.Name = "露露"
	fmt.Println(p.Name)
}


func main(){
	//创建结构体对象：
	var p Person 
	p.Name = "丽丽"
	fmt.Printf("p的地址为：%p",&p)
	p.test()
	fmt.Println(p.Name)
}

