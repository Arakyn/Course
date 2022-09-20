package main

import "fmt"

func main() {
	x := make(chan int)
	go func() {
		x <- 42
		close(x)
	}()
	fmt.Println("=========")
	v, ok := <-x
	fmt.Println(v, ok)
	fmt.Println(<-x)
	v, ok = <-x
	fmt.Println(v, ok)
}
