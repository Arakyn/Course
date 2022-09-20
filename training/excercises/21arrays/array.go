package main

import "fmt"

func main() {
	x := [4]int{1, 2, 3, 4}
	fmt.Println(x)
	fmt.Println("done.")
	var y [5]int
	fmt.Println(y)
	y[4] = 42
	fmt.Println(y)
	fmt.Println(len(y))

}
