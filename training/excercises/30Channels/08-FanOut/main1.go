package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	ye := make(chan int)
	even := make(chan int)

	go send(ye)
	go recieve(even, ye)
	for v := range even {
		fmt.Println("From Even", v)
	}

	fmt.Println("Abou to exit")
}
func send(fo chan<- int) {
	for i := 0; i <= 100; i++ {

		fo <- i
	}
	close(fo)
}
func recieve(e chan int, fo chan int) {
	var wg sync.WaitGroup
	for v := range fo {
		wg.Add(1)
		go func(v int) {
			e <- v
			fmt.Println("Number of GORoutines: ", runtime.NumGoroutine())
			wg.Done()

		}(v)

	}
	wg.Wait()
	close(e)

}
