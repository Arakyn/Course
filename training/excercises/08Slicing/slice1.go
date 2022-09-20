package main

import "fmt"

func main() {

	even := [5]int{0, 2, 4, 6, 8}
	for _, value := range even {
		fmt.Println(value)
	}

	// slicing an array
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	sliced := slice[:3]
	fmt.Println(sliced)
	fmt.Println(slice)
	for _, value1 := range slice {
		fmt.Println(value1)
	}

	slice2 := make([]int, 5, 6)
	fmt.Println(slice2)
	copy(slice2, sliced)
	fmt.Println(slice2)

	slice3 := append(even[:], 3, 4, 5)
	for _, value3 := range slice3 {
		fmt.Print(value3)
	}

}
