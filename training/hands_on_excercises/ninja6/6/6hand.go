package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 56, 6, 7, 7, 68, 7, 4, 534, 23, 543, 5}
	func(x ...int) {
		total := 0
		for _, v := range x {
			total += v
		}
		fmt.Println(total)
	}(ii...)

}
