package main

import (
	"context"
	"fmt"
)

func main() {
	// 在独立的goroutine中生成整数，通过channel传递出去。
	// 一旦context关闭，该goroutine也将安全退出。
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // 当完成整数生产时，关闭Context

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}
