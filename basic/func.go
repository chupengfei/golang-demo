package main

import (
	"fmt"
	"reflect"
	"runtime"
)

func main() {

	q, r := div1(11, 2) // 函数内部定义好的返回值名称
	fmt.Println(q,r)

	a, b := div(11, 2)
	fmt.Println(a,b)

	q2, r2 := div2(11, 2) // 函数内部定义好的返回值名称
	fmt.Println(q2, r2)

	result, err := apply(func(a int, b int) int {
		return a / b
	}, 0, 3)
	fmt.Println(result, err)

	fmt.Println(sum(2,2,3,4))
	a, b = 1, 2
	swap(&a,&b)
	fmt.Println(a, b)
}

/**
  可以有多返回值
*/
func div(a, b int) (int, int) {
	return a / b, a % b
}

/**
  可以有多返回值,且给返回值取名字  建议方法
*/
func div1(a, b int) (q, r int) {
	return a / b, a % b
}

/**
  可以有多返回值,且给返回值取名字，q r分别赋值，直接return   容易混乱 不建议
*/
func div2(a, b int) (q2, r2 int) {
	q2 = a / b
	r2 = a % b
	return
}
/**
  入参可以是函数
 */
func apply(op func(int, int) int, a, b int) (result int, nil error) {
	p := reflect.ValueOf(op).Pointer()
	name := runtime.FuncForPC(p).Name()
	fmt.Printf("Calling function %s with args (%d, %d)", name, a, b)
	return op(a, b), nil
}
/**
 可变参数
 */
func sum(a int, numbers ...int) (sum int) {
	s := 0
	// 遍历可变参数
	for i := range numbers {
		s += numbers[i]
	}
	return s + a
}

/**
 指针
 交换外部传入的a b变量
 */
func swap(a, b *int)  {
	*a, *b = *b, *a
}

