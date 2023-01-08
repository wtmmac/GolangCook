package main

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

type Student struct {
	ID      int64
	Name    string
	Address string
	Age     int64
}

// OrderFields 作为一个 数据库Order对象+fields字段的组合
// fields用来指定Order中的哪些字段生效
type StudentFields struct {
	Student *Student
	Fields  []string
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:root@tcp(127.0.0.1:3306)/test2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	//defer db.Close()

	studentFields := &StudentFields{
		Student: &Student{
			Name:    "wyf1",
			Address: "nj2",
		},
		Fields: []string{"name", "address"},
	}

	var students []Student

	db = db.Where(studentFields.Student, studentFields.Fields)
	err = db.WithContext(ctx).Debug().
		Find(&students).Error

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(students)
	}
}
