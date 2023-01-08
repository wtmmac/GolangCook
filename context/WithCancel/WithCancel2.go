package main

import (
	"context"
	"fmt"
	"time"
)

func parent() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	defer func() {
		fmt.Println("parent ok")
		cancel()
	}()
	go child(ctx)
	time.Sleep(time.Second * 1)
}

func child(ctx context.Context) {

	ctx, cancel := context.WithCancel(ctx)
	defer func() {
		fmt.Println("child ok")
		cancel()
	}()
	time.Sleep(time.Second * 10)
}

func main() {
	parent()
	time.Sleep(time.Second * 10)
}
