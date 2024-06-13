package main

import (
	"fmt"

	"github.com/bitly/go-simplejson"
)

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
	data := []byte(`[{"id":1,"title":"5斤","price":0.01,"code":"","stock":0,"goodsId":1,"online":false,"picture":"https://52baobei.oss-cn-shanghai.aliyuncs.com/assets/2024/06/03/ac2439d8-1c3d-4d53-a69c-c923e24682e2.ico","specs":"[{\"keyId\":1,\"key\":\"苹果小箱\",\"valueId\":1,\"value\":\"5斤\"}]"},{"id":2,"title":"10斤","price":0.01,"code":"","stock":5,"goodsId":1,"online":true,"picture":"","specs":"[{\"keyId\":1,\"key\":\"苹果小箱\",\"valueId\":2,\"value\":\"10斤\"}]"}]`)

	// sku := CMSSKUReq{}

	js, err1 := simplejson.NewJson(data) // 反序列化
	if err1 != nil {
		panic(err1.Error())
	}

	arr, err := js.Array()
	if err != nil {
		fmt.Println(err)
	}
	for _, v := range arr {
		fmt.Println(v) //prints 0, 1, 2
	}
	// jsonData, _ := json.Marshal(mediaInfo)
	// fmt.Println(string(jsonData))

}
