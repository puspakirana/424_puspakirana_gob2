package main

import (
	"fmt"
	"strings"
)

func main() {
	//Implementing channel
	c := make(chan string)

	go introduce("Ariel", c)
	go introduce("Merida", c)
	go introduce("Margo", c)

	msg1 := <-c
	fmt.Println(msg1)

	msg2 := <-c
	fmt.Println(msg2)

	msg3 := <-c
	fmt.Println(msg3)

	close(c)

	fmt.Println(strings.Repeat("-", 25))

}

// Implementing channel
func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}
