package main

import (
	"fmt"
	"regexp"
	"strings"
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
	test()
}

func test() {
	str := `小明上街买菜
买了1斤黄瓜花了5元
买了2斤葡萄花了10.5元
`
	arr := strings.Split(str, "\n")
	for _, line := range arr {
		pattern := regexp.MustCompile(`(\d)斤(.*)花了(\d+(\.\d+)?)元`)
		match := pattern.FindAllStringSubmatch(line, 10)
		//fmt.Println(match)
		if len(match) > 0 {
			fmt.Println(match[0][1] + "," + match[0][2] + "," + match[0][3])
		}
	}
}
