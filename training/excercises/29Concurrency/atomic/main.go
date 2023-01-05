package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	var wg sync.WaitGroup
	var counter int64
	const gs = 101
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			fmt.Println("GoRoutines", runtime.NumGoroutine())
			atomic.AddInt64(&counter, 1)

			runtime.Gosched()
			wg.Done()
		}()

	}
	wg.Wait()
	fmt.Println("CPU", runtime.NumCPU())
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	fmt.Println(atomic.LoadInt64(&counter))
	fmt.Println(counter)
}
