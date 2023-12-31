package main

import(
	"fmt"
	"net" //所需的网络编程全部都在net包下
)

func process(conn net.Conn){
	//连接用完一定要关闭：
	defer conn.Close()

	for{
		//创建一个切片，准备：将读取的数据放入切片：
		buf := make([]byte,1024)

		//从conn连接中读取数据：
		n,err := conn.Read(buf)
		if err != nil{
			return
		}
		//将读取内容在服务器端输出：
		fmt.Println(string(buf[0:n]))
	}
}

func main(){
	//打印：
	fmt.Println("服务器端启动了。。")
	//进行监听：需要指定服务器端TCP协议，服务器端的IP+PORT
	listen,err := net.Listen("tcp","127.0.0.1:8888")
	if err != nil{//监听失败
		fmt.Println("监听失败，err:",err)
		return
	}

	//监听成功以后：
	//循环等待客户端的链接：
	for{
		conn,err2 := listen.Accept()
		if err2 != nil {//客户端的等待失败
			fmt.Println("客户端的等待失败,err2:",err2)
		}else{
			//连接成功：
			fmt.Printf("等待链接成功，con=%v ，接收到的客户端信息：%v \n",conn,conn.RemoteAddr().String())
		}

		//准备一个协程，协程处理客户端服务请求：
		go process(conn)//不同的客户端的请求，连接conn不一样的
	}
}
