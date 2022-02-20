package main

import (
	"GolangTest/initImport/config"
	_ "GolangTest/initImport/config"
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	once.Do(func() {
		fmt.Println(config.LogVer)
	})
}
