package main

import (
	"fmt"

	"github.com/leoantony72/goswift"
)

func main() {
	cache := goswift.NewCache()

	// Value 0 indicates no expiry
	// cache.Set("key", "value", 0)

	val, err := cache.Get("key")
	if err != nil {
		if err.Error() == goswift.ErrKeyNotFound {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("key", val)
}
