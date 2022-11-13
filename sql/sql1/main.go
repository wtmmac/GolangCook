package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//db, err := sql.Open("mysql", "root:root@tcp(localhost)/test?charset=utf8&allowOldPasswords=1")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select * from user")

	for rows.Next() {
		//row.Scan(...)
	}
	rows.Close()
}
