package main

import "fmt"

func main() {
	var a interface{} = nil
	aa, ok := a.([]int)
	fmt.Println(aa)
	fmt.Println(ok)
}
