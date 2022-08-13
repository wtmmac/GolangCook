package main

import (
	"fmt"
	"sync"
)

var (
	wg sync.WaitGroup
)

func printNumber(ch chan int) {
	var once sync.Once
	var tag int

	defer wg.Done()
	for {
		num, ok := <-ch
		if !ok {
			return
		}
		once.Do(func() { tag = num })

		fmt.Printf("routine %d:%d\n", tag, num)

		if num > 99 {
			close(ch)
			return
		}
		num++
		ch <- num
	}
}

func printNumber2(name string, ch chan int) {
	defer wg.Done()
	for {
		num, ok := <-ch
		if !ok {
			return
		}

		fmt.Printf("routine %s:%d\n", name, num)

		if num > 99 {
			close(ch)
			return
		}
		num++
		ch <- num
	}
}

func main() {
	ch := make(chan int)
	wg.Add(2)
	go printNumber2("x", ch)
	go printNumber2("y", ch)
	ch <- 1
	wg.Wait()
}
