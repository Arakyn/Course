package main

import "fmt"

var x int

func incrementer() {
	x++
}
func main() {
	incrementer()
	incrementer()
	fmt.Println("x = ", x)
	incrementer()
	fmt.Println("Increased the Value of x by 1")
	fmt.Println(x)
	kb := 1
	fmt.Println("kb = ", kb)
	const (
		_   = iota
		kb1 = 1 << 10
		mb1 = 1 << 20
	)
	fmt.Println("kb1 = ", kb1)
	fmt.Println("mb1 = ", mb1)

}
