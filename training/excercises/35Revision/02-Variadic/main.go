package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Total numbers", adder(slice))
	fmt.Println("Total numbers", unfurleradder(slice...))
	// unfurling the data using ... and passing it to the function
	fmt.Println("Average of your numbers", average(slice1...))

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
func unfurleradder(x ...int) int {
	adder := 0
	for _, v := range x {
		adder += v
	}
	return adder
}

// ... is called a variadic parameter. this x is taking in data one by one and storing it into a slice.
func average(x ...float64) float64 {
	total := 0.0
	for _, v := range x {
		total += v
	}
	return total / float64(len(x))
}
