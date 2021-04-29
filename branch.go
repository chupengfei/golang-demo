package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

func main() {

	//readFile()
	fmt.Println(string(switchTest(1, 2, "+")))
	forTest()
	fmt.Println(convertToBin(5))
}

func readFile() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
}

func switchTest(a, b int, op string) int {
	var result int
	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		panic("unsupported operator:" + op)
	}
	fmt.Println(string(result))
	return result
}

func forTest() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
}

func convertToBin(n int) string {
	result := ""
	for ; n > 0; n /= 2 {
		lsb := n % 2
		// int转string
		result = strconv.Itoa(lsb) + result
	}
	return result
}

func printFile()  {
	
}
