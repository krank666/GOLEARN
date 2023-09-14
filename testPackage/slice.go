package testPackage

import "fmt"

func Slice() {
	a := [3]int{0, 1, 2}
	// a[2] 得到的只是数组元素，不是切片
	cl := a[1:] // 这才是数组切片
	// 表示从a的第n位开始切到末尾

	fmt.Println(cl) // 打印结果为[2]

	cl[0] = 5      // 改变切片中第0位的值
	fmt.Println(a) // 原数组被改变，打印结果为[0,5,2]

	// cl[ : ] 全取值
	// cl[0:2] 前闭后开原则，只能取到0，1两位数
	// cl[ :2] 只能取到0，1两位数
	// cl[1: ] 从指定下标取到最后

}

func SliceFunc() {
	arr := [3]int{1, 2, 5}
	sl := arr[1:]
	fmt.Println(sl) // [2,5]
	/*sl[2] = 6
	fmt.Println(sl) // index out of range数组越界，因为sl中只有2位
	*/
	sl = append(sl, 999) // 向sl末位增加一个元素, append不会改变原数组
	fmt.Println(sl)      // [2 5 999]
	fmt.Println(arr)     // [1 2 5]

	// 总结： 切片不是数组，切片属于数组
}
