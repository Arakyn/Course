package main

import "fmt"

func main() {
	x := []int{11, 22, 33, 44, 56}
	fmt.Println(len(x))
	fmt.Println(x[1])
	fmt.Println(x[1:])
	fmt.Println(x[1:3])
	fmt.Println(x[:4])
	fmt.Println(x[1:3])

	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

}
