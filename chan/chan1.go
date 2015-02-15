package main

import (
	"fmt"
)

var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	c <- 2
}

func main() {
	go f()
	intOut := <-c
	print(a)
	fmt.Println("Hello World!")
	print(intOut)
}
