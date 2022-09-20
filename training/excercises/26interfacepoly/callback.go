package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	kkk := sum(ii...)
	fmt.Println("Sum of all numbers", kkk)
	t := evenSum(sum, ii...)
	fmt.Println(t)
	odd := oddSum(sum, ii...)
	fmt.Println(odd)
}
func sum(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}
func evenSum(f1 func(x ...int) int, vi ...int) int {
	slice := []int{}
	for _, v := range vi {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f1(slice...)
	return total
}
func oddSum(f2 func(x ...int) int, k ...int) int {
	var lol []int
	for _, v := range k {
		if v%2 != 0 {
			lol = append(lol, v)

		}

	}
	return f2(lol...)
}
