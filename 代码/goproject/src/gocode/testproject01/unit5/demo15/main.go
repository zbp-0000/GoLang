package main
import (
	"fmt"
	"strings"
)


func main(){
	//字符串的替换：
	str1 := strings.Replace("goandjavagogo", "go", "golang", -1) 
	str2 := strings.Replace("goandjavagogo", "go", "golang", 2) 
	fmt.Println(str1)
	fmt.Println(str2)

	//按照指定的某个字符，为分割标识，将一个学符串拆分成字符串数组:
	arr := strings.Split("go-python-java", "-")
	fmt.Println(arr)

	//将字符串的字母进行大小写的转换:
    fmt.Println(strings.ToLower("Go"))
	fmt.Println(strings.ToUpper("go"))

	//将字符串左右两边的空格去掉:
	fmt.Println(strings.TrimSpace("     go and java    "))

	//将字符串左右两边指定的字符去掉:
	fmt.Println(strings.Trim("~golang~", "~")) 

	//判断字符串是否以指定的字符串开头: 
	fmt.Println(strings.HasPrefix("http://java.sun.com/jsp/jstl/fmt", "http"))
	
	//判断字符串是否以指定的字符串结束: 
    fmt.Println(strings.HasSuffix("demo.png", ".jpg"))
}

