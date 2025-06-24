package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string = "{hello world"
	fmt.Println(strings.Index(a, "ll"))
	fmt.Println(a[0] == '{')
}
