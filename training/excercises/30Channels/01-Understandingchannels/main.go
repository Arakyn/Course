package main

import "fmt"

func main() {
	c := make(chan int)
	// c <- 42    // to make it work we are just gonna add this into a go func
	go func() {
		c <- 42
	}()
	fmt.Println(<-c)
	c1 := make(chan int, 1)
	c1 <- 132
	fmt.Println(<-c1)
	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2
	fmt.Println(<-c2)

	fmt.Println(<-c2)

}

// to make it work
