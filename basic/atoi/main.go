package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	userId := "10010"
	fmt.Println(reflect.TypeOf(userId))
	_, err := strconv.Atoi(userId)
	if err != nil {
		//
	}
	fmt.Println(userId)
	fmt.Println(reflect.TypeOf(userId))
}
