package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

// 定义与JSON数据对应的结构体
type CMSSKUReq struct {
	Id      int    `json:"id"`
	Title   string `json:"title" validate:"required"`
	Price   string `json:"price" validate:"required"`
	Code    string `json:"code"`
	Stock   int    `json:"stock"`
	GoodsId int    `json:"goodsId" validate:"required"`
	Online  int    `json:"online"`
	Picture string `json:"picture"`
	Specs   string `json:"specs" validate:"required"`
}

func main() {
	// 无法处理的数据，外层括号和数据中间的逗号
	// jsonData := `[{"id":1,"title":"5斤","price":0.01,"code":"","stock":0,"goodsId":1,"online":false,"picture":"https://52baobei.oss-cn-shanghai.aliyuncs.com/assets/2024/06/03/ac2439d8-1c3d-4d53-a69c-c923e24682e2.ico","specs":"[{\"keyId\":1,\"key\":\"苹果小箱\",\"valueId\":1,\"value\":\"5斤\"}]"},{"id":2,"title":"10斤","price":0.01,"code":"","stock":5,"goodsId":1,"online":true,"picture":"","specs":"[{\"keyId\":1,\"key\":\"苹果小箱\",\"valueId\":2,\"value\":\"10斤\"}]"}]`
	// 多行JSON数据
	jsonData := `{"id":1,"title":"5斤","price":"0.01","code":"","stock":0,"goodsId":1,"online":1,"picture":"https://52baobei.oss-cn-shanghai.aliyuncs.com/assets/2024/06/03/ac2439d8-1c3d-4d53-a69c-c923e24682e2.ico","specs":"[{\"keyId\":1,\"key\":\"苹果小箱\",\"valueId\":1,\"value\":\"5斤\"}]"}{"id":2,"title":"10斤","price":"0.01","code":"","stock":5,"goodsId":1,"online":1,"picture":"","specs":"[{\"keyId\":1,\"key\":\"苹果小箱\",\"valueId\":2,\"value\":\"10斤\"}]"}`

	// 用于接收解析后的数据
	var sku CMSSKUReq

	// 解析JSON数据到结构体
	dec := json.NewDecoder(strings.NewReader(jsonData))
	for dec.More() {
		err := dec.Decode(&sku)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Title: %s, Price: %s\n", sku.Title, sku.Price)
	}
}
