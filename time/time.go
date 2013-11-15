package main

import (
	"fmt"
	"time"
)

func main() {
	//休眠一段时间以防程序立即退出
	time.Sleep(4 * 1e9)

	fmt.Println(4 * 1e9)
}
