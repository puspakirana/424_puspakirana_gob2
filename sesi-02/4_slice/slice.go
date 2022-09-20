package main

import (
	"fmt"
	"strings"
)

func main() {
	//Slice
	var fruits = []string{"apple", "banana", "mango"}
	_ = fruits
	fmt.Printf("%#v\n", fruits)
	fmt.Println("-----------------")

	//Slice (make function)
	var cars = make([]string, 3)
	_ = cars
	fmt.Printf("%#v\n", cars)
	fmt.Println("-----------------")

	//Slice (append function)

	//Assign with index
	var princesses = make([]string, 3)
	_ = princesses

	princesses[0] = "Anna"
	princesses[1] = "Mulan"
	princesses[2] = "Merida"

	fmt.Printf("%#v\n", princesses)

	fmt.Println(strings.Repeat("#", 25))

	//Append to slice

	princesses = append(princesses, "Ariel", "Cinderella", "Aurora")

	fmt.Printf("%#v\n", princesses)

	fmt.Println("-----------------")

	//Slice (append function with ellipsis)
	var fruits1 = []string{"apple", "banana", "mango"}
	var fruits2 = []string{"durian", "pineapple", "starfruit"}
	fruits1 = append(fruits1, fruits2...)
	fmt.Printf("%#v\n", fruits1)
	fmt.Println("-----------------")

	//Slice (copy function)
	var car1 = []string{"toyota", "honda", "suzuki"}
	var car2 = []string{"bmw", "merceder", "maserati"}

	nn := copy(car1, car2)

	fmt.Println("Car1 => ", car1)
	fmt.Println("Car2 => ", car2)
	fmt.Println("Copied element => ", nn)
	fmt.Println("-----------------")

	//Slice (Slicing)
	var foods = []string{"noodle", "rendang", "burger", "pizza", "meatball"}
	var food1 = foods[1:4]
	fmt.Printf("%#v\n", food1)

	var food2 = foods[0:]
	fmt.Printf("%#v\n", food2)

	var food3 = foods[:3]
	fmt.Printf("%#v\n", food3)

	var food4 = foods[:]
	fmt.Printf("%#v\n", food4)
	fmt.Println("-----------------")

	//Slice (combining slicing and append)
	var drinks = []string{"latte", "cappucino", "green tea", "jasmine tea", "milk"}
	drinks = append(drinks[:3], "water")
	fmt.Printf("%#v\n", drinks)
	fmt.Println("-----------------")

	//Backing array
	var superheroes1 = []string{"superman", "batman", "ironman", "spiderman", "hulk"}
	var superheroes2 = superheroes1[2:4]
	superheroes2[0] = "aquaman"

	fmt.Println("Superheroes 1 =>", superheroes1)
	fmt.Println("Superheroes 2 =>", superheroes2)
	fmt.Println("-----------------")

	//Slice (cap function)
	var singers1 = []string{"tulus", "agnezmo", "raisa", "isyana", "vidi"}
	fmt.Println("Singers 1 cap:", cap(singers1))
	fmt.Println("Singers 1 len:", len(singers1))

	fmt.Println(strings.Repeat("#", 25))

	var singers2 = singers1[0:3]
	fmt.Println("Singers 2 cap:", cap(singers2))
	fmt.Println("Singers 2 len:", len(singers2))

	fmt.Println(strings.Repeat("#", 25))

	var singers3 = singers1[1:]
	fmt.Println("Singers 3 cap:", cap(singers3))
	fmt.Println("Singers 3 len:", len(singers3))

	fmt.Println("-----------------")

	//Slice (creating a new backing array)
	animals := []string{"panda", "tiger", "seal", "lion"}
	newAnimals := []string{}

	newAnimals = append(newAnimals, animals[0:2]...)

	animals[0] = "rabbit"
	fmt.Println("animals: ", animals)
	fmt.Println("newAnimals: ", newAnimals)

	fmt.Println("-----------------")

}
