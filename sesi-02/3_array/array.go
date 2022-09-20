package main

import (
	"fmt"
	"strings"
)

func main() {
	//Array
	var numbers [4]int
	numbers = [4]int{1, 2, 3, 4}

	var princesses = [3]string{"Ariel", "Aurora", "Belle"}

	fmt.Printf("%#v\n", numbers)
	fmt.Printf("%#v\n", princesses)

	fmt.Println("-----------------")

	//Modify element through index

	var fruits = [3]string{"Apple", "Banana", "Mango"}
	fruits[0] = "Orange"
	fruits[1] = "Strawberry"
	fruits[2] = "Pineapple"

	fmt.Printf("%#v\n", fruits)

	fmt.Println("-----------------")

	//Loop through elements
	var cars = [3]string{"audi", "mercedes", "porsche"}

	//Cara Pertama
	for i, v := range cars {
		fmt.Printf("Index: %d, Value: %s\n", i, v)
	}

	fmt.Println(strings.Repeat("#", 25))

	//Cara Kedua
	for i := 0; i < len(cars); i++ {
		fmt.Printf("Index: %d, Value: %s\n", i, cars[i])
	}
	fmt.Println("-----------------")

	//Multidimensional Array
	balances := [2][3]int{{5, 6, 7}, {8, 9, 10}}

	for _, arr := range balances {
		for _, value := range arr {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
	fmt.Println("-----------------")
}
