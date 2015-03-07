package main

import "fmt"

//import "strings"

func main() {
ages := map[string]int{
		"lili":  13,
		"nick":  23,
		"jacky": 55,
	}

	if _, dup := ages["lili"]; dup {
		fmt.Println("finded the key", dup)
	}
}
