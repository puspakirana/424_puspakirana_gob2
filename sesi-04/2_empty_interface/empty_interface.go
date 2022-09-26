package main

import (
	"fmt"
	"strings"
)

func main() {
	//EMPTY INTERFACE
	var randomValues interface{}

	_ = randomValues
	randomValues = "Jalan Sudirman"
	randomValues = 20
	randomValues = true
	randomValues = []string{"Ariel", "Aurora"}

	// ketika suatu variable mempunyai tipe data empty interface, maka variable tersebut dapat diberikan nilai dengan tipe data apapun.

	fmt.Println(randomValues)
	fmt.Println(strings.Repeat("-", 25))

	//EMPTY INTERFACE: TYPE ASSERTION
	var v interface{}
	v = 20
	// v = v * 9 -> Salah
	//Error tersebut disebabkan karena kita hanya bisa melakukan perkalian terhadap tipe data int yang konkrit atau asli, sedangkan variable v memiliki tipe data empty interface yang diberikan nilai dengan tipe data int.
	if value, ok := v.(int); ok == true {
		v = value * 9
	}
	fmt.Println(v)
	fmt.Println(strings.Repeat("-", 25))

	//EMPTY INTERFACE: MAP & SLICE WITH EMPTY INTERFACE
	rs := []interface{}{1, "Ariel", true, 2, "Aurora", true}

	rm := map[string]interface{}{
		"Name":   "Ariel",
		"Status": true,
		"Age":    15,
	}

	_, _ = rs, rm
	fmt.Println(rs)
	fmt.Println(rm)
	fmt.Println(strings.Repeat("-", 25))

}
