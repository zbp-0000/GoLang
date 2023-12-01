package main

import(
	"fmt"
	"io/ioutil"
)

func main(){
	//定义源文件：
	file1Path := "d:/Demo.txt"
	//定义目标文件：
	file2Path := "d:/Demo2.txt"

	//对文件进行读取：
	content,err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Println("读取有问题!")
		return
	}

	//写出文件：
	err = ioutil.WriteFile(file2Path,content,0666)
	if err != nil {
		fmt.Println("写出失败！")
	}
}