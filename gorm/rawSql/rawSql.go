package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Result struct {
	ID   int64
	Name string
	Age  int
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	//defer db.Close()

	var result Result
	db.Raw("SELECT id, name, age FROM user WHERE name = ?", "小刚").Scan(&result)
	fmt.Println(result)

	fmt.Println("query all records")
	// Raw SQL
	rows, err := db.Raw("SELECT id, name, age FROM user").Rows()
	defer rows.Close()
	for rows.Next() {
		db.ScanRows(rows, &result)
		fmt.Println(result.Name)
	}
}
