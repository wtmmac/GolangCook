package main

import "fmt"
import "strings"

func main() {
	str := "hel"
	fmt.Println(strings.HasPrefix(str, "a"))
	fmt.Println(strings.HasPrefix(str, "he"))
}
