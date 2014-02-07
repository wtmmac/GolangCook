package main

import "fmt"
import "strings"

func main() {

	ages := map[string]int{
		"lili":  13,
		"nick":  23,
		"jacky": 55,
	}

	str := "hel"
	fmt.Println(strings.Index(str, "a"))
}
