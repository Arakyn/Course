package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t", runtime.GOOS)
	fmt.Println("ARCH\t", runtime.GOARCH)
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Add(1) // ek cheez hai uske liye wait krna before to end the program
	go foo()
	boo()
	fmt.Println("CPU\t", runtime.NumCPU())
	fmt.Println("GoRoutines\t", runtime.NumGoroutine())
	wg.Wait() // jaise hi khatam hoga wait khatam
}
func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo :", i)
	}
	wg.Done()
	// after the loop ends wg.Done() is checked ki ho gaya and now the waiting stops and the program exits lmao ded mai bina dekhe keyboard chala rahe lesgo

}
func boo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Boo: ", i)
	}
}
