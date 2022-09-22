package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	greet("Aurora", 16)
	fmt.Println(strings.Repeat("-", 25))

	greet2("Aurora", "Castle")
	fmt.Println(strings.Repeat("-", 25))

	var names = []string{"Jane", "Doe"}
	var printMsg = greet3("Heii", names)
	fmt.Println(printMsg)
	fmt.Println(strings.Repeat("-", 25))

	var diameter float64 = 15
	var area, circumference = calculate(diameter)
	fmt.Println("Area: ", area)
	fmt.Println("Circumference: ", circumference)
	fmt.Println(strings.Repeat("-", 25))

	studentLists := print("Moana", "Merida", "Jasmine", "Ariel", "Belle")
	fmt.Printf("%v\n", studentLists)
	fmt.Println(strings.Repeat("-", 25))

	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}
	result := sum(numberLists...)
	fmt.Println("Result:", result)
	fmt.Println(strings.Repeat("-", 25))

	profile("Ariel", "pasta", "ayam geprek", "ikan roa", "sate padang")
	fmt.Println(strings.Repeat("-", 25))

}

//Functions

func greet(name string, age int8) {
	fmt.Printf("Hello there! My name is %s and I'm %d years old\n", name, age)
}

func greet2(name, address string) {
	fmt.Println("Hello there! My name is", name)
	fmt.Println("I live in", address)
}

// Function (Return)
func greet3(msg string, names []string) string {
	var joinStr = strings.Join(names, " ")
	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result
}

// Function (Returning multiple values)
func calculate(d float64) (float64, float64) {
	//Luas
	var area float64 = math.Pi + math.Pow(d/2, 2)

	//Keliling
	var circumference = math.Pi + d

	return area, circumference
}

// Function (Variadic function #1)
func print(names ...string) []map[string]string {
	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("Student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

// Function (Variadic function #2)
func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}

// Function (Variadic function #3)
func profile(name string, favFoods ...string) {
	mergeFavFoods := strings.Join(favFoods, ", ")

	fmt.Println("Hello There! I'm", name)
	fmt.Println("I really love to eat", mergeFavFoods)
}
