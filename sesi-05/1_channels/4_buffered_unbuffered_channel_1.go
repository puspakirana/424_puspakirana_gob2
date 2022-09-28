package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {

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

}
