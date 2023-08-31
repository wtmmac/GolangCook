package main

import (
	"fmt"
	"github.com/wtmmac/GolangCook/initImport/config"
	"sync"
)

var once sync.Once

func main() {
	once.Do(func() {
		fmt.Println(config.LogVer)
	})
}
