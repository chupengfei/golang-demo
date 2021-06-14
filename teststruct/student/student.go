package student

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (s Student) UseStudent() {
	fmt.Printf("student %v", s)
}

