package main

import "fmt"

func main() {
	f()
	fmt.Println("Returned normally from f.")
}
func f() {
	defer func() {

		if r := recover(); r != nil {
			fmt.Println("Recovered in f.", r)

		}
	}()
	fmt.Println("Calling G.")
	g(0)
	fmt.Println("Sucessfully returned from G.")
}
func g(x int) {
	if x > 3 {
		fmt.Println("Panncikkkcing")
		panic(fmt.Sprintf("%v", x))
	}
	defer fmt.Println("Defer in g", x)
	fmt.Println("printing in g .", x)
	g(x + 1)
}
