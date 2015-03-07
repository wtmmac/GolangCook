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

	if strings.Index(str, "a") > 0 {
		fmt.Println("find")
	} else {
		fmt.Println("not find")
	}

	fmt.Println(ages)
}
