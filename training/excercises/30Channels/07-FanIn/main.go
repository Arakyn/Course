package main

import (
	"fmt"
	"sync"
)

func main() {
	even := make(chan int)
	oddd := make(chan int)
	fanin := make(chan int)
	go send(even, oddd)
	go recieve(even, oddd, fanin)
	for v := range fanin {
		fmt.Println(v)
	}
}

// send  channel
func send(even, odd chan<- int) {

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			even <- i

		} else {
			odd <- i
		}
	}

	close(odd)
	close(even)
}

// recieve channel
func recieve(even, oddd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		for v := range even {
			fanin <- v
		}
		wg.Done()
	}()

	go func() {
		for v1 := range oddd {
			fanin <- v1

		}
		wg.Done()
	}()

	wg.Wait()
	close(fanin)
}
