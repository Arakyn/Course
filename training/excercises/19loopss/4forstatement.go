package main

import "fmt"

func main() {
	x := 1
	for {
		if x > 10 {

			break
		}
		fmt.Println(x)
		x++
	}
	fmt.Println("done")
	fmt.Println("Another Loop")
	for j := 0; j < 100; j = j + 2 {

		if j == 50 {
			continue // it will skip 50

		}
		fmt.Println(j)
	}
	fmt.Println("done 2.")
	for x := 0; x <= 100; x++ {
		if x%2 != 0 {
			continue

		}
		fmt.Println(x)
	}
	fmt.Println("done3.")
}
