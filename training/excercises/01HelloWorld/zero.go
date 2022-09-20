package main

import "fmt"

var y string

func main() {
	fmt.Println("y =", y)
	fmt.Printf(" %T \n", y)
	y = "What are you doing you bich"
	n, err := fmt.Println(y)
	fmt.Println(len(y))
	fmt.Println(n)
	fmt.Println(err)

}
