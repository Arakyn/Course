package main

import "fmt"

func main() {
	fmt.Printf("this is %v\n", true && true)
	fmt.Printf("this is %v\n", true && false)
	fmt.Printf("this is %v\n", true || true)
	fmt.Printf("this is %v\n", true || false)
	fmt.Printf("this is %v\n", !true)

}
