package main

import (
	"log"
	"time"
)

func main() {
	log.Printf("run...")
	defer log.Printf("stop.")
	i := 0
	c := make(chan int, 1)
	for {
		go func(i int) {
			mem := make([]int, 100*1024*1024)
			log.Printf("i=%d,mem:%p", i, mem)
			mem[0] = <-c
		}(i)
		i++
		time.Sleep(200 * time.Microsecond)
	}
}
