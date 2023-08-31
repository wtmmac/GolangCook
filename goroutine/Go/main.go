package main

import (
	"fmt"
	"github.com/wtmmac/GolangCook/goroutine/Go/GoWrapper"
	"time"
)

func main() {
	GoWrapper.Go(hello)
	time.Sleep(1 * time.Second)
}

func hello() {
	fmt.Println("hello")
}
