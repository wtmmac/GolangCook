package main

import (
	"fmt"
	"sync"
)

var (
	counter = make(chan int)
	wg      sync.WaitGroup
)

func printTester(counter chan int, isLetter bool) {
	defer wg.Done()
	for {
		c, ok := <-counter
		if !ok {
			return
		}

		switch isLetter {
		case true:
			fmt.Printf("%c\n", c)
		case false:
			fmt.Printf("%d\n", c)
		}

		if c >= 90 {
			close(counter)
			return
		}

		//time.Sleep(time.Millisecond * 50)
		c++
		counter <- c
	}
}

func main() {
	wg.Add(2)
	go printTester(counter, false)
	go printTester(counter, true)

	counter <- 65
	wg.Wait()
}
