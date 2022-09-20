package main

import (
	"fmt"
	"strings"
)

func main() {

	//Map
	//Cara 1
	var person map[string]string //Deklarasi
	person = map[string]string{} //Inisialisasi

	person["name"] = "Jane"
	person["age"] = "25"
	person["address"] = "Holywood"

	fmt.Println("name: ", person["age"])
	fmt.Println("age: ", person["age"])
	fmt.Println("address: ", person["address"])

	fmt.Println(strings.Repeat("#", 25))

	//Cara 2
	var person2 = map[string]string{
		"name":    "Aurora",
		"age":     "16",
		"address": "Castle",
	}

	fmt.Println("name: ", person2["age"])
	fmt.Println("age: ", person2["age"])
	fmt.Println("address: ", person2["address"])

	fmt.Println("-----------------")

	//Map (Looping with map)
	var person3 = map[string]string{
		"name":    "Cinderella",
		"age":     "13",
		"address": "Castle Somewhere Over There",
	}

	for key, value := range person3 {
		fmt.Println(key, ":", value)
	}

	fmt.Println("-----------------")

	//Map (Deleting value)
	var person4 = map[string]string{
		"name":    "Merida",
		"age":     "17",
		"address": "Into the Wood",
	}

	fmt.Println("Before deleting: ", person4)

	delete(person4, "age")

	fmt.Println("After deleting: ", person4)

	fmt.Println("-----------------")

	//Map (Detecting a value)

	var person5 = map[string]string{
		"name":    "Red Riding Hood",
		"age":     "9",
		"address": "Woods",
	}

	value, exist := person5["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value doesn't exist")
	}

	delete(person5, "age")

	value, exist = person5["age"]

	if exist {
		fmt.Println(value)
	} else {
		fmt.Println("Value has been deleted")
	}

	fmt.Println("-----------------")

	//Map (Combining slice with map)
	var pets = []map[string]string{
		{"name": "Milo", "age": "5"},
		{"name": "Peter", "age": "2"},
		{"name": "Grace", "age": "3"},
	}

	for i, pet := range pets {
		fmt.Printf("Index: %d, Name: %s, Age: %s\n", i, pet["name"], pet["age"])
	}

	fmt.Println("-----------------")

}
