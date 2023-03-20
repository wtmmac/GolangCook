package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()

	go func(ctx context.Context) {
		execResult := make(chan bool)

		// 模拟业务逻辑
		go func(execResult chan<- bool) {
			// 模拟处理超时

			for {
				time.Sleep(1 * time.Second)
				fmt.Println("biz")
			}
			execResult <- true
		}(execResult)

		// 等待结果
		select {
		case <-ctx.Done():
			fmt.Println("超时退出")
			return
		case <-execResult:
			fmt.Println("处理完成")
			return
		}
	}(ctx)
	time.Sleep(1000 * time.Second)
}
