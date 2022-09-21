package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	defer foo()
	bar()
}
func foo() {
	fmt.Println("Foo")

}
func bar() {
	fmt.Println("Bar")
}
