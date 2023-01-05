package main

import "fmt"

func main() {

	x := "Helo World"
	fmt.Println(x)
	// this is a closure
	{
		fmt.Println("this was accessed in the inner scoep", x)
		y := "I am a disco dancer"
		fmt.Println(y)
		// the inner scope can access the outer scope
		// the outer scope cannot access the inner scope
		// the scope of y is just insife the inner scope
		// the scope of x inside the whole func main
		// y cannot be accessed the outer scopecd
	}
}
