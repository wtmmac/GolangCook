package main

import (
	"GolangCook/initImport/config"
	_ "GolangCook/initImport/config"
	"fmt"
	"sync"
)

var once sync.Once

func main() {
	once.Do(func() {
		fmt.Println(config.LogVer)
	})
}
