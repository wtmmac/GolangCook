package main

import (
	"fmt"
	"os"
)
import "strings"

func main() {
	str := "hel"
	fmt.Println(strings.Index(str, "a"))
	fmt.Println(os.PathSeparator)
}
