package main

import (
	"fmt"
	"math"
)
/**

go语言变量定义了必须使用 否则报错

 */

var aa = 3 //包内部的变量 没有全局变量的概念
var ss = "kkk"
// 不能用 aa := 3
// 可以用var()集中定义
var(
	bb = 3
	cc = 44
)

func main() {
	fmt.Println("nihao shijie")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	triangle()
	constant()
	fmt.Println(k)
}
/**
初值
*/
func variableZeroValue()  {
	var a int // 初值0
	var s string // 初值 ""
	fmt.Printf("%d %q\n",a,s)
}

func variableInitialValue()  {
	var a,b = 3, 4
	var s = "abc"
	fmt.Println(a,b,s)
}
/**
 可以类型推断
 */
func variableTypeDeduction()  {
	var a, b, c, d = 1, "a", true, 2
	fmt.Println(a,b,c,d)
}
/**
 上面的简单写法
 */
func variableShorter()  {
	 a, b, c, d := 1, "a", true, 2
	 b = "b"
	fmt.Println(a,b,c,d)
}
/**
 golang没有隐式类型转换，只有显示
 */
func triangle()  {
	var a,b = 3,4
	var c int
	c=int(math.Sqrt(float64(a*a+b*b)))
	fmt.Println(c)
}
/**
常量类型
 */
func constant()  {
	const k = "kk"
	fmt.Println(k)
	const (
		filename = "a.txt"
		a,b = 3,4
	)
	//k = "5" 不能赋值
	fmt.Println(filename)
}

const k = "2"

