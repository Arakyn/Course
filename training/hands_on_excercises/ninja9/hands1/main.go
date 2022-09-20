package main

import (
	"fmt"
	"sync"
)

func foo() {
	fmt.Println("Hello World from foo.")
}
func bar() {
	fmt.Println("Hello from some another world.")
}
func yay() {
	fmt.Println("I am speaking from yay.")
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		foo()
		wg.Done()
	}()
	go func() {
		bar()
		wg.Done()
	}()
	wg.Wait()
	yay()

}
