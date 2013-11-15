package main

import (
	"encoding/json"
	//	"github.com/ziutek/mymysql/godrv"
	"fmt"
	"time"
)

type Book struct {
	Title       string
	Author      []string
	Publisher   string
	Price       float64
	IsPublished bool
}

type People struct {
	Name string
	Age  int
}

////////

func main() {
	b := []byte(`{
    "Title":"go programming language",
    "Author":["john","ada","alices"],
    "Publisher":"qinghua",
    "IsPublished":true,
    "Price":99
  }`)
	//先创建一个目标类型的实例对象，用于存放解码后的值
	var book Book
	err := json.Unmarshal(b, &book)
	if err != nil {
		fmt.Println("error in translating,", err.Error())
		return
	}
	fmt.Println(book.Author)

	fmt.Println(time.Now())
	//	test()
	//	file()
	maps()
	//	println("你好")

	fileWrite()

}
