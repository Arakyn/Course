package main

import "fmt"

func main() {
	const a = 42
	const b = 43.3223
	const c = "James Bond"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

}
