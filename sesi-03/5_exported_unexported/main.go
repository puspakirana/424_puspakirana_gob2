package main

import "5_exported_unexported/helpers"

func main() {
	helpers.Greet()

	var person = helpers.Person{}
	person.Invokegreet()
}
