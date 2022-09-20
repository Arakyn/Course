package main

import "fmt"

var x int = 1

func main() {
	for i := 0; i < 10; i++ {

		fmt.Println(x)

		x = x + x
	}
	for i := 1; i <= 10; i++ {
		fmt.Println(i * 5)
	}

}
