package main

import "fmt"

func main() {
	xd := factorial(5)
	fmt.Println(xd)
	y := 5
	for i := 1; i < 5; i++ {
		y = y * i
	}
	fmt.Println(y)
	fmt.Println("Factor loop from func factorlrooo[")
	y1 := loopFac(5)
	fmt.Println(y1)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)

}
func loopFac(n int) int {
	total := 1
	for ; n > 0; n-- {
		total *= n
	}
	return total
}
