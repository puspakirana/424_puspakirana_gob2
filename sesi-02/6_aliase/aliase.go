package main

import (
	"fmt"
)

func main() {
	//Slide 3
	//deklarasi variable dengan tipe data uint8
	var a uint8 = 10
	var b byte // byte adalah alias dari tipe data uint8

	b = a //no error, karena byte memiliki data type yang sama dengan uint8
	_ = b

	//Slide 4
	//Mendeklarasi tipe data alias bernama second
	type second = uint

	var hour second = 3600
	fmt.Printf("Hour type: %T\n", hour) // => hour type: uint
	fmt.Println("-----------------")

}
