package testPackage

import "fmt"

func Array() {
	a := [3]int{0, 1, 2}                       // 指定长度
	b := [...]any{1231231, "sss", 5656, "qqq"} // 不指定长度及内容

	var c = new([10]int)
	c[0] = 9 // 除了指定元素的值其他都是默认0
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	zoo := [...]string{"monkey", "cheval", "chine"}
	//  第一种常规循环
	for i := 0; i < len(zoo); i++ {
		fmt.Println(zoo[i] + " run")
	}

	// i=index 数组元素的下标  v=value 数组元素的值
	for i, v := range zoo {
		fmt.Println(i, v)
	}

	fmt.Println(len(zoo), cap(zoo)) //len是数组长度，cap是空间
	// 当到达一定临界值时， cap会自动扩容

	// 多维数组 以二维数组举例
	twoArr := [3][2]int{ // [3]表示数组内包含3个元素 [2]表示元素内位数
		{1, 3},
		{0, 1},
		{-1, 0},
	}
	fmt.Println(twoArr)
}
