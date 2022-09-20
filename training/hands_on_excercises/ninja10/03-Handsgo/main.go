package main

import (
	"fmt"
)

func main() {
	c := gen()
	recieve(c)

	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}

func gen() <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < 100; i++ {
			c <- i

		}
		close(c)
	}()

	return c
}
func recieve(c <-chan int) <-chan int {

	for v := range c {
		fmt.Println(v)
	}

	return c

}
