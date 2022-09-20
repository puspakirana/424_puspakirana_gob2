package main

import "fmt"

func main() {
	//constant variable
	const full_name string = "John Doe"
	fmt.Printf("Hello %s\n", full_name)
	fmt.Println("-----------------")

	//Operators
	var value = 2 + 2*3
	fmt.Println(value)

	var value2 = (2 + 2) * 3
	fmt.Println(value2)

	fmt.Println("-----------------")

	//Relational Operators

	var firstCondition bool = 2 < 3
	var secondCondition bool = "joey" == "Joey"
	var thirdCondition bool = 10 != 2.3
	var forthCondition bool = 11 <= 11

	fmt.Println("first condition:", firstCondition)
	fmt.Println("second condition:", secondCondition)
	fmt.Println("third condition:", thirdCondition)
	fmt.Println("forth condition:", forthCondition)

	fmt.Println("-----------------")

	//Logical Operators
	var right = true
	var wrong = false

	var wrongAndRight = wrong && right
	fmt.Printf("Wrong && Right \t(%t) \n", wrongAndRight)

	var wrongOrRight = wrong || right
	fmt.Printf("Wrong || Right \t(%t) \n", wrongOrRight)

	var wrongReverse = !wrong
	fmt.Printf("!wrong \t\t(%t) \n", wrongReverse)
}
