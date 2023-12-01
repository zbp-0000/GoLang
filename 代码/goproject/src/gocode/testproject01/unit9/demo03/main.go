package main
import "fmt"
func main(){
	//定义map
	b := make(map[int]string)
	//增加：
	b[20095452] = "张三"
	b[20095387] = "李四"
	b[20098833] = "王五"

	//获取长度：
	fmt.Println(len(b))

	//遍历：
	for k,v := range b {
		fmt.Printf("key为：%v value为%v \t",k,v)
	}
	fmt.Println("---------------------------")
	//加深难度：
	a := make(map[string]map[int]string)
	//赋值：
	a["班级1"] = make(map[int]string,3)
	a["班级1"][20096677] = "露露"
	a["班级1"][20098833] = "丽丽"
	a["班级1"][20097722] = "菲菲"

	a["班级2"] = make(map[int]string,3)
	a["班级2"][20089911] = "小明"
	a["班级2"][20085533] = "小龙"
	a["班级2"][20087244] = "小飞"

	for k1,v1:= range a {
		fmt.Println(k1)
		for k2,v2:= range v1{
			fmt.Printf("学生学号为：%v 学生姓名为%v \t",k2,v2)
		}
		fmt.Println()
	}



}