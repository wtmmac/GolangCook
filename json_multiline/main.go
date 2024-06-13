package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// 定义与JSON数据对应的结构体
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 多行JSON数据
	jsonData := `
    [{
        "name": "Alice",
        "age": 30
    },
    {
        "name": "Bob",
        "age": 25
    }]`

	// 用于接收解析后的数据
	var person Person

	// 解析JSON数据到结构体
	dec := json.NewDecoder(strings.NewReader(jsonData))
	for dec.More() {
		err := dec.Decode(&person)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
	}
}
