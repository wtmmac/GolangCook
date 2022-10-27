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

	//根据逐渐查询最后一条记录
	var animal2 Animal
	db.Last(&animal2)
	fmt.Println(animal2)

	//指定某条记录（仅当主键为整型时可用）
	var animal3 Animal
	db.First(&animal3, 2)
	fmt.Println(animal3)

	//where条件

	//符合条件的第一条记录
	var animal4 Animal
	db.Where("name = ?", "demotest2").First(&animal4)
	fmt.Println("where : ", animal4, animal4.ID, animal4.Name, animal4.Age)

	//符合条件的所有记录
	var animals5 []Animal
	db.Where("name = ?", "galeone").Find(&animals5)
	fmt.Println(animals5)
	for k, v := range animals5 {
		fmt.Println("k:", k, "ID:", v.ID, "Name:", v.Name, "Age:", v.Age)
	}

	//IN
	var animals6 []Animal
	db.Where("name IN (?)", []string{"demo-test", "demotest2"}).Find(&animals6)
	fmt.Println(animals6)

	//LIKE
	var animals7 []Animal
	db.Where("name like ?", "%jim%").Find(&animals7)
	fmt.Println(animals7)

	//AND
	var animals8 []Animal
	db.Where("name = ? AND age >= ?", "jim", "24").Find(&animals8)
	fmt.Println(animals8)

	//总数
	var count int
	var animals9 []Animal
	db.Where("name = ?", "galeone").Or("name = ?", "jim").Find(&animals9).Count(&count)
	fmt.Println(animals9)
	fmt.Println(count)

	//Scan, 原生查询
	var animals10 []Animal
	db.Raw("SELECT id, name, age From Animals WHERE name = ? AND age = ? ", "galeone", "30").Scan(&animals10)
	fmt.Println("Scan: ", animals10)

	//原生查询，select all
	var animals11 []Animal
	rows, _ := db.Raw("SELECT id,name FROM Animals").Rows()
	//注意：上面的 select id,name 后面不能写成 * 代替，不然出来的结果都是默认0值
	//像这样结果： ALL: [{0 0} {0 0} {0 0} {0 0} {0 0} {0 0} {0 0}]
	//Scan 后面是什么字段，select 后面就紧跟什么字段
	for rows.Next() {
		var result Animal
		rows.Scan(&result.ID, &result.Name)
		animals11 = append(animals11, result)
	}
	fmt.Println("ALL: ", animals11)
	//output:ALL: [{1 demo-test 0} {2 galeone 0} {3 demotest2 0} {4 galeone 0} {5 galeone 0} {6 jim 0} {7 jimmy 0}]

	//select 查询
	var animal12 Animal
	db.Select("name,age").Find(&animal12) //只查询name，age字段，相当于select name,age from user
	fmt.Println("select: ", animal12)
	// db.Select([]string{"name", "age"}).Find(&animal12)
	// fmt.Println("select2: ", animal12)
}
