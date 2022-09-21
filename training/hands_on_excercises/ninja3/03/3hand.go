package main

import "fmt" // jedi level 3 excercise part 1

func main() {
	for x := 0; x <= 10000; x++ {
		fmt.Println(x)
	}
	fmt.Println("done.")
	y := 0
	for y <= 10000 {
		fmt.Println(y)
		y++
	}
	fmt.Println("also done.")
}
