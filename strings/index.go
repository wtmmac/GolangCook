package main

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	str := "hel"
	fmt.Println(strings.Index(str, "a"))
	fmt.Println(os.PathSeparator)

	emoji := "🤦"
	fmt.Println(utf8.RuneCountInString(emoji))
}
