package main

import "fmt"

func main() {

	chs := make(chan int, 10)
	fmt.Println(len(chs))

	for i := 0; i < 11; i++ {
		chs <- i
	}

	fmt.Println(len(chs))

	for {
		select {
		case value := <-chs:
			fmt.Println(value)
		default:
			fmt.Println("default")
			fmt.Println(len(chs))
			break
		}

	}

    fmt.Println(len(chs))
}
