package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("hello.txt") //(开门)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close() //(关门)

	fileInfo, err := file.Stat() //获取文件属性
	if err != nil {
		fmt.Println(err)
		return
	}

	fileSize := fileInfo.Size()      //文件大小
	buffer := make([]byte, fileSize) //设置一个byte的数组(buffer)

	n, err := file.Read(buffer) //读文件(拿取)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("字节数:", n)
	fmt.Println("bytestream to string:", string(buffer))
}
