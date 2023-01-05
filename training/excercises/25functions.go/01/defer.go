package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	defer foo()
	defer bar()
	fmt.Println("Hello World 2")
}
func foo() {
	fmt.Println("Foo")

}
func bar() {
	fmt.Println("Bar")
}
