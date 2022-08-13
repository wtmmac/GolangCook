package main

import (
	"fmt"

	models "github.com/wtmmac/GolangTest/modDemo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	user := &models.User{}
	// var users []User
	// works because destination struct is passed in
	db.First(user)
	// SELECT * FROM `users` ORDER BY `users`.`id` LIMIT 1
	fmt.Println(user)

	student := &models.Student{}
	db.First(student)

	fmt.Println(student)
}
