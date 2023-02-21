package main

import (
	"fmt"

	"github.com/arakyn/Course/training/hands_on_excercises/ninja13/01-hands/mymaththree"
)

type canine struct {
	name string
	age  int
}

func main() {

	fido := canine{
		name: "Jacky",
		age:  7,
	}
	fmt.Println("Age of Jacky in human years")
	fmt.Println(mymaththree.Years(fido.age))

	fmt.Println("My age is 63, if i was a dog my age would be", mymaththree.YearsTwo(63))

}
