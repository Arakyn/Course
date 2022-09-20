package main

import "fmt"

func main() {
	s := 132
	fmt.Println(s)
	bs := byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Printf("%b\n", bs)
	fmt.Printf("%x\n", bs)
	fmt.Printf("%d\n", bs)
	fmt.Printf("%o\n", bs)

}
