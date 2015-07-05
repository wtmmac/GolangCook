package main

//var c = make(chan int, 10)
var a string

func f() {
	a = "hello, world"
	print("function f")
	//c <- 0
}

func main() {
	go f()
	//<-c
	print("main function")
	print(a)
}
