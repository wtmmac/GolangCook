package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
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
	db, err := gorm.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test2?charset=utf8")

	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移模式
	//db.AutoMigrate(&Student{})

	// 创建
	//db.Create(&Student{Name: "wyf1", Address: "nj22"})

	// 读取
	//student := Student{}
	//db.First(&student)                   // 查询id为1的product
	//fmt.Println(student.Name)
	//condition := &StudentFields{
	//	Student: &Student{
	//		Name: "wyf1",
	//		//Address: "nj2",
	//	},
	//	Fields: []string{"name1"},
	//}

	var students []Student

	//db = db.Where(condition.Student, condition.Fields)
	db = db.Where(&Student{
		Name: "wyf1",
		Age:  0,
	})
	err = db.Debug().
		// Select("id","name").
		//Limit(pageSize).
		//Offset((pageNumber - 1) * pageSize).
		Find(&students).Error

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Println(students)
	}

}
