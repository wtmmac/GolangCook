package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t.Format("15:04:05"))
	fmt.Println(t.Second())
}
