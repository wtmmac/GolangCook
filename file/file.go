package main

import (
	"fmt"
	"log"
	"os"
)

func file() {

	f, err := os.Create("asdf.txt")
	if err != nil {
		log.Fatal("could not create file :", err)
	}

	defer func() {
		_ = f.Close()
	}()

	fmt.Printf("asdf测试\n")
}
