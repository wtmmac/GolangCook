package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := `asdf asdf 
	
	
	as 
	 df `
	//将HTML标签全转换成小写
	//re, _ := regexp.Compile("\\s")
	//替换掉注释和一些标签
	reg := regexp.MustCompile("\\s")
	str = reg.ReplaceAllString(str, "")
	fmt.Println(str)
}
