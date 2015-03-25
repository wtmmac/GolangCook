package main

import "fmt"

func main() {

	chs := make(chan int, 10)

	for i := 0; i < 10; i++ {
		chs <- i
	}

	fmt.Println("Channel chs的长度:", len(chs))

loopLabel:
	for {
		select {
		case value := <-chs:
			fmt.Println(value)
		default:
			break loopLabel
		}
	}

	fmt.Println("Channel chs的长度:", len(chs))
}
