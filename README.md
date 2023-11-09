# go环境搭建

https://golang.google.cn/dl/

点击安装，一路下一步，安装路径 `d:go`

安装到d盘的go目录里面

然后在d盘创建go_project用来存放开发代码

go：安装go sdk

goProject：写项目代码

1. 安装go后，
2. 打开终端，以管理员身份，输入`go env -w GO111MODEULE=on` 这句代码表示开启 go model 模式，不从 GOPATH下面去找包
3. 在goProject里面打开终端，输入 `go mod init goProject`，初始化 go mod 要加上文件夹名
4. goProject里面建一个src文件夹，里面书写项目要代码



# 分支/函数/循环/defer

```go
package main

import (
	"fmt"
	// "goCode/testProJect01/myFoo"
)
// name := "zhang san" 这种声明方式，只能在局部作用域声明
var name string = "zhang"
var age int = 18

func main() {
	fmt.Println("Hello goLang")
	// 分支语句 if 和 swith
	// ifAndSwith()
	// 循环
	// forAndRane()
	// 函数的使用
	// foo()
	// defer关键词
	deferFoo(10, 20)
	// 引入其他包 函数
	// MyFoo.MyFunc()
}

// 自定义函数
func foo() {
	sum := cal(10, 20)
	fmt.Println("函数的调用-一个参", sum)
	sum1,result1 := cal2(10, 20)
	fmt.Println("函数的调用-多参", sum1, result1)
}
func cal (num1 int, num2 int) (int) { // 如果返回值类型就一个的话，那么()可以省略; 没有返回值，可以不写返回类型
	return num1 + num2
}
func cal2 (num1 int, num2 int) (int,int) {
	return num1,num2
}

// 分支语句
func ifAndSwith() {
	// -----if
	if age >= 18 {
		fmt.Println("您的年龄为", age)
	} else if age < 18 {
		fmt.Println("您的年龄为", age)
	} else {
		fmt.Println("您的年龄为", age)
	}
	// -----swith
	switch age {
		case 10 : 
			fmt.Println("switch您的年龄小于11岁")
		case 18 :
			fmt.Println("switch您的年龄是18岁")
	}
}

// 循环语句
func forAndRane() {
	// for
	for i := 0; i < 5; i++ {
		// fmt.Println("for循环 %i", i)
		fmt.Printf("for循环%v \n", i)
	}
	// for range
	for i, value := range name {
		fmt.Printf("for range循环，索引:%d，值:%c \n", i, value)
	}
}

// defer关键词
func deferFoo (num1 int, num2 int) {
	// 遇到defer关键字，不会立即执行defer语句，有点像异步的意思
	 defer fmt.Println("num1", num1)
	 defer fmt.Println("num2", num2)
	 fmt.Println("deferFoo", num1 + num2)
}
```





# 数据类型

```go
	// ***基本数据类型
	// 在保证程序正确运行下，尽量使用占用空间小的数据类型
	// 整数 int int8 int16 int32 int64 uint uint8 uint16
	var num int = 8
	// 浮点数 float32 float64
	var num2 float32 = 3.2
	// 字符型 没有单独的字符型，使用byte来保存单个字母字符 ：打印出来就是 数字 本质上是一个整数，可以参与运算
	var byte1 byte = 'a'
	var byte2 byte = '6'
	var byte3 byte = '{'
	// 布尔型 bool
	var isShow bool = false
	// 字符串 string 不能修改字符串下标：name[0] = "s"，可以直接整个赋值 name = "wang wu"
	var name string = "zhang san"
	fmt.Println(num,num2,byte1,byte2,byte3,isShow,name)

	// ***派生数据类型 / 复杂数据类型
	// 指针
	// 数组
	// 结构体
	// 管道
	// 函数
	// 切片
	// 接口
	// map
```



# 数组

数组是值类型，传递值的时候，直接传递值，所以更改数据时，不会同时改变

# 指针

1. & 操作法 `&` 用于获取变量的地址，即取址操作。它返回一个指向该变量的指针。 

   ```go
   x := 10
   ptr := &x  // 获取 x 的地址，并赋值给 ptr
   fmt.Println(ptr)  // 打印指针的值 打印：0xc00000a0a8
   ```

