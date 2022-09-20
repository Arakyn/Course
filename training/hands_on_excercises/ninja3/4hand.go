package main

import "fmt"

func main() {

	for y := 65; y <= 90; y++ {
		fmt.Println(y)
		for x := 0; x < 3; x++ {
			fmt.Printf("\t%#U\n", y)

		}
		fmt.Printf("\n")
	}
}
