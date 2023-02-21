package main

import (
	"fmt"

	"github.com/Course/training/hands_on_excercises/ninja12/dog"
)

func main() {
	var x int
	fmt.Println("Lets convert the ages")
	fmt.Scanln(&x)
	fmt.Println("Your Age converted into Doc Years. : ", dog.Convert(x))

}
