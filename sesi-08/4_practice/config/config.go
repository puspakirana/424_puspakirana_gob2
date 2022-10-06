package config

import (
	"../structs"
	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3360)/godb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to conenct to database")
	}

	db.AutoMigrate(structs.Person{})
	return db
}
