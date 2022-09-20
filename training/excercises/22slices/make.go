package main

import "fmt"

func main() {
	x := make([]int, 10, 14)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	fmt.Println("")
	x[9] = 1

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
	x = append(x, 50)
	x = append(x, 50)
	x = append(x, 50)
	x = append(x, 50)
	x = append(x, 50)

	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))
}
