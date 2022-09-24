package main

import (
	"fmt"

	mymaths "github.com/github.com/Arakyn/Course/training/excercises/33Documentation/02-godoc/01-maths"
)

func main() {
	fmt.Println("2 + 2 =", mymaths.Sum(2, 2))
	fmt.Println("3 + 5 =", mymaths.Sum(3, 5))
	fmt.Println("12 + 15 =", mymaths.Sum(12, 15))

}
