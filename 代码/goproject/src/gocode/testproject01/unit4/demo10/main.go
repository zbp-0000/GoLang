package main
import "fmt"
func main(){
	//双重循环：
	lable2:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {
			fmt.Printf("i: %v, j: %v \n",i,j)
			if i == 2 && j == 2 {
				break lable2
			}
		}
	}

	fmt.Println("-----ok")
}