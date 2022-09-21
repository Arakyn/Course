package main

import "fmt"

func main() {
	for x := 33; x <= 122; x++ {
		fmt.Printf("This is the ascii symbol for: %d\t is : %#U\n", x, x)
	}

}
