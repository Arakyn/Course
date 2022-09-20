package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x[0] = 42
	x[1] = 43
	x[2] = 44
	x[3] = 45
	x[4] = 46
	x[5] = 47
	x[6] = 48
	x[7] = 49
	x[8] = 50
	x[9] = 51
	for i, v := range x {
		fmt.Println(i, v)

	}
	fmt.Printf("%T\n", x)
	fmt.Println(x[0:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])

}
