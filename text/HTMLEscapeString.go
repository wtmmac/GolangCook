package main

import (
	"fmt"
	"text/template"
)

func main() {
	str := "<&"
	fmt.Println(template.HTMLEscapeString(str))
}
