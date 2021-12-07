package main

import (
	"fmt"
)

var c = make(chan bool)
var str1 string

func chan1() {
	str1 = "chan1() => hello, world"
	fmt.Println(str1)
	c <- true
}

func main() {
	go chan1()
	_ = <-c
	//for i := 0; i < 2; i++ {
	//	time.Sleep(1 * time.Second)
	//	fmt.Println("sleep")
	//}
	fmt.Println("main() => Hello World!")
}
