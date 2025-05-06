package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "7354 5567 961877:9500"
	str = strings.TrimSpace(str)            // 去除首尾空白
	str = strings.ReplaceAll(str, " ", "")  // 去除空格
	str = strings.ReplaceAll(str, "\t", "") // 去除制表符
	str = strings.ReplaceAll(str, "\n", "") // 去除换行符
	fmt.Println(str)
}
