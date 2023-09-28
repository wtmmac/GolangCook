package main

import "fmt"

func a() *int {
	v := 111
	return &v
}

func main() {
	i := a()
	fmt.Println(i)
}
