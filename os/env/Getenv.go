package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("**************************")
	goPath := os.Getenv("GOPATH")
	fmt.Printf("GOPATH is %s\n",goPath)
}
