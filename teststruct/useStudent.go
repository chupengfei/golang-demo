package main

import "fmt"

func main() {
	ss := Student{
		"fei", 11,
	}

	ss.useStudent()
}

func (s Student) useStudent() {
	fmt.Printf("student %v", s)
}
