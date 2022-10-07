package config

import (
	"fmt"
	"practice/structs"

	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:mysql@123@tcp(127.0.0.1:3306)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to conenct to database")
	}

	fmt.Println("Connected")

	db.AutoMigrate(structs.Person{})
	return db
}
