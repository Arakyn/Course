package main

import "fmt"

func main() {
	{
		x := foo()
		fmt.Println(x)
	}
	x, y := bar()
	fmt.Println(x, y)
}
func foo() int {
	x := 32
	x = x*x - 2
	return x
}
func bar() (string, int) {
	return "Hello World", 5423232323
}
