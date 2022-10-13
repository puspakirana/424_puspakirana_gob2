package main

import (
	"reload/controllers"
	config "reload/database"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var PORT = ":8081"
	db := config.DBInit()
	inDB := &controllers.InDB{DB: db}

	router := gin.Default()

	router.POST("/reload/", inDB.CreateData)
	router.GET("/reload/latest", inDB.GetLatestData)

	router.Run(PORT)
}
