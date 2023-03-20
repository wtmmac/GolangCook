package main

import (
	"fmt"
	"go/ast"
	"go/token"
	"os"
)

func GuessType(obj interface{}) {
	fset := token.NewFileSet()
	ast.Print(fset, obj)
}

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

	// print information by format %+v
	sprintf := fmt.Sprintf("%+v", fileInfo)
	fmt.Println(sprintf)

	// by ast
	GuessType(sprintf)

	fileSize := fileInfo.Size()      //文件大小
	buffer := make([]byte, fileSize) //设置一个byte的数组(buffer)

	n, err := file.Read(buffer) //读文件(拿取)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("字节数:", n)
	fmt.Println("bytestream to string:\n", string(buffer))
}
