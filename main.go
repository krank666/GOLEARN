package main

import (
	"awesomeProject/testPackage" //包的别名.
	"fmt"
	// . "awesomeProject/testPackage" // 使用 . 不需要写任何别名
)

func main() {
	fmt.Println(testPackage.A) // 包别名使用
	fmt.Println(testPackage.B) // 包别名使用
	//fmt.Println(A)
}
