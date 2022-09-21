package main

import "fmt"

func main() {

	//Temporary Variable
	var currentYear = 2021

	if age := currentYear - 2010; age < 17 {
		fmt.Printf("Kamu belum boleh membuat kartu sim karena umur anda => %d\n", age)
	} else {
		fmt.Println("Kamu sudah boleh membuat kartu sim")
	}
	fmt.Println("-----------------")

	//Switch
	var score = 8

	switch score {
	case 8:
		fmt.Println("perfect")
	case 7:
		fmt.Println("awesome")
	default:
		fmt.Println("not bad")
	}
	fmt.Println("-----------------")

	//Switch with relational operators
	var score2 = 6

	switch {
	case score2 == 8:
		fmt.Println("perfect")
	case (score2 < 8) && (score2 > 3):
		fmt.Println("not bad")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("you need to learn more")
		}
	}
	fmt.Println("-----------------")

	//Switch (fallthrough keyword)
	var score3 = 6

	switch {
	case score3 == 8:
		fmt.Println("perfect")
	case (score3 < 8) && (score3 > 3):
		fmt.Println("not bad")
		fallthrough
	case score3 < 5:
		fmt.Println("It is ok, but please study harder")
	default:
		{
			fmt.Println("study harder")
			fmt.Println("You don't have a good score yet")
		}
	}
	fmt.Println("-----------------")

	//Nested condition
	var score4 = 10

	if score4 > 7 {
		switch score4 {
		case 10:
			fmt.Println("perfect")
		default:
			fmt.Println("nice!")
		}
	} else {
		if score4 == 5 {
			fmt.Println("not bad")
		} else if score4 == 3 {
			fmt.Println("keep trying")
		} else {
			fmt.Println("You can do it")
			if score4 == 0 {
				fmt.Println("try harder!")
			}
		}
	}
	fmt.Println("-----------------")

}
