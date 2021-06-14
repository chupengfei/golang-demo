package main

import (
	"fmt"
	"runtime"
	"time"
)
// 协程  轻量级"线程"
// 非抢占式多任务处理，由协程主动交出控制权
// 编译器/解释器/虚拟机层面的多任务
// 多个协程可能在一个或多个线程上运行
func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for  {
				a[i]++
				runtime.Gosched()  // 手动交出执行权
				//fmt.Println("hello")
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}





