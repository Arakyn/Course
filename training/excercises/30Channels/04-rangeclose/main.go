package main

import (
	"fmt"
)

func main() {

	c := make(chan int) //creating a channel

	go foo(c) // putting in stuff the channel

	// now we range over the channel
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("About to Exit.")
}
func foo(c chan<- int) {
	for i := 1; i < 100; i++ {

		c <- 32

	}
	close(c)

}
