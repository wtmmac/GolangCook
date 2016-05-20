package main

import (
	"fmt"
	//	"fmt"
	"net/url"
)

func main() {
	str := "dm=%7B%22msgtype%22%3A2%2C%22effect%22%3A0%2C%22alpha%22%3A1%2C%22size%22%3A1%2C%22pos%22%3A3%2C%22color%22%3A16711680%2C%22commit_time%22%3A1444465369879%2C%22user_code%22%3A%2246313317%22%2C%22data%22%3A%22%2F%5C%5C%E6%8E%A5%5C%22%E5%8F%A3%E5%B0%B1%22%7D&liveid=1&sign=d2b03574387bcc2a34e54203e9301c17&time=1449157464"
	//	urlObj, err := url.ParseRequestURI(str)
	urlObj, err := url.ParseQuery(str)

	if err != nil {
		panic("invalid url string")
	}

	fmt.Println(urlObj)

	if value, ok := urlObj["dm"]; ok {
		fmt.Println(value[0])
	}

	if value, ok := urlObj["ab"]; ok {
		fmt.Println(value[0])
	}

}
