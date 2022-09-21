package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	var pointer *int = &x[3]
	fmt.Println(*pointer)
	fmt.Println(pointer)
	fmt.Println()
	fmt.Println(x)
	x = append(x[:2], x[4:]...)
	fmt.Println(*pointer)
	fmt.Println(pointer)

	y := []int{1, 2, 3, 4, 5, 6, 7, 78, 9, 90, 12212, 1212, 4456}
	// 0  1  2  3  4  5  6   7  8  9     10         11  12
	y = append(y[:3], y[5:8]...)
	fmt.Println(y)

}
