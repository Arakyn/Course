package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 42

	fmt.Println(<-c)

	c1 := make(chan int)
	go func() {
		c1 <- 42

	}()
	defer fmt.Println(<-c1)
}
