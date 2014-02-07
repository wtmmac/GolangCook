package main

import "fmt"

var string1 interface{}

func main() {
	string1 = "asdf"
	fmt.Println(string1)

	if k, ok := string1.(string); ok {
		fmt.Println("k is string:", k)
	}
}
