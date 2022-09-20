package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	//Looping over String (byte-by-byte)
	city := "Jakarta"
	for i := 0; i < len(city); i++ {
		fmt.Printf("index: %d, byte: %d\n", i, city[i])
	}

	fmt.Println(strings.Repeat("#", 25))

	var city2 []byte = []byte{74, 97, 107, 97, 114, 116, 97}

	fmt.Println(string(city2))

	fmt.Println("-----------------")

	//Looping over String (rune-by-rune)
	str1 := "makan"
	str2 := "mânca"

	fmt.Printf("str1 byte length => %d\n", len(str1))
	fmt.Printf("str2 byte length => %d\n", len(str2))

	fmt.Println(strings.Repeat("#", 25))

	fmt.Printf("str1 character length => %d\n", utf8.RuneCountInString(str1))
	fmt.Printf("str2 character length => %d\n", utf8.RuneCountInString(str2))

	fmt.Println(strings.Repeat("#", 25))

	str := "mânca"

	for i, s := range str {
		fmt.Printf("idex => %d, string => %s\n", i, string(s))
	}

	fmt.Println("-----------------")

}
