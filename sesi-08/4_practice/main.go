package main

import (
	"config"
	"fmt"
)

func main() {
	db := config.DBInit()
	_ = db
	fmt.Println("success")
}
