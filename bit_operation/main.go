package main

import "fmt"

func main() {
	kindMask := (1 << 5) - 1
	fmt.Println(kindMask)

	const size = 64 << 10
	fmt.Println(size)
}
