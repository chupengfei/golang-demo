package main

import "fmt"
// slice不是值类型 视图 数组变成slice  arr[:]
func main() {
	arr := [...]int{0,1,2,3,4,5,6,7}
	fmt.Printf("%T\n",arr) // 数组
	s := arr[2:6]
	fmt.Printf("%T\n",s)  // 切片
	fmt.Println(s) // [2 3 4 5]
	sliceTest(s)
	fmt.Println(arr) // [0 1 100 3 4 5 6 7]
	s2 := s[1:]
	s2[2] = 11
	fmt.Println(s2) // [3 4 5]
	fmt.Println(s) // [100 3 4 11]
	slice2()
}

func sliceTest(s []int)  {
	s[0] = 100
}

func slice2()  {
	s := make([]int, 16)
	s[2] = 2
	fmt.Println(len(s))
	s1 := append(s, 1)
	fmt.Println(len(s1))
	s2 := append(s1, 2)
	fmt.Println(len(s2))
	fmt.Println(s2)


}


