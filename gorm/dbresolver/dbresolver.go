package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"log"
)

func main() {
	dsn1 := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	dsn2 := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"
	dsn3 := "root:root@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn1), &gorm.Config{})
	if err != nil {
		log.Default()
	}

	db.Use(dbresolver.Register(dbresolver.Config{
		// `db2` 作为 sources，`db3`、`db4` 作为 replicas
		Sources:  []gorm.Dialector{mysql.Open(dsn2)},
		Replicas: []gorm.Dialector{mysql.Open(dsn3)},
		// sources/replicas 负载均衡策略
		Policy: dbresolver.RandomPolicy{},
	}))



}
