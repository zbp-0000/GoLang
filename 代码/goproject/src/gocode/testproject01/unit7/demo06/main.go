package main
import "fmt"
func main(){
	var arr3 = [3]int{3,6,9}
	test1(&arr3)//传入arr3数组的地址
	fmt.Println(arr3)//[3 6 9]
}

func test1(arr *[3]int){
	(*arr)[0] = 7
}


	//定义一个数组：
	// var arr1 = [3]int{3,6,9}
	// fmt.Printf("数组的类型为：%T",arr1)

	// var arr2 = [6]int{3,6,9,1,4,7}
	// fmt.Printf("数组的类型为：%T",arr2)


	// func main(){
	// 	var arr3 = [3]int{3,6,9}
	// 	test1(arr3)
	// 	fmt.Println(arr3)//[3 6 9]
	// }
	
	// func test1(arr [3]int){
	// 	arr[0] = 7
	// }