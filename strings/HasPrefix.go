package main

import "fmt"
import "strings"

func main() {
	str := "hel"
	fmt.Println(strings.HasPrefix(str, "a"))
	fmt.Println(strings.HasPrefix(str, "he"))

	// fetch the first string
	stringParts := strings.Split("192.168.0.1,220.111.10.12,3,4,5", ",")
	if len(stringParts) > 0 {
		fmt.Println(stringParts[0])
	}

}
