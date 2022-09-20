package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	q := make(chan int)
	send(c)
	FanIn(c, q)
	for v := range q {
		fmt.Println(v)
	}
}
func send(c chan<- int) {
	go func() {
		for i := 0; i <= 100; i++ {
			c <- i
		}
		close(c)
	}()

}
func FanIn(c <-chan int, q chan<- int) {

	go func() {
		for v := range c {

			q <- v

		}
		close(q)
	}()

}
