package main

import (
	"fmt"
)

func main() {
	//Defer 1
	defer fmt.Println("defer function starts to execute")
	fmt.Println("Hai everyone")
	fmt.Println("Welcome back to Go learning center")

}
