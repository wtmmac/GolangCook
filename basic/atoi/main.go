package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	userId := "10010"
	fmt.Println(reflect.TypeOf(userId))
	userIdInt, err := strconv.Atoi(userId)
	if err != nil {
		//
	}
	fmt.Println(userId)
	fmt.Println(reflect.TypeOf(userId))
	fmt.Printf("TypeOf of userIdInt is : %+v\n", reflect.TypeOf(userIdInt))

	fmt.Printf("%f", 1e15)
}
