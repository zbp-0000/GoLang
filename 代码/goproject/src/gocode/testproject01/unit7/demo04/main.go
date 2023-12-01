package main
import "fmt"
func main(){
	//实现的功能：给出五个学生的成绩，求出成绩的总和，平均数：
	//给出五个学生的成绩：--->数组存储：
	//定义一个数组：
	var scores [5]int
	//将成绩存入数组：(循环 + 终端输入)
	for i := 0; i < len(scores);i++ {//i：数组的下标
		fmt.Printf("请录入第个%d学生的成绩",i + 1)
		fmt.Scanln(&scores[i])
	}
	//展示一下班级的每个学生的成绩：（数组进行遍历）
	//方式1：普通for循环：
	for i := 0; i < len(scores);i++ {
		fmt.Printf("第%d个学生的成绩为：%d\n",i+1,scores[i])
	}
	fmt.Println("-------------------------------")
	//方式2：for-range循环
	for key,value := range scores {
		fmt.Printf("第%d个学生的成绩为：%d\n",key + 1,value)
	}
	fmt.Println("-------------------------------")
	for _,value := range scores {
		fmt.Printf("学生的成绩为：%d\n",value)
	}
}
