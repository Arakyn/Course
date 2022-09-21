package main

import "fmt"

func main() {
	switch {
	case 2 == 3:
		fmt.Println("this should not print")
	default:
		fmt.Println("This should print as its default")
	}
}
