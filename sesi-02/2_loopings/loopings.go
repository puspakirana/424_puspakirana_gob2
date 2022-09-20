package main

import "fmt"

func main() {
	//First way of looping
	for i := 0; i < 3; i++ {
		fmt.Println("Angka", i)
	}
	fmt.Println("-----------------")

	//Second way of looping
	var j = 0
	for j < 3 {
		fmt.Println("Angka", j)
		j++
	}
	fmt.Println("-----------------")

	//Third way of looping
	var k = 0
	for {
		fmt.Println("Angka", k)
		k++
		if k == 3 {
			break
		}
	}
	fmt.Println("-----------------")

	//Break and continue keywords
	for i := 1; i <= 10; i++ {
		if i%2 == 1 {
			continue
		}

		if i > 8 {
			break
		}

		fmt.Println("Angka", i)
	}
	fmt.Println("-----------------")

	//Nested Looping
	for i := 0; i < 5; i++ {
		for j := i; j < 5; j++ {
			fmt.Print(j, " ")
		}
		fmt.Println()
	}
	fmt.Println("-----------------")

	//Label
outerLoop:
	for i := 0; i < 3; i++ {
		fmt.Println("Perulangan ke - ", i+1)
		for j := 0; j < 3; j++ {
			if i == 2 {
				break outerLoop
			}
			fmt.Print(j, " ")
		}
		fmt.Print("\n")
	}
	fmt.Println("-----------------")

}
