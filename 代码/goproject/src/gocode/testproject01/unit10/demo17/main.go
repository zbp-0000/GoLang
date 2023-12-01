package main

import "fmt"

type CInterface interface{
	c()
}

type BInterface interface{
	b()
}

type AInterface interface{
	BInterface
	CInterface
	a()
}

type Stu struct{

}

func (s Stu) a(){
	fmt.Println("a")
}
func (s Stu) b(){
	fmt.Println("b")
}
func (s Stu) c(){
	fmt.Println("c")
}

type E interface{

}

func main(){
	var s Stu
	var a AInterface = s
	a.a()
	a.b()
	a.c()

	var e E = s
	fmt.Println(e)
	var e2 interface{} = s
	fmt.Println(e2)
	var num float64 = 9.3
	var e3 interface{} = num
	fmt.Println(e3)
}
