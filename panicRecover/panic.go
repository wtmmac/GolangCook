package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() { //必须要先声明defer，否则不能捕获到panic异常
		fmt.Println("enter defer func")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("end of defer")
	}()
	raisePanic()
}

func raisePanic() {
	fmt.Println("before panic")
	panic(errors.New("!!fatal error!"))
	fmt.Println("after panic")
}
