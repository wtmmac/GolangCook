package main

import (
	"fmt"
	"github.com/cloudflare/ahocorasick"
	"strconv"
)

func main() {

	strings := []string{"测试", "Steel", "tee"}

	strings = append(strings, "你好")

	m := ahocorasick.NewStringMatcher(strings)

	m = ahocorasick.NewStringMatcher(strings)

	hits := m.Match([]byte("The Man Of Steel: Superman测试你好"))

	fmt.Println("匹配个数:" + strconv.Itoa(len(hits)))

	for _, i := range hits {
		fmt.Println(strings[i])
	}
}
