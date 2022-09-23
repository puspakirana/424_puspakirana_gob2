package main

import (
	"fmt"
	"strings"
)

func main() {
	//Memory Address
	var firstNumber int = 4
	var secondNumber *int = &firstNumber
	fmt.Println("firstNumber (value) :", firstNumber)
	fmt.Println("firstNumber (memori address) :", &firstNumber)
	fmt.Println("secondNumber (value) :", *secondNumber)
	fmt.Println("secondNumber (memori address) :", secondNumber)
	fmt.Println(strings.Repeat("-", 25))

	//Changing value through pointer
	var firstPerson string = "John"
	var secondPerson *string = &firstPerson
	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("firstPerson (value) :", *secondPerson)
	fmt.Println("firstPerson (memori address) :", secondPerson)

	fmt.Println("\n", strings.Repeat("#", 30), "\n")

	*secondPerson = "Doe"

	fmt.Println("firstPerson (value) :", firstPerson)
	fmt.Println("firstPerson (memori address) :", &firstPerson)
	fmt.Println("firstPerson (value) :", *secondPerson)
	fmt.Println("firstPerson (memori address) :", secondPerson)
	fmt.Println(strings.Repeat("-", 25))

	//Pointer as parameter
	var a int = 10
	fmt.Println("Before:", a)
	changeValue(&a)
	fmt.Println("After:", a)
	fmt.Println(strings.Repeat("-", 25))

}

func changeValue(number *int) {
	*number = 20
}
