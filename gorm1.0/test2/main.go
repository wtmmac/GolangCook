package main

import (
	"fmt"

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Animal struct {
	ID   int64
	Name string
	Age  int64
}

//https://gorm.io/zh_CN/docs/query.html
func main() {
	db, err := gorm.Open("mysql", "root:root@/gormdemo?charset=utf8&parseTime=true&loc=Local")
	if err != nil {
		fmt.Println("connect db error: ", err)
	}
	defer db.Close()

	//根据逐渐查询第一条记录
	var animal Animal
	db.First(&animal)
	fmt.Println(animal)
}
