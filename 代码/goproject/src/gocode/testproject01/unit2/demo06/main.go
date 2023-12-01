package main
import "fmt"

func main(){
	//练习转义字符：
	//\n  换行
	fmt.Println("aaa\nbbb")
	//\b 退格
	fmt.Println("aaa\bbbb")
	//\r 光标回到本行的开头，后续输入就会替换原有的字符
	fmt.Println("aaaaa\rbbb")
	//\t 制表符
	fmt.Println("aaaaaaaaaaaaa")
	fmt.Println("aaaaa\tbbbbb")
	fmt.Println("aaaaaaaa\tbbbbb")

	//\"
	fmt.Println("\"Golang\"")
}

