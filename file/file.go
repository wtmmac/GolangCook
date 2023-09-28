package main

import (
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
}
