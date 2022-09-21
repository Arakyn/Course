package main

import "fmt"

func main() {
	x := foo(1, 2, 3, 4, 5)
	fmt.Println(x)
}
func foo(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	sum := 0
	for _, v := range x {
		sum += v

	}
	fmt.Println(sum)
	return sum
}
