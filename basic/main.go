package main

import "fmt"

func main() {
	a := 10
	b := &a
	fmt.Printf("b=%v\n", b)
	fmt.Printf("type b: %T\n", b)
	c := 11

	b = &c
	fmt.Printf("b=%v\n", b)
	fmt.Printf("type b: %T\n", b)
	//d := "33"
	//b = &d //报错  Cannot use '&d' (type *string) as type *int
	fmt.Printf("b=%d\n", *b) // *取地址对应的值

	//var s *int  // a是一个int类型的指针 报错 b=11panic: runtime error: invalid memory address or nil pointer dereference
	//*s = 11
	// 正确的
	var s = new(int) // 得到一个int类型的指针 值类型的用new
	fmt.Printf("s=%v *s=%v\n", s, *s)  //s=0xc000134030 *s=0
	*s = 11
	fmt.Printf("s=%v *s=%v\n", s, *s)  // s=0xc000134030 *s=11

	arr := new([3]int)
	arr[0] = 1
	(*arr)[1]=2
	fmt.Printf("arr=%v *arr=%v\n", arr, *arr) //arr=&[1 2 0] *arr=[1 2 0]

	slice := new([]int)
	//(*slice)[0]=1 //报错 panic: runtime error: index out of range [0] with length 0
	// 使用make初始化
	fmt.Printf("*b = %v\n",*slice )
	*slice = make([]int, 5,100)
	(*slice)[0]=1
	(*slice)[2]=2
	(*slice)[3]=3
	(*slice)[4]=3
	(*slice)[1]=5
	fmt.Printf("*b = %v\n",*slice )

	d := make([]int,3)
	d[0] = 1
	d[1] = 1
	d[2] = 1
	fmt.Printf("d = %v &d= %p\n",d, d )  //d = [1 1 1] &d= 0xc0000181e0
	d = append(d, 11)
	fmt.Printf("d = %v &d= %p\n",d, d ) // d = [1 1 1 11] &d= 0xc00001a0f0
}
