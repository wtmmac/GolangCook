package main

import (
	"fmt"
	"runtime"
	//	"strconv"
	//	"time"
)

func makeCakeAndSend(s string) {
	fmt.Println("func:makeCakeAndSend")
	panic("test..")
	fmt.Println(s)
}

func receiveCakeAndPack(cs chan string) {
	for i := 1; i <= 3; i++ {
		s := <-cs //接收channel上的蛋糕
		fmt.Println("=======>Packing received cake: ", s)
	}
}

type myRoutine func(s string)

func safeRoutine(s string, fn myRoutine) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Runtime error caught: %v\n", r)
			go safeRoutine(s, makeCakeAndSend)
			runtime.Goexit()
		}
	}()
	fn(s)
}

func main() {
	//	cs := make(chan string)
	//	go makeCakeAndSend(cs)
	s := "字符串"
	safeRoutine(s, makeCakeAndSend)
	//go receiveCakeAndPack(cs)

	//休眠一段时间以防程序立即退出
	//time.Sleep(5 * time.Second)
}
