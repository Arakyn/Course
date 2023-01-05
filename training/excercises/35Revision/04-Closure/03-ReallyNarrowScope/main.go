package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
func main() {
	incrementer := wrapper()
	fmt.Println("x = ", incrementer())
	fmt.Println("x = ", incrementer())
	x := incrementer()
	x++
	x++
	x++
	x++
	x++
	fmt.Println("x from func main", x)
	fmt.Println("x from incerememrber", incrementer())
}
