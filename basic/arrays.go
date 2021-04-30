package main

import "fmt"

func main() {

	var arr1 [5]int
	// 此种方式需要给初始化数组
	arr2 := [2]int{1, 2}
	// 数组长度可以任意 编译器来数几个int
	arr3 := [...]int{2, 3, 4, 5, 1}
	fmt.Println(arr1, arr2, arr3) // [0 0 0 0 0] [1 2] [2 3 4 5 1]

	var grid [2][3]int
	fmt.Println(grid) // [[0 0 0] [0 0 0]]

	for i := 0; i < len(arr3); i++ {
		fmt.Print(arr3[i])
	}
	fmt.Println()
	// i是下标 v是值
	for i, v := range arr3 {
		fmt.Println(i, v)
	}

	// 只想要值 v是值
	for _, v := range arr3 {
		fmt.Println(v)
	}

	// 数组是值类型  [10]int 和 [20]int 是不同的类型  go语言一般不实用数组，使用切片
}
