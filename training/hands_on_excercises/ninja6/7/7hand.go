package main

import "fmt"

func main() {
	x := func(x ...int) {
		total := 0
		for _, v := range x {
			total += v
		}
		fmt.Println(total)
	}
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8}
	x(ii...)
}
