package main

import "fmt"

func main() {
	// to see the memory
	a := 42
	fmt.Println(a)
	fmt.Println(&a)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a)
	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b)
	fmt.Println(*&a)
	*b = 43
	fmt.Println(*b)
	fmt.Println(b)
	fmt.Println(a)

}
