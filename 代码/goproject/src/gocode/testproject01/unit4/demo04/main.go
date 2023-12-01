package main
import "fmt"
func main(){
	//实现功能：根据给出的学生分数，判断学生的等级：
	// >=90  -----A
	// >=80  -----B
	// >=70  -----C
	// >=60  -----D
	// <60   -----E

	//给出一个学生分数：
	//var score int = 187
	//根据分数判断等级：
	//switch后面是一个表达式，这个表达式的结果依次跟case进行比较，满足结果的话就执行冒号后面的代码。
	//default是用来“兜底”的一个分支，其它case分支都不走的情况下就会走default分支
	//default分支可以放在任意位置上，不一定非要放在最后。
	var a int32 = 7
	//var b int64 = 9
	switch a {
		case 10 , 11 , 12 :
			fmt.Println("您的等级为A级")
		case 9 :
			fmt.Println("您的等级为A级")
		case 8 :
			fmt.Println("您的等级为B级")
		case 7 :
			fmt.Println("您的等级为C级")
			fallthrough
		case 6 :
			fmt.Println("您的等级为D级")
		case 5 :
			fmt.Println("您的等级为E级")
		case 4 :
			fmt.Println("您的等级为E级")
		case 3 :
			fmt.Println("您的等级为E级")
		case 2 :
			fmt.Println("您的等级为E级")
		case 1 :
			fmt.Println("您的等级为E级")
		case 0 :
			fmt.Println("您的等级为E级")
		default:
			fmt.Println("您的成绩有误")
	}

	// switch {
	// 	case a == 1 :
	// 		fmt.Println("aaa")
	// 	case a == 2 :
	// 		fmt.Println("bbb")
	// }

	// switch b := 7 ; {
	// 	case b > 6 :
	// 		fmt.Println("大于6")
	// 	case b <= 6 :
	// 		fmt.Println("不大于6")
	// }
	
}