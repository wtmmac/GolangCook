package main

import (
	"fmt"
	"time"
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
	for i := 0; i < 10000; i++ {
		time.Sleep(5 * time.Second)
		fmt.Println("sleep")
	}
	fmt.Println("main() => Hello World!")
}
