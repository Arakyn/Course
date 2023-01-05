package main

import "fmt"

func main() {
	x := greetermaker()
	x()
	x1 := greetermaker()
	x1()

}
func greetermaker() func() {
	return func() {
		fmt.Println("Hello World, I was born throug another functon")
	}
}
