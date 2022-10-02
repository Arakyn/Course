package main

import "fmt"

func main() {
	var val int
	fmt.Println("Welcome to the Login Machine")
	fmt.Println("Please Select a Option")
	fmt.Println("1. Login        2. Sign Up           3. Exit")
	fmt.Println("---------------------------------------------")
	fmt.Scanln(&val)
	switch {
	case val == 1:
	default:
		fmt.Println("Please Enter a Valid Value")

	}
}
func login() {
	fmt.Println("Please Enter your Username")
}
