package main

import (
	"fmt"
	"strings"
)

// Struct
type Employee struct {
	name     string
	age      int
	division string
}

type Person struct {
	name string
	age  int
}

type Employee2 struct {
	division string
	person   Person
}

func main() {
	//Giving value to struct
	var employee Employee
	employee.name = "Aurora"
	employee.age = 16
	employee.division = "Curse Security"
	fmt.Println(employee.name)
	fmt.Println(employee.age)
	fmt.Println(employee.division)
	fmt.Println(strings.Repeat("-", 25))

	//Initializing Struct
	var employee1 = Employee{}
	employee1.name = "Cinderella"
	employee1.age = 20
	employee1.division = "Fashion Design"
	var employee2 = Employee{name: "Merida", age: 17, division: "War General and City Safety"}
	fmt.Printf("Employee1: %+v\n", employee1)
	fmt.Printf("Employee: %+v\n", employee2)
	fmt.Println(strings.Repeat("-", 25))

	//Pointer to a struct
	var employee3 = Employee{name: "Ariel", age: 15, division: "Sea and Land Peace"}
	var employee4 *Employee = &employee3
	fmt.Println("Employee3 name:", employee3.name)
	fmt.Println("Employee4 name:", employee4.name)
	fmt.Println(strings.Repeat("#", 20))
	employee4.name = "Moana"
	fmt.Println("Employee3 name:", employee3.name)
	fmt.Println("Employee4 name:", employee4.name)
	fmt.Println(strings.Repeat("-", 25))

	//Embedded struct
	var employee5 = Employee2{}
	employee5.person.name = "Elsa"
	employee5.person.age = 25
	employee5.division = "Winter Guardian"
	fmt.Printf("%+v", employee5)
	fmt.Println(strings.Repeat("-", 25))

	//Anonymous Struct
	//Anonymous Struct tanpa pengisian field

	var employee6 = struct {
		person   Person
		division string
	}{}
	employee6.person = Person{name: "Ariel", age: 15}
	employee6.division = "Sea & Land Peace"

	//Anonymous Struct dengan pengisian field
	var employee7 = struct {
		person   Person
		division string
	}{
		person:   Person{name: "Moana", age: 17},
		division: "Ocean Guardian",
	}

	fmt.Printf("Employee6: %+v\n", employee6)
	fmt.Printf("Employee7: %+v\n", employee7)
	fmt.Println(strings.Repeat("-", 25))

	//Slice of Struct
	var people = []Person{
		{name: "Belle", age: 18},
		{name: "Jasmine", age: 15},
		{name: "Sofia", age: 10},
	}

	for _, v := range people {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println(strings.Repeat("-", 25))

	//Slice of anonymous struct
	var employee8 = []struct {
		person   Person
		division string
	}{
		{person: Person{name: "Belle", age: 18}, division: "Personal Development"},
		{person: Person{name: "Jasmine", age: 15}, division: "Magic Carpet Training"},
		{person: Person{name: "Sofia", age: 10}, division: "Friendship and Relation"},
	}

	for _, v := range employee8 {
		fmt.Printf("%+v\n", v)
	}
	fmt.Println(strings.Repeat("-", 25))

}
