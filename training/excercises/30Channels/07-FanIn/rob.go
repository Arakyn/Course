package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Hello")
	c := FanIn(boring("Aman"), boring("YTOLO"))

	for i := 0; i < 5; i++ {
		fmt.Println(<-c)
	}

}
func boring(msg string) <-chan string {
	c := make(chan string)
	go func() {

		for i := 0; i <= 5; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()

	return c
}
func FanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for v := range input1 {
			c <- v
		}
	}()
	go func() {
		for v := range input2 {
			c <- v
		}

	}()
	return c
}
