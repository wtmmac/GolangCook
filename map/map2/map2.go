package main

import (
	"fmt"
)

// A Header represents the key-value pairs in an HTTP header.
type Header map[string][]string


const (
    pi = 3.14
)

func main() {
	Header := map[string][]string{
		"Accept-Encoding": {"gzip, deflate"},
		"Accept-Language": {"en-us"},
		"Connection": {"keep-alive"},
	}
	fmt.Println(Header)
	fmt.Println(pi)
}