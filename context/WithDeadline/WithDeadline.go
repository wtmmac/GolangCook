package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	d := time.Now().Add(50 * time.Millisecond)
	ctx, cancel := context.WithDeadline(context.Background(), d)

	// 尽管ctx会过期，但在任何情况下调用它的cancel函数都是很好的实践。
	// 如果不这样做，可能会使上下文（context）及其父类存活的时间超过必要的时间。
	defer cancel()

	for {
		// 哪个先调用，就执行那个case
		select {
		case <-time.After(1 * time.Second): // 一秒钟取一次值
			fmt.Println("overslept")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
