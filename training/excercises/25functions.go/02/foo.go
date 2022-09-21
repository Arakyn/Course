package main

import (
	"fmt"
)

var n int
var t int

func main() {
	foo()
	bar("Aman")
	x := bar1(123)
	fmt.Println(x)
	factor(5)
	s1 := woo("Monneypenny")
	fmt.Println(s1)
	x1, y1 := mouse("James", "Amandeep")
	fmt.Println(x1)
	fmt.Println(y1)

}
func foo() {
	fmt.Println("Hello from foo")
} //     s is the variable jo assign hota hai
// string is the indentifier or parameter, jabh argument dalte hai to wo parameter usse identify krta hai then its get assigned to the reciever.

func bar(s string) {
	fmt.Println("Hello,", s)
}
func bar1(l int) int {
	l = l * 5
	return l
}
func factor(x int) int {
	x = x * (x - 1)
	fmt.Println(x)
	return x

}
func woo(st string) string {

	return fmt.Sprint("Hello from woo,", st)
}
func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, "\t", ln, `, says "Hello"`)
	b := false
	return a, b

}
