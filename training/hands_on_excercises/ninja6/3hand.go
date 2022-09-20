package main

import "fmt"

func main() {
	x := func() {
		fmt.Println("I am speaking from this package main")
	}
	defer x()
	foo()
	bar()
}

func foo() {
	fmt.Println("I ran first")
}
func bar() {
	fmt.Println(" I Rand second")
}
