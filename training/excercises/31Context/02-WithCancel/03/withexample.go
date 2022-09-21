package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	fmt.Println("Checking Error 1:", ctx.Err())
	fmt.Println("Number of Go Routines: ", runtime.NumGoroutine())
	go func() {
		n := 0
		for {
			select {
			case <-ctx.Done():
				return
			default:
				n++
				time.Sleep(time.Millisecond * 200)
				fmt.Println("Working on it", n)
			}
		}
	}()
	time.Sleep(time.Second * 3)
	fmt.Println("------------------------------------------------------")
	fmt.Println("Error Check 2:", ctx.Err())
	fmt.Println("Go Routune check 2:", runtime.NumGoroutine())
	fmt.Println("About to cancel context")
	fmt.Println("--------------------------------------------------------")
	cancel()
	time.Sleep(time.Second * 3)
	fmt.Println("Cancelled Context")
	fmt.Println("Error Check 3:", ctx.Err())
	fmt.Println("Go Routune check 3:", runtime.NumGoroutine())

}
