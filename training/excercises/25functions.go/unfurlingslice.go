package main

import "fmt"

func main() {
	// unfurling a slice means opening up a slice
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := foo(xi...) // when you are asking for a variadic parameter you could also pass nothing
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
