package main

import (
	"fmt"
)

type A struct{}

func (*A) hello() {
	fmt.Println("hello, world")
}

type People struct {
	Name string
	Age  int
}

func main() {
	var a *A
	var b interface{}
	// b = a
	fmt.Printf("a is nil? %v\n", a == nil)
	fmt.Printf("b is nil? %v\n", b == nil)
	a.hello()

	p := People{
		Name: "name",
		Age:  10,
	}
	fmt.Println(p)
}
