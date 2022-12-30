package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Total numbers", adder(slice))

}
func adder(x []int) int {
	fmt.Println(x)

	sum := 0
	for i, v := range x {
		if i <= 4 {
			continue
		}

		sum += v
		fmt.Println("For item in index position", i, "we are now adding", v, "to the total which is now", sum, "")
	}
	return sum
}
