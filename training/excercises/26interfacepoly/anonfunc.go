package main

import "fmt"

func main() {
	foo()
	fmt.Println("Hello World")
	func() {
		fmt.Println("Hello from anonymous func")
	}() // anonymous function
	func(x int) {
		fmt.Println(x)
	}(42)
}
func foo() {
	fmt.Println("foo ran")
}
