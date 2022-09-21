package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 8, 23, 453, 1, 1, 1} // composite literal x := types{ data with commas} this is a slice. not a array
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Printf("%T\n", x)
	a := [3]string{"Hello"}
	fmt.Println(a)
	fmt.Println(len(a))
}

// a slice allows you to group VALUES of the SAME TYPE.
