package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	counter := 0
	const gs = 100
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("GoRoutines", runtime.NumGoroutine())

	}
	wg.Wait()

	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	fmt.Println(counter)

}
