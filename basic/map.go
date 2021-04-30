package main

import "fmt"

func main() {
	map1()
	map2() // empty map
	map3() // nil
}

func map1() {
	m := map[string]string{
		"name":   "cc",
		"course": "golang",
	}
	// 遍历map
	for k, v := range m {
		fmt.Println(k, v)
	}
	// 获取key为name的map的value值 ，返回值第一个参数为value，第二个为是否存在 存在为true，否则为false
	name, ok := m["name"]
	if course, ok := m["course"]; ok {
		fmt.Printf("11 %s\n",course)
	}else {
		fmt.Println(course)
	}

	fmt.Println(name, ok)
	fmt.Println(m)
}

func map2() {
	m := make(map[string]string)

	fmt.Println(m)
}

func map3() {
	var m3 map[string]string

	fmt.Println(m3)
}
