package main

import (
	"fmt"
	"sort"
)

func main() {
	//a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	//b := []int{13, 14, 15, 16}
	// a = append(a, b...)
	//sort.Ints(a)

	//fmt.Println("Length of array", len(a))
	//	fmt.Println("index poistion at middle position of array", a[len(a)/2]-1)

	// if len(a)%2 == 1 {
	// 	median := a[len(a)/2]
	// 	fmt.Println(median)
	// } else {
	// 	median := float64((a[len(a)/2]-1)+a[len(a)/2]) / 2
	// 	fmt.Println(median)
	// }
	arr := []int{1, 3}
	arr1 := []int{2}
	a1 := shit(arr, arr1)
	fmt.Println("With func", a1)

}
func shit(num1 []int, num2 []int) float64 {
	num1 = append(num1, num2...)
	sort.Ints(num1)
	var median float64
	if len(num1)%2 == 1 {
		median = float64(num1[len(num1)/2])

	} else {

		median = float64((num1[len(num1)/2])+num1[len(num1)/2]) / 2

	}
	return median

}
