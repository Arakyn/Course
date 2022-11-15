package main

import "fmt"

func main() {
	a := "Hellop"
	switch {
	case a == "Hello":
		fmt.Println("Sup guy")
	case a != "Hellop":
		fmt.Println("You are a Hello")
	default:
		fmt.Println("I am from Delhi")
	}
	fmt.Println("Enter your name")
	fmt.Scan(&a)
	fmt.Println("Hi", a)

}
