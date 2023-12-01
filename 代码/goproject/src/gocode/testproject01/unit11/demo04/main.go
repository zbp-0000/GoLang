package main

import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	//写入文件操作：
	//打开文件：
	file , err := os.OpenFile("d:/Demo.txt",os.O_RDWR | os.O_APPEND | os.O_CREATE,0666)

	if err != nil {//文件打开失败
		fmt.Printf("打开文件失败",err)
		return
	}

	//及时将文件关闭：
	defer file.Close()

	//写入文件操作：---》IO流---》缓冲输出流(带缓冲区)
	writer := bufio.NewWriter(file)
	for i := 0; i < 10;i++ {
		writer.WriteString("你好 马士兵\n")
	} 

	//流带缓冲区，刷新数据--->真正写入文件中：
	writer.Flush()

	s :=os.FileMode(0666).String()
	fmt.Println(s)
}