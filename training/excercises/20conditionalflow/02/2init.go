package main

import "fmt"

func main() {
	x := 42
	if x == 42 { // scope of x is lmited
		fmt.Println("-12312-")
	}
	fmt.Println(x)
}
