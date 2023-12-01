package main
import "fmt"

func main(){
	//定义字符类型的数据：
	var c1 byte = 'a'
	fmt.Println(c1)//97
	var c2 byte = '6'
	fmt.Println(c2)//54
	var c3 byte = '('
	fmt.Println(c3 + 20)//40
	//字符类型，本质上就是一个整数，也可以直接参与运算，输出字符的时候，会将对应的码值做一个输出
	//字母，数字，标点等字符，底层是按照ASCII进行存储。

	var c4 int = '中'
	fmt.Println(c4)
	//汉字字符，底层对应的是Unicode码值
	//对应的码值为20013，byte类型溢出，能存储的范围：可以用int
	//总结：Golang的字符对应的使用的是UTF-8编码（Unicode是对应的字符集，UTF-8是Unicode的其中的一种编码方案）

	var c5 byte = 'A'
	//想显示对应的字符，必须采用格式化输出
	fmt.Printf("c5对应的具体的字符为：%c",c5)

}