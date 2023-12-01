package main

import(
	"fmt"
	"io/ioutil"
)

func main(){
	//备注：在下面的程序中不需要进行 Open\Close操作，因为文件的打开和关闭操作被封装在ReadFile函数内部了
	//读取文件：
	content,err := ioutil.ReadFile("d:/Test.txt")//返回内容为：[]byte,err

	if err != nil {//读取有误
		fmt.Println("读取出错，错误为：",err)
	}

	//如果读取成功，将内容显示在终端即可：
	//fmt.Printf("%v",content)
	fmt.Printf("%v",string(content))
}