package main

import (
	"fmt"
)

func main() {
	//Defer 2
	callDeferFunc()
	fmt.Println("Hai everyone!")
}

func callDeferFunc() {
	defer deferFunc()
}

func deferFunc() {
	fmt.Println("defer function starts to execute")
}
