package main
import "fmt"
func main(){
	//双重循环：
	lable:
	for i := 1; i <= 5; i++ {
		for j := 2; j <= 4; j++ {			
			if i == 2 && j == 2 {
				continue lable
			}
			fmt.Printf("i: %v, j: %v \n",i,j)
		}
	}
	fmt.Println("-----ok")
}