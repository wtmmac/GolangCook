package main

import (
	"fmt"
	"github.com/wtmmac/webstrings"
)

func main() {
	str := `a\aa'sf"\\2022'"\测试`
	str_addslashes := webstrings.Addslashes(str)
	str_stripslashes := webstrings.Stripslashes(str_addslashes)
	fmt.Println(str)
	fmt.Println(str_addslashes)
	fmt.Println(str_stripslashes)
	if str == str_stripslashes {
		fmt.Println("true")
	}
}
