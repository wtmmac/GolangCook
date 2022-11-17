package main

import "fmt"

func main() {
	v1 := false
	v2 := true
	v3 := false

	if v1 && v2 && v3 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}

	if !v1 || !v2 || !v3 {
		fmt.Println("false")
	} else {
		fmt.Println("true")
	}

}
