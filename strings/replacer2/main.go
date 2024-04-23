package main

import (
	"fmt"
	"strings"
)

const refString = `{
    
	"text": "消息内容测试",
	"title": "sampleLink消息测试",
	"picUrl": "@lADOADmaWMzazQKA",
	"messageUrl": "http://dingtalk.com"
   
}`

func main() {
	replacer := strings.NewReplacer(" ", "", "\n", "", "\t", "")
	out := replacer.Replace(refString)
	fmt.Println(out)

	refStringBak := strings.ReplaceAll(refString, "\n", "")
	refStringBak = strings.ReplaceAll(refStringBak, " ", "")
	fmt.Println(refStringBak)
}
