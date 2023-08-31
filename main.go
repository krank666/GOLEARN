package main

import (
	"fmt"
	"strconv"
)

/*
	func main() {
		var num1 uint
		var num2 int
		num1 = 999
		num2 = -999
		fmt.Println(num1, num2)
	}
*/
/*func main() {
	var num3 float64
	num3 = 3.1415926
	fmt.Println(num3)
}*/
/*func main() {
	var str1 string
	str1 = "hello"
	fmt.Println(str1)
}
*/

/*func main() {
	var str1 string
	str1 = "我是一个粉刷匠"
	fmt.Printf("%T", str1)
}*/

// 数据类型转换
func main() {
	var string1 string
	string1 = "123"
	int1, _ := strconv.Atoi(string1) // string to int
	fmt.Printf("%T", int1)           // 查看类型
	fmt.Println(int1)                // 查看打印结果

	int2, _ := strconv.ParseInt(string1, 10, 8) //string to int64
	// strconv.ParseInt需要三个参数 1. 需要转化的string值 2. 转化后所需要的进制，2~32进制 3. 限制位 这个包只能输出int64的类型
	fmt.Printf("%T", int2)
}
