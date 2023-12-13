package main

import "fmt"

func main() {
	response := "this is a response"
	for _, c := range response {
		fmt.Println(string(c))
	}
}
