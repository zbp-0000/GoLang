package main

import(
	"fmt"
	"os"
)

func main(){
	//打开文件：
	file,err := os.Open("d:/Test.txt");

	if err != nil {//出错
		fmt.Println("文件打开出错，对应错误为：",err)
	}

	//没有出错，输出文件：
	fmt.Printf("文件=%v",file)
	//.........一系列操作

	//关闭文件：
	err2 := file.Close();
	if err2 != nil {
		fmt.Println("关闭失败")
	}
}

