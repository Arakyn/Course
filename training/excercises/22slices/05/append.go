package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	y := []int{5, 6, 7, 8}
	fmt.Println(x)
	fmt.Println(y)
	y = append(y, 9, 10, 11)
	x = append(x, y...)
	z := []int{1, 2}
	fmt.Println(x)
	for i := 0; i <= 100; i++ {
		z = append(z, i)
		fmt.Println(z)
	}
}
