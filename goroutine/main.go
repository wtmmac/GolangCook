package main

import (
	"fmt"
	"strconv"
	"time"
)

func makeCakeAndSend(cs chan string) {
	for i := 1; i <= 3; i++ {
		cakeName := "Strawberry Cake " + strconv.Itoa(i)
		fmt.Println("********Making a cake and sending ...", cakeName)
		cs <- cakeName //发送草莓蛋糕
	}
}

func receiveCakeAndPack(cs chan string) {
	for i := 1; i <= 3; i++ {
		s := <-cs //接收channel上的蛋糕
		fmt.Println("=======>Packing received cake: ", s)
	}
}

func main() {
	cs := make(chan string)
	go makeCakeAndSend(cs)
	go receiveCakeAndPack(cs)

	//休眠一段时间以防程序立即退出
	time.Sleep(4 * time.Second)

	fmt.Println(4 * time.Second)
}
