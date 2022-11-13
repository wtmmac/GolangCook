package main

import "fmt"

type A struct{}

func (*A) hello() {
	fmt.Println("hello, world")
}

func main() {
	var a *A
	var b interface{}
	//b = a
	fmt.Printf("a is nil? %v\n", a == nil)
	fmt.Printf("b is nil? %v\n", b == nil)
	a.hello()
}
