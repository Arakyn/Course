package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	c := make(chan int)
	c1 := make(chan int)
	populate(c)
	go FanOutIn(c, c1)
	for v := range c1 {
		fmt.Println(v)
	}
	fmt.Println("About to exit.")
}
func populate(c chan int) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for i := 0; i <= 100; i++ {
			c <- i
		}
		wg.Done()
		wg.Wait()
		close(c)
	}()
}
func FanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup

	for v := range c1 {
		wg.Add(1)
		c2 <- v
		wg.Done()
	}
	fmt.Println(runtime.NumGoroutine())

	wg.Wait()
	close(c2)

}
