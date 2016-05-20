package main

import (
	"fmt"
	"log"
	"time"
)

func init() {
	time.Sleep(time.Second)
	log.Println("init")
}

func main() {
	fmt.Println("test")
}
