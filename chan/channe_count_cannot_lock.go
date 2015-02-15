package main

import "fmt"

func main() {

	chs := make(chan int, 10)
	fmt.Println("init line",len(chs),"\r\n")

	for i := 0; i < 20; i++ {

		select {
		case chs <- i:

		default:
			fmt.Println("push")
			fmt.Println(len(chs))
			//break
		}
	}

	fmt.Println(fmt.Sprint("chs length is :", len(chs)))

	for i := 0; i < 20; i++ {
		select {
		case value := <-chs:
			fmt.Println(value)
		default:
			fmt.Println("pull")
			fmt.Println(len(chs))
			break
		}

	}
}
