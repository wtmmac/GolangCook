package main

import (
	"fmt"
	"log"
)

// 业务处理1

func main() {
	isDone := make(chan int)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				log.Println("skip panic!")
			}
			isDone <- 1
		}()
		a := []string{"a", "b"}
		fmt.Println(a[3])          // 这里slice越界异常了
		fmt.Println("end of biz1") // 无法执行
	}()
	fmt.Println("before isDone")
	<-isDone
	fmt.Println("after isDone")

	// 业务处理2
	// 打印
	fmt.Println("exit")
}
