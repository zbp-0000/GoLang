package logs_ope

import (
	"github.com/sirupsen/logrus"
	"os"
)

//初始化记录器一个实例：
var Logrus = logrus.New()

func init(){
	//先读取日志的配置文件：
	log_conf := LoadLogConfig()

	//设置日志的输出文件：
	file,err := os.OpenFile(log_conf.LogDir,os.O_APPEND|os.O_CREATE,0666)

	if err != nil {
		panic(err)
	}

	//将上面打开的file文件设置为  日志的输出文件：
	Logrus.Out = file


	//设置日志的级别：
	//定义一个map，专门存储日志级别：
	log_level_map := map[string]logrus.Level{
		"trace" : logrus.TraceLevel,
		"panic": logrus.PanicLevel,
		"fatal": logrus.FatalLevel,
		"error": logrus.ErrorLevel,
		"warn": logrus.WarnLevel,
		"info": logrus.InfoLevel,
		"debug": logrus.DebugLevel,
	}
	Logrus.SetLevel(log_level_map[log_conf.LogLevel])

	//日志格式化：设置文本格式
	Logrus.SetFormatter(&logrus.TextFormatter{})
}
