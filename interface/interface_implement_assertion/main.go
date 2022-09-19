package main

import (
	"fmt"
	"os"
)

func main() {

	_, err := os.Open("testFile.txt")
	if err != nil {
		// 类型转换
		mErr, ok := err.(*MyError)
		if ok {
			fmt.Println("MyError:" , mErr.Code)
		} else {
			fmt.Println("false")
		}
	}

}
