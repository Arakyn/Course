package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	odd := make(chan int)
	yess := make(chan int)

	go send(even, odd)
	go FanIn(even, odd, yess)
	for v := range yess {
		fmt.Println(v)
	}
	fmt.Println("About to exit")

}
func send(e, odd chan<- int) {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			odd <- i
		}
	}
	close(e)
	close(odd)

}
func FanIn(e, odd <-chan int, fi chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range e {
			fi <- v
		}
		wg.Done()
	}()
	go func() {
		for v := range odd {
			fi <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(fi)

}
