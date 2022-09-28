package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

	//Select
	c6 := make(chan string)
	c7 := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		c6 <- "Hello!"
	}()

	go func() {
		time.Sleep(1 * time.Second)
		c7 <- "Salut!"
	}()

	for i := 1; i <= 2; i++ {
		select {
		case msg1 := <-c6:
			fmt.Println("Received", msg1)
		case msg2 := <-c7:
			fmt.Println("Received", msg2)
		}
	}
	fmt.Println(strings.Repeat("-", 25))

}
