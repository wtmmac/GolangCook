package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("time.Now is:", time.Now())

	start := time.Now()
	time.Sleep(2 * time.Second)
	t1 := time.Now()

	spend_time := t1.Sub(start)

	fmt.Printf("sleep cost: %v\n", spend_time)
	fmt.Printf("spend_time TYPE: %T\n", start)
}
