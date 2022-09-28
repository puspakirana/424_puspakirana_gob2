package main

import (
	"fmt"
	"strings"
)

func main() {

	//Channel directions
	c3 := make(chan string)
	students2 := []string{"Cinderella", "Aurora", "Maleficent"}

	for _, v := range students2 {
		go introduce(v, c3)
	}

	for i := 1; i <= 3; i++ {
		print(c3)
	}

	close(c3)
	fmt.Println(strings.Repeat("-", 25))

}

// Channel directions
func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}

// Channel directions
func print(c chan string) {
	fmt.Println(<-c)
}
