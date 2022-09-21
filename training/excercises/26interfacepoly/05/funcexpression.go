package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("This is my first func expression")
	}
	f()
	f1 := func(x int) {
		x = x + 20
		fmt.Println("Added 20 to your number", x)
	}
	f1(23)

}
