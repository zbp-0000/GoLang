package main
import "fmt"
func main(){
	//声明数组：
	var arr [3]int16
	//获取数组的长度：
	fmt.Println(len(arr))
	//打印数组：
	fmt.Println(arr)//[0 0 0]
	//证明arr中存储的是地址值：
	fmt.Printf("arr的地址为：%p",&arr)
	//第一个空间的地址：
	fmt.Printf("arr的地址为：%p",&arr[0])
	//第二个空间的地址：
	fmt.Printf("arr的地址为：%p",&arr[1])
	//第三个空间的地址：
	fmt.Printf("arr的地址为：%p",&arr[2])

	//赋值：
	arr[0] = 10
	arr[1] = 20
	arr[2] = 30
}