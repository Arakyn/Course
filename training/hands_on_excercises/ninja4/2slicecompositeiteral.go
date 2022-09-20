package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x[0] = 1
	x[1] = 10
	x[2] = 100
	x[3] = 1000
	x[4] = 10000
	x[5] = 100000
	x[6] = 1000000
	x[7] = 10000000
	x[8] = 100000000
	x[9] = 1000000000
	for i, v := range x {
		fmt.Println(i, v)

	}
	fmt.Printf("%T\n", x)

}
