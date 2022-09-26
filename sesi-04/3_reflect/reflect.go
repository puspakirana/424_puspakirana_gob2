package main

import (
	"fmt"
	"reflect"
	"strings"
)

// Idetifying Method Information
type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	//Identifying Data type & Value
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("Tipe variabel: ", reflectValue.Type())
	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variabel : ", reflectValue.Int())
	}
	fmt.Println(strings.Repeat("-", 25))

	//Accessing Value Using Interface Method
	fmt.Println("Tipe variabel: ", reflectValue.Type())
	fmt.Println("Nilai variabel: ", reflectValue.Interface())

	var nilai = reflectValue.Interface().(int)
	fmt.Println("Nilai: ", nilai)
	fmt.Println(strings.Repeat("-", 25))

	//Idetifying Method Information
	var s1 = &student{Name: "John Wick", Grade: 2}
	fmt.Println("Nama: ", s1.Name)

	var reflectValue2 = reflect.ValueOf(s1)
	var method = reflectValue2.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("Wick"),
	})

	fmt.Println("Nama: ", s1.Name)
	fmt.Println(strings.Repeat("-", 25))

}
