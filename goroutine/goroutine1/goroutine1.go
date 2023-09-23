package main

import (
	"fmt"
	"time"
)

func ready(w string, sec time.Duration) {
	time.Sleep(sec * time.Second)
	fmt.Println(w, "is ready!")
}

func main() {
	go ready("Tee", 2)
	go ready("Coffee", 1)
	fmt.Println("I'm waiting")
	time.Sleep(5 * time.Second)
}

//Output:
//I'm waiting // 立刻 Coffee is ready // 1秒后 Tee is ready // 2秒后
