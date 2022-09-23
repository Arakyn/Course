package main

import "fmt"

func main() {
	fmt.Println("Average for this bitch is", Average(1, 2, 3, 4))
}
func Average(x ...int) float64 {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return float64(sum) / float64(len(x))
}
