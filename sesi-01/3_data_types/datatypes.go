package main

import "fmt"

func main() {
	//Data type Number (Integer)
	var num1 uint8 = 89
	var num2 int8 = -12
	fmt.Printf("tipe data num1: %T\n", num1)
	fmt.Printf("tipe data num2: %T\n", num2)
	fmt.Println("-----------------")

	//Data type Numer (Float)
	var decimal1 float32 = 3.63
	fmt.Printf("decimal number: %f\n", decimal1)
	fmt.Printf("decimal number: %.3f\n", decimal1)
	fmt.Println("-----------------")

	//data type boolean
	var cond bool = true
	fmt.Printf("is it permitted? %t\n", cond)
	fmt.Println("-----------------")

	//data type string
	var message = `Selamat datang di "Hacktiv8". Salam kenal :). Mari belajar "Scalable Web Service With Go"`
	fmt.Println(message)
	fmt.Println("-----------------")
}
