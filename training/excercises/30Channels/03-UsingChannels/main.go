package main

import (
	"fmt"
)

func main() {

	c := make(chan int)

	go foo(c)
	bar(c)

	fmt.Println("About to Exit.")
}
func foo(c chan<- int) {
	fmt.Println("Putting in the Numbers.")
	c <- 1
	c <- 2

}
func bar(c <-chan int) {
	fmt.Println("Taking out the first number", <-c)
	fmt.Println("Taking out the Second number", <-c)

}