2. *操作符  * 用于[解引用](https://so.csdn.net/so/search?q=%E8%A7%A3%E5%BC%95%E7%94%A8&spm=1001.2101.3001.7020)指针，即获取指针指向的值。它应用于指针类型的变量，用于访问该指针指向的实际值。

   ```go
   x := 10
   ptr := &x  // 获取 x 的地址，并赋值给 ptr
   fmt.Println(*ptr)  // 解引用 ptr，获取 x 的值 打印：10
   ```

   

# 二维数组

> 二维数组

```go
var arr [4][6]int 
// 默认值 [[0,0,0,0,0,0], [0,0,0,0,0,0], [0,0,0,0,0,0], [0,0,0,0,0,0]]
  
var arr [4][6]int
arr[1][2] = 1
arr[2][1] = 2
arr[2][3] = 3
for i := 0; i < 4; i++ {
    for j := 0; j < len(arr[i]); j++ {
        fmt.Println(arr[i][j], " ")
    }
    fmt.Println("\n")
}

// 直接赋值
var arr2 [2][3]int = [2][3]int{{1,2,3}, {4,5,6}}
fmt.Println(arr2) //[[1 2 3] [4 5 6]]

// ...
var arr3 [2][3]int = [...][3]int{{1,2,3}, {4,5,6}}
// 
var arr4 = [2][3]int{{1,2,3}, {4,5,6}}
//
var arr5  = [...][3]int{{1,2,3}, {4,5,6}}
```

### 语法

```go
var 数组名 [大小][大小]类型
```

### 遍历

```go
var arr2 [2][3]int = [2][3]int{{1,2,3}, {4,5,6}}
// for
for i := 0; i < 4; i++ {
    for j := 0; j < len(arr[i]); j++ {
        fmt.Println(arr[i][j], " ")
    }
    fmt.Println("\n")
}
// for rang
for i, v := range arr {
    for j, v2 := range v {
        fmt.Println("arr[%v][%v]=%v \t", i, j, v2)
    }
}
```



# 切片

```go
	// 创建切片 方式一
	var arr [5]int = [...]int {1,2,3,4,5}
	var slice = arr[1:3] // 左闭右开 取arr数组的小标 1 ~ 3 的值，不包括下标3
	fmt.Println("arr", arr)
	fmt.Println("slice", slice)
	// 方式二 直接创建切片 
	// 使用内置函数 make ；
	// 语法 var 切片名 []type = make([], len, [cap])
	// 参数说明：type就是类型； len：大小； cap：指定切片容量，可选  
	var slice2 []int = make([]int, 4, 10)
	fmt.Println(slice2) // [0,0,0,0] 容量为10

	// 循环切片 for
	var arr2 [5]int = [...]int{10,20,30,40,50}
	slice3 := arr2[1:4] // 20 30 40
	for i := 0; i < len(slice3); i++ {
		fmt.Printf("slice[%v]=%v \n",i,slice3[i])
	}
	// 循环切片 for range
	for i,value := range slice3 {
		fmt.Printf("i=%v value=%v \n", i ,value)
	}
```

### 增加 append

```go
arr append(6)
```



# map

**引用类型** 是一个数据结构，又叫集合，关联数组，key-value的数据结构

**注意：声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用**

> 语法：var 变量名 map[keyType]valueType

`var a map[string]string`

```go
// 声明map变量
var a map[string]string
var a map[int]string
var a map[string]int
var a map[string]map[string]string

// make 给 map 分配数据空间
a = make(map[string]string, 2) // 相当于可以放2对 key-value 如：[xx:yy, oo:pp]
// 并且 key 不能重复
```

例：key不能重复，重复赋值，覆盖前面的key的值

```go
var a map[string]string
a = make(map[string]string)
a["no1"] = "宋江"
a["no2"] = "吴用"
a["no1"] = "武松"
fmt.Println(a) // [no2:吴用, no1:武松]
```

### 使用方式

**方式一：先声明，再赋值**

```go
var a map[string]string
a = make(map[string]string, 2) // 2个容量
```

**方式二：声明，就直接make**

```go
var a = make(map[string]string) // 没有声明容量
```

**方式三：声明，直接赋值**

```go
var a map[string]string = map[string]string{"no1": "成都"}
a["no2"] = "北京"
```

### 增删改查

1. **增加和更新**

   map["key"] = value 如果key还没有，就增加，如何key存在就修改

   ```go
   var a map[string]string
   a = make(map[string]string)
   a["no1"] = "张三" // 增加
   a["no1"] = "李四" // 修改
   ```

2. **删除 delete**

   delete(map, "key")，delete是一个内置函数，

   如何key存在，就删除该key-value，

   如果不存在，不删除，不报错

   ```go
   var a map[string]string
   a = make(map[string]string)
   a["no1"] = "宋江"
   a["no2"] = "吴用"
   
   delete(a, "no1")
   fmt.Println(a) // [no2:吴用]
   ```

   **删除所有的话，直接make一个新的空间**

   ```go
   a = make(map[string]string)
   ```

3. **查询**

   ```go
   var heroes map[string]string
   heroes = make(map[string]string, 10)
   heroes["no1"] = "张三" // 增加
   heroes["no1"] = "李四" // 修改
   val, findRes := heroes["no1"] // findRes 表示'有' || ‘没有’
   if findRes {
       fmt.Println("找到了val=",val)
   } else {
       fmt.Println("没有no1这个key")
   }
   ```

### 遍历

```go
cities := make(map[string]string)
cities["no1"] = "北京"
cities["no2"] = "天津"

for k,v := range cities {
    fmt.Println("k=%v, v=%v",k, v)
}
// k=no1 v=北京
// k=no2 v=天津
```

### map切片

相当于动态的map，可以动态添加

```go
	var mosters []map[string]string // 初始化切片
	mosters = make([]map[string]string, 2) // 准备放2个妖怪
	if mosters[0] == nil {
		mosters[0] = make(map[string]string, 2)
		mosters[0]["name"] = "牛魔王"
		mosters[0]["age"] = "500"
	}
	if mosters[1] == nil {
		mosters[1] = make(map[string]string, 2)
		mosters[1]["name"] = "玉兔精"
		mosters[1]["age"] = "400"
	}
	newMonster := map[string]string{
		"name":"新的妖怪",
		"age":"200",
	}
	mosters = append(mosters, newMonster)
	fmt.Println(mosters) //[ map[age:500 name:牛魔王] map[]]
```

### map排序

先通过内置函数`sort.Ints(keys)`拿到map的key值，再通过遍历key来获取value；

```go
import {
    "fmt"
    "sort"
}
...
map1 := make(map[int]int, 10)
map1[10] = 100
map1[1] = 13
map[4] = 56
map[8] = 90

var keys []int // 存放key的数组
for k, v := range map1 { // 遍历map1 拿到key
    keys = append(keys, k)
}
// 排序 sort.Ints() 内置函数
sort.Ints(keys) // 将 key 进行排序
fmt.Println(keys) // 1 4 8 10

for _, k := range keys { // 遍历 keys 这样就拿到key
    fmt.Printf("map1[%v]=%v \n", k, map1[k])
}
```

### 细节

1. map是引用类型

```go
// 修改map，会修改原来的map
map1 := make(map[int]int)
map1[1] = 90
map1[2] = 80
map1[10] = 1
map1[20] = 2
func modify(map1 map[int]int) {
    map1[10] = 900
}
fmt.Printf("map1", map1) // map[1:90 2:80 10:900 20:2]
```

2. map容量达到后，在向map增加元素，会自动扩容，并不会发生panic，能自动增长
3. map的value也经常用**struct**类型(结构体)，更适合管理复杂的数据



# 结构体/面向对象

**值类型**

定义一个Cat结构体，结构体里面的字段，可以是`数组`，`基本数据类型`，`引用类型`

如果结构体的字段类型是：**指针，slice，和map那么初始值都是nil**

如果需要使用这杨的字段，**需要先make，才能使用**

```go
type Cat struct { // Cat 表示外面的包能访问
    Name string // Name 表示其他的包可以访问
    age int // age 小写的a，表示私有，其他包不能访问
    Color string
    Hobby string
    slice []int // 切片(数组)
    map1  map[string]string // 动态切片
}
var cat1 Cat
cat1.Name = "小白"
cat1.Age = 18
cat1.Color = "白色"
cat1.Hobby = "吃"
cat1.slice = make([]int, 10)
cat1.slice[0] = 10
cat1.map1 = make(map[string]string)
cat1.map1["on1"] = "张三"
```

### 使用方式

```go
type Person struct {
    Name string
    Age int
}
// 方式一 
var p Person
p.Name = "小白"
// 方式二 推荐使用
p2 := Person{"张三", 18}

// 方式三 获取到指针
var p3 *Person = new(Person)
(*p3).Name = "李四"
(*p3).Age = 19
// 简化 p3.Age = 19

// 方式四
var p4 *Person = &Person{}
(*p4).Name = "李四"
// 简化 p4.Name = "李四"
```

### 细节/注意事项

1. 结构体的所有字段在内存是连续的（可以通过指针内存的加减获取结构体的值）

2. 结构体是用户单独定义的类型，和其他类型进行转换时需要有完全相同的字段(名字、个数和类型)

   ```go
   type A struct {
       Num int
   }
   type B struct {
       Num int
   }
   func main() {
       var a A
       var b B
       b = a // 这样会报错
       b = B(a) // 强制转换 这里正确吗？正确 因为：类型相同、个数相同、名字相同
   }
   ```

3. 结构体进行**type**重新定义（相当于取别名），Golang认为是新的数据类型，但是相互间可以强转

   ```go
   type Student struct {
       Name string
       Age int
   }
   type Stu struct // Golang认为是新的数据类型
   func main() {
       var stu1 Student
       var stu2 Stu
       stu2 = stu1 // 报错，可以这样修改 stu2 = Stu(stu1)
   }
   // 例：
   type integer int
   func main() {
       var i interger = 10
       var j int = 20
       j = i // 错误 不能这样赋值；可以强制转换 j = int(i)
   }
   ```

4. struct的每个字段上，可以写上一个**tag**，该tag可以通过反射机制获取，常见的使用场景的就是序列化和反序列化。

   ```go
   import "encoding/json"
   type Monster struct {
       Name string `json:name` // tag 的使用
       Age int `json:age`
       Skill string `json:skill`
   }
   func main() {
       var monster Monster
       monster.Name = "红孩儿"
       monster.Age = 18
       monster.Skill = "吐火"
       data, err := json.Marshal(monster)
       if err != nil {
           fmt.Println("json encoding err", err)
           return
       }
       fmt.Println("json后的数据", string(data))
   }
   ```

### 方法

自定义类型，都可以有方法，而不仅仅是srtuct

```go
type Person struct {
    Name int
}
// Person 这个结构体 有个test方法 
func (p Person) test() { // 这个函数 这能被 a 这种数据类型调用
    fmt.Println(p.Name) // 这个时候为空
}

func main() {
    var p Person
    p.test()
    p.Name = "张三" // 赋值 再调用
    p.test() // 打印 Name 为 张三
}
```

**举例 两数之和**

```go
type Person struct {
    Name int
}
func (p Person) getSum(num1 int, num2 int) int {
    return num1 + num2
}
var p Person
res := p.getSum(10, 20)
fmt.Println(res) // 30
```

为了提高效率，通常我们**方法和结构体的指针类型绑定**

```go
type Circle struct {
    radius int
}
func (c *Circle) area2() float64 { // 这里绑定的指针类型
    return 3.14 * (*c).radius * (*c).radius
    return 3.14 * c.radius * c.radius // 也可以简化为这种写法，一样的效果
}
var c Circle
c.radius = 5.0
res2 := (&c).area2() // 直接用指针调用方法，这样便是引用传递，修改的值，直接改变指针
res2 := c.area2() // 可以简化为这种写法，底层还是 (&c)
```

如果手动的给类型，挂载了一个 String 方法，就会自动调用

```go
type Student struct {
    Name string
}
func (stu *Student) String() int { // 这里绑定的指针类型
    str := fmt.Speintf("Name = %v", stu.Name)
    return str // 也可以简化为这种写法，一样的效果
}
stu := Student{Name: "张三"}
fmt.Println(&stu) // 打印 Name = 张三 因为手动实现了 String
```

### 工厂模式/构造函数

Golang的结构体**没有构造函数**，通常可以使用工厂模式来解决这个问题

使用工厂模式实现跨包创建结构体实例的案例：

```go
// package model
package model
type Student struct {
    Name string
    Score float64
}
// package main
package main
import (
	"fmt"
    "goCode/testProJect01/model" // 引包
)
func main() {
    var stu = model.Student{
        Name : "张三"
        Score: 78.9
    }
    fmt.Println(stu)
}
```

如果我要让 Student 变成小写 student 封装一下，并且外面的包还能访问

```go
// package model
package model
type student struct { // 小写的 student
    Name string
    Score float64
}
// 通过工厂模式来解决
func NewStudent(n string, s float64) *student { // 返回指针地址
    return &student{
        Name: n,
        Score: s
    }
}

// package main
package main
import (
	"fmt"
    "goCode/testProJect01/model" // 引包
)
func main() {
    //var stu = model.Student{
    //    Name : "张三"
    //    Score: 78.9
    //}
    var stu = model.NewStudent("张三", 18)
    fmt.Println(stu) // 这个时候，这个stu是个指针，可以用 *stu
    fmt.Println("name=", (*stu).Name) // 简写 stu.Name
}

```

**如果这个时候，结构体的 Score，变成score**

```go
// package model
package model
type student struct { // 小写的 student
    Name string
    score float64
}
// 通过工厂模式来解决
func NewStudent(n string, s float64) *student { // 返回指针地址
    return &student{
        Name: n,
        Score: s
    }
}
func (s *student) GetScore() float64 {
    return s.score
}

// package main
package main
import (
	"fmt"
    "goCode/testProJect01/model" // 引包
)
func main() {
    //var stu = model.Student{
    //    Name : "张三"
    //    Score: 78.9
    //}
    var stu = model.NewStudent("张三", 18)
    fmt.Println(stu) // 这个时候，这个stu是个指针，可以用 *stu
    fmt.Println("name=", (*stu).Name, "score=", stu.GetScore()) // 简写 stu.Name
}

```

### 封装

```go
// foo 包
package foo

import (
	"fmt"
)
func MyFoo () {
	fmt.Print("MyFoo函数")
}
type account struct {
	account string
	password string
	balance float64

}
func Newaccount (a string, p string, b float64) *account {
	return &account{
		account: a,
		password: p,
		balance: b,
	}
}
func (a *account) Setaccount(account string) {
	a.account = account
}
func (a *account) SetPassword(pas string) {
	a.password = pas
}
func (a *account) SetBalance(balance float64) {
	a.balance = balance
}
func (a *account) Getaccount() string {
	return a.account
}
func (a *account) GetPassword() string {
	return a.password
}
func (a *account) GetBalance() float64 {
	return a.balance
}
func main() {
	fmt.Print("我是foo")
}
```



```go
// main 入口
package main

import (
	"fmt"
	"goProject/testProJect01/foo"
)


func main() {
	foo.MyFoo()
	account := foo.Newaccount("农业银行:123", "666666", 999.9)
	// foo.SetAccount("123123")
	if account != nil {
		fmt.Println("创建成功")
	}
	fmt.Print(account) // &{农业银行:123 666666 999.9}
    fmt.Print(*account)// {农业银行:123 666666 999.9}
	fmt.Println("获取账户名称：",account.Getaccount()) // 获取账户名称： 农业银行:123
 }
```

### 继承 =>结构体嵌套

```go
package main

import "fmt"

// 封装
type Student struct {
	Name string
	Age int
	Score int // 成绩
}
func (stu *Student) ShowInfo() {
	fmt.Printf("学生名=%v 年龄=%v 成绩=%v \n", stu.Name,stu.Age,stu.Score)
}
func (stu *Student) SetScore(score int) {
	stu.Score = score
}
// 小学生 结构体
type Pupil struct {
	Student
}
func (p *Pupil) testing() {
	fmt.Println("小学生考试中")
}
// 大学生结构体
type Graduate struct {
	Student // 嵌入了 Student 匿名结构体；完整的理解：匿名字段+嵌套结构体，匿名字段的字段类型为结构体；例：stu Student stu 是名称 Student是类型，好比：Name string
}
func (p *Graduate) testing() {
	fmt.Println("大学生考试中")
}

func main() {
	// 当我们对结构体嵌入了匿名结构体，使用方法会发生改变
	// 实例化 小学生 Pupil ：var pupil1 Pupil
	pupil1 := &Pupil{}
	pupil1.Student.Name = "李四"
	pupil1.Student.Age = 16
	pupil1.Student.SetScore(60) // 设置成绩函数
	pupil1.Student.ShowInfo() // 打印学生信息 => 学生名=李四 年龄=16 成绩=60 

	// 实例化 大学生 Pupil ：var graduate1 Graduate
	graduate1 := &Pupil{}
	graduate1.Student.Name = "王五"
	graduate1.Student.Age = 20
	graduate1.Student.SetScore(99) // 设置成绩函数
	graduate1.Student.ShowInfo() // 打印学生信息 => 学生名=王五 年龄=20 成绩=99 
}
```





# 打印输出

### fmt.Printf

```go
	%v,原样输出	      %T，打印类型 	  %t,bool类型	  %s，字符串 	    %f，浮点
    %d，10进制的整数 	 %b，2进制的整数  %o，8进制      %x，%X，16进制    %x：0-9，a-f
    %X：0-9，A-F     %c，打印字符     %p，打印地址
fmt.Printf("slice[%v]=%v \n",i,slice3[i])
```

### fmt.Println

原样输出



# 字符串方法

```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	name := "zhang san"
	// len 系统函数 获取字符串长度
	fmt.Println(len(name))
	// 字符串遍历
	r := []rune(name)
	fmt.Println("字符串遍历", r)
	// 字符串转整数
	str,_ := strconv.Atoi("6887")
	fmt.Println("字符串转整数", str)
	// 整数转字符串
	str2 := strconv.Itoa(6887)
	fmt.Println("整数转字符串", str2)
	// 查找子串是否在指定的字符串中
	strings.Contains("javaandgolang", "go")
	fmt.Println(strings.Contains("javaandgolang", "go")) // 结果 true
	// 统计一个字符串有几个指定的字串：
	strings.Count("javaandgolang", "a")
	fmt.Println("统计",strings.Count("javaandgolang", "a"))
	// 不区分大小写的字符串比较：
	fmt.Println(strings.EqualFold("go", "Go"))
	// 区分大小写
	fmt.Println("go" == "Go") // false
	// 返回子串在字符串第一次出现的索引值，如果没有就返回-1
	strings.Index("javaandgolang", "a")
	fmt.Println(strings.Index("javaandgolang", "a"))
	// 字符串的替换
	n := -1
	strings.Replace("goandjavagogo", "go", "golang", n) // n可以指定你希望替换几个，n=-1表示全部替换，替换两个n就是2
	// 字符串分割为数组
	strings.Split("go-python-java", "-")
	// 字符串的字母进行大小写转换
	strings.ToLower("Go") // go
	strings.ToUpper("go") // Go
	// 去除左右两侧空格
	strings.TrimSpace("  go and java  ")
	// 去除左右两侧指定字符
	strings.TrimLeft("~golang~", "~") // 左侧
	strings.TrimRight("~golang~", "~") // 右侧
	// 判断是否以指定字符串开头
	strings.HasPrefix("http://java.sun.com", "http")

}
```



# 案例

## 二维数组

定义二维数组，用于保存三个班，每个班五名同学成绩，并求出每个班评价分、以及所有班平均分

```go
	var myArr = [3][5]int{{10,20,30,40,50}, {60, 70, 80, 90, 100}, {40, 50, 60, 70, 80}}
	var myArr2 int
	for _, v := range myArr {
		var myArr1 int
		for _, v2 := range v {
			myArr1 += v2
		}
		fmt.Println("每个班平均分：", myArr1 / 5)
		myArr2 += myArr1
	}
	fmt.Println("所有班平均分：", myArr2 / 3)
```

## map

1. 存放3个学生信息，每个学生有name和sex信息

`我的理解：相当于一个数组，里面3个对象，每个对象里面有 name 和 sex`

```go
studentMap := make(map[string]map[string]string)
studentMap["stu01"] = make(map[string]string, 2)
// 创建 stu01 的 map，所有还要make一下；2 代表里面有2个值，一个name，一个sex
studentMap["stu01"]["name"] = "张三"
studentMap["stu01"]["sex"] = "男"
```

2. 综合练习

   使用 map[string]map[string]string 的map类型

   key：表示用户名，是唯一的，不可以重复

   如何某个用户名存在，就将其密码修改为”888888“，如果不存在就增加这个用户信息，包括昵称nickname和pwd

   编写一个函数 modifyUser(users map[string]map[string]string, name string) 完成上述功能

   ```go
   map[string]map[string]string
   ```

## 结构体/方法

打印一个 10*8 的矩阵

```go
package main
import "fmt"
type MethodUtilsType struct {
	num1 int
	num2 int
}
func (m MethodUtilsType) MethodUtils() {
	for i:=0; i < m.num1; i++ {
		for j:=0; j < m.num2; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
func main() {
	var method MethodUtilsType
	method.num1 = 10
	method.num2 = 8
	method.MethodUtils()
}

```

