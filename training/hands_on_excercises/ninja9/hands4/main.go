package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	const gs = 100
	var mu sync.Mutex
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			mu.Lock()
			v := counter
			v++
			counter = v
			fmt.Println(counter)
			fmt.Println("Number of Go Routines:", runtime.NumGoroutine())
			mu.Unlock()
			wg.Done()

		}()

	}
	wg.Wait()
	fmt.Println("Number of Go Routines:", runtime.NumGoroutine())
	fmt.Println(counter)
}
