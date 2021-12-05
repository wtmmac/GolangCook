package main

import "fmt"

func testa() {
	fmt.Println("aaaa")
}

func testb() {
	//显示调用panic函数
	panic("this is a panic test")
	fmt.Println("bbbb")
}

func testc() {
	fmt.Println("cccc")
}

func main() {
	testa()
	//不会打印
	testb()
	testc()
}
