package main

import (
	"sync"
)

type MyChannel struct {
	C    chan struct{}
	once sync.Once
}

func NewMyChannel() *MyChannel {
	return &MyChannel{C: make(chan struct{})}
}

func (mc *MyChannel) SafeClose() {
	mc.once.Do(func() {
		close(mc.C)
	})
}

func main() {
	// fmt.Println("<-chan struct{}")
	// ch := make(<-chan struct{}) // 只读channel
	// select {
	// case <-ch:
	// 	fmt.Println("readed h")
	// }
}
