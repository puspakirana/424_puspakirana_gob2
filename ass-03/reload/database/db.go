package config

import (
	"fmt"
	"reload/models"

	"github.com/jinzhu/gorm"
)

// DBInit create connection to database
func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:mysql@123@tcp(127.0.0.1:3306)/autoReload?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to conenct to database")
	}

	fmt.Println("Connected")

	db.AutoMigrate(models.Reload{})
	return db
}
