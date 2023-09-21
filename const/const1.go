package main

import "fmt"

const (
	e, f = iota, 1 << iota
	g, h
	i, j
)

func main() {
	fmt.Println(e, f, g, h, i, j) //	0 1 1 2 2 4
}
