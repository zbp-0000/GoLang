package main
import "fmt"
func main(){
	//定义map
	b := make(map[int]string)
	//增加：
	b[20095452] = "张三"
	b[20095387] = "李四"
	//修改：
	b[20095452] = "王五"
	//删除：
	delete(b,20095387)
	delete(b,20089546)
	fmt.Println(b)
	//查找：
	value,flag := b[200]
	fmt.Println(value)
	fmt.Println(flag)
}