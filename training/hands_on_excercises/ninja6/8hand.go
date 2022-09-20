package main

import "fmt"

func main() {
	x := boor()
	fmt.Println(x())
}
func boor() func() int {
	return func() int {
		ii := []int{1, 2, 3, 4, 5, 6}
		total := 0
		for _, v := range ii {
			total += v
		}
		return total
	}
}
