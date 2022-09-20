package main

import "fmt"

var a int = 42

type hotdog int

var b hotdog

func main() {

	a := 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T \n", a)
	fmt.Printf("%T \n", b)

	// a = b cant do this but lets get around it
	a = int(b)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("%T \n", b)

	b = hotdog(a)
	fmt.Println("changed valei", b)
	fmt.Printf("%T \n", b)

}
