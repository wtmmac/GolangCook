package main

import (
	"fmt"
)

var c = make(chan bool)
var str1 string

func f() {
	str1 = "f() => hello, world"
	fmt.Println(str1)
	c <- true
}

func main() {
	go f()
	_ = <-c
	fmt.Println("main() => Hello World!")
}
