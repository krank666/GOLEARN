package testPackage

import "fmt"

func Process() {
	/*	// 递增语句
		a := 1
		a++
		fmt.Println(a)*/

	/*	// if语句
		a := 1
		if a > 0 {
			fmt.Println("大于")
		} else {
			fmt.Println("小于")
		}*/

	/*	// switch判断语句
		a := 1
		switch a {
		case 0:
			fmt.Println("0")
		case 1:
			fmt.Println("1")
			fallthrough
		case 2:
			fmt.Println("2")
		default:
			fmt.Println("none")
		}*/

	/*// 循环语句
	a := 0
	for a < 10 {
		a++
		fmt.Println(a)
	}
	for a := 0; a < 10; a++ {

	}*/

	// 跳转语句
	// goto 可以定义循环体
	a := 0
A:
	for a < 10 {
		a++
		fmt.Println(a)
		if a == 5 {
			break A
			goto B
		}
	}
B:
	fmt.Println("rush B!")

}
