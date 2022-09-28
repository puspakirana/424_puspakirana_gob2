package main

import (
	"fmt"
	"strings"
)

func main() {
	//Channel with anonymous function
	c2 := make(chan string)
	students := []string{"Belle", "Sofia", "Agnes"}

	for _, v := range students {
		go func(student string) {
			fmt.Println("Student", student)
			result := fmt.Sprintf("Hai, my name is %s", student)
			c2 <- result
		}(v)
	}

	for i := 1; i <= 3; i++ {
		print(c2)
	}

	close(c2)
	fmt.Println(strings.Repeat("-", 25))

}

// Channel with anonymous function
func print(c chan string) {
	fmt.Println(<-c)
}
