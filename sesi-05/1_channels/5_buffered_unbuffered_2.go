package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

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

}
