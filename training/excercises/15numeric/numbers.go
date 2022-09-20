package main

import (
	"fmt"
	"runtime"
)

func main() {
	x := 42
	y := 42.3456
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\t%T\n", x, y)
	var u int = 12
	fmt.Println(u)
	fmt.Printf("%T\n", u)
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
