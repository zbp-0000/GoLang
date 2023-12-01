package main

//简单的几行代码，就可以实现web服务
import (
	"github.com/gin-gonic/gin" //程序所需要的导包
)

func Hello(context *gin.Context) { //gin.Context - 把请求和响应都封装到gin.Context上下文环境中了
	//对页面的渲染效果（多种），你要给浏览器响应什么效果
	context.String(200, "这是我的第一个Gin项目")
}
func main() {
	//Default()返回的是一个引擎Engine，它是框架非常重要的数据结构，是框架的入口。
	//引擎 - 框架核心发送机 - 默认服务器 - 整个web服务都是由它来驱动的
	//Default()底层调用了New()，相当于New()的升级，New()返回的是一个引擎，在此基础上多增加了中间件处理-engine.Use(Logger(), Recovery())
	r := gin.Default()
	//r := gin.New()
	//路由：通过访问"/"的GET请求走这一条处理逻辑，走对应的函数中的内容
	//"/" : 路由规则  函数：路由函数
	//路由请求方式：GET、POST、DELETE、PATCH、PUT、OPTIONS、HEAD、Any
	//函数：可以直接写匿名函数，还可以在外部定义函数使用
	//r.GET("/", func(context *gin.Context) {
	//	context.String(200, "这是我的第一个Gin项目")
	//})
	r.GET("/testdemo", Hello)

	//启动引擎 ，服务器启动
	//Run可以传入参数：host+port
	//中间拼接的冒号一定不要忘记
	//r.Run("192.168.199.217:9999")
	r.Run(":9999")
}
