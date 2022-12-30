package main

import (
	"fmt"

	"github.com/github.com/Arakyn/Course/training/excercises/35Revision/doggo"
)

func main() {
	x1 := foo("James Bond")
	x1(42)
	doggo.Speak()

}

// function returning a func
func foo(x string) func(x int) {
	fmt.Println("Hellooo", x)
	return func(x int) {
		println("Printed yor number fool", x)
	}

}
