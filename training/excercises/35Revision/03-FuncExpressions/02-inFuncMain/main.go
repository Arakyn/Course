package main

import "fmt"

func main() {

	greeting := func() {
		fmt.Println("Hello World, Greetings I am talking from main and was born in func main")
	}
	greeting()
}
