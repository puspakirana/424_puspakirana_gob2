package main

import (
	"fmt"
	"strings"
	"time"
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

	//Buffered vs Unbuffered Channel
	c4 := make(chan int)
	go func(c4 chan int) {
		fmt.Println("func goroutine starts sending sata into the channel")
		c4 <- 10
		fmt.Println("func goroutine after sending data into the channel")
	}(c4)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Println(("main goroutine starts receiving data"))
	d := <-c4
	fmt.Println("main goroutine received data:", d)

	close(c4)
	time.Sleep(time.Second)
	fmt.Println(strings.Repeat("-", 25))

	//Buffered vs Unbuffered Channel #2
	c5 := make(chan int, 3)

	go func(c5 chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("func goroutine %d starts sending sata into the channel\n", i)
			c5 <- i
			fmt.Printf("func goroutine %d after sending data into the channel\n", i)
		}

		close(c5)
	}(c5)

	fmt.Println("main goroutine sleeps 2 second")
	time.Sleep(time.Second * 2)

	for v := range c5 {
		fmt.Println("main goroutine received vakue from channel: ", v)
	}
	fmt.Println(strings.Repeat("-", 25))

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

// Implementing channel, Channel directions
func introduce(student string, c chan string) {
	result := fmt.Sprintf("Hai, my name is %s", student)
	c <- result
}

// Channel with anonymous function, Channel directions
func print(c chan string) {
	fmt.Println(<-c)
}
