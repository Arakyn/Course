package main

import "fmt"

// var x int
// .\closure.go:17:2: undefined: x ye ayega result

func main() {

	var x int
	fmt.Println(x)
	x++
	{
		var y int
		y = 32
		fmt.Println(y)
		fmt.Println("Inside a function main ")
	}
	// fmt.Println(y  )     # command-line-arguments
	// .\closure.go:19:14: undefined: y            cant do this also cause of the scoop of y is narrowed down to those curly braces
	fmt.Println(x)
	// foo()
	fmt.Println(x)
	zo := incermenter()
	fmt.Println(zo())
	fmt.Println(zo())
	fmt.Println(zo())
	fmt.Println(zo())
	fmt.Println(zo())
	fmt.Println(zo())
	fmt.Println(zo())
	fmt.Println(zo())
	fmt.Println(zo())
	fmt.Println(zo())
	fmt.Println(zo())

}

// func foo() {
// 	fmt.Println("Hello")
// 	//x++ cant do this now
// }

func incermenter() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
