package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	time.Sleep(time.Second)
	log.Println("init...........")
}

func init() {
	time.Sleep(time.Second)
	log.Println("init2...........")
}

func test() {
	log.Println("test--------")
}

func main() {
	fmt.Println("test")
	test()
}
