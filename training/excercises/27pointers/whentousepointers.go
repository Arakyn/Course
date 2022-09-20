package main

import "fmt"

// func main() {
// 	x := 0
// 	foo(x)
// 	fmt.Println(x)
// }
// func foo(y int) {
// 	fmt.Println(y)
// 	y = 43
// 	fmt.Println(y)
// }
func main() {
	x := 0
	fmt.Println("from main before foo", &x)
	fmt.Println("from main before foo", x)
	foo(&x)
	fmt.Println("from main after foo", x)
	fmt.Println("from main after foo", &x)
}
func foo(y *int) {
	fmt.Println("from foo", y)
	fmt.Println("from foo", *y)
	*y = 43
	fmt.Println("from foo", y)
	fmt.Println("from foo", *y)
}
