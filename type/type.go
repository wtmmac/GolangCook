package main

import (
	"fmt"
)

func IsBigEndian() bool {
	var i int32 = 0x12345678
	var b byte = byte(i)
	if b == 0x12 {
		return true
	}
	return false
}

func main() {
	// 大类型向小类型转换
	var gid int32 = 0x12345678
	var uid int8 = int8(gid)
	fmt.Printf("uid=0x%02x, gid=0x%02x\n", uid, gid)

	if IsBigEndian() {
		fmt.Println("大端序")
	} else {
		fmt.Println("小端序")
	}
}
