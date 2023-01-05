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

}
