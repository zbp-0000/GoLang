package logs_ope

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

//对应结构体：
type LogConfig struct {
	LogDir string `json:"log_dir"`
	LogLevel string  `json:"log_level"`
}


//读取配置文件：
func LoadLogConfig() *LogConfig{
	log_conf := LogConfig{}

	//打开文件：
	file,err := os.Open("part16/confs/log_config.json")

	if err != nil{//错误处理
		panic(err)
	}
	//资源释放：
	defer file.Close()

	//用流读取文件中内容：
	data,err2  := ioutil.ReadAll(file)

	if err2 != nil {
		panic(err2)
	}
	//Unmarshal将json字符串解码到对应的数据结构中：
	//第一个参数：json字符串，第二个参数：接收json解析的数据结构
	err3 := json.Unmarshal(data,&log_conf)

	if err3 != nil {
		panic(err3)
	}

	return &log_conf
}
