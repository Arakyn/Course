package main

import "fmt"

func main() {

	ii := []int{1, 2, 3, 4, 5, 6, 7, 7, 8}
	x := foo(ii...)
	fmt.Println(x)
	x1 := bar(ii)
	fmt.Println(x1)
}
func foo(x1 ...int) int {
	total := 0
	for _, v := range x1 {
		total += v

	}

	return total
}
func bar(x []int) int {
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
