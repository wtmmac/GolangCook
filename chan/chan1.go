package main

import (
	"fmt"
	"time"
)

// 无缓冲通道
var c = make(chan bool)
var str1 string

func chan1() {
	defer fmt.Println("defer of chan1")
	str1 = "chan1() => hello, world"
	fmt.Println(str1)
	c <- true //通道锁定 等待取出
	fmt.Println("end of chan1")
}

func main() {
	go chan1()
	_ = <-c
	//for i := 0; i < 2; i++ {
	time.Sleep(2 * time.Second)
	//	fmt.Println("sleep")
	//}
	fmt.Println("main() => Hello World!")
}
