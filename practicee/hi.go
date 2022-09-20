package main

import (
	"fmt"
)

var name string
var prefix string

func main() {
	prefix := "Hello"
	fmt.Print("Enter your name pls:")
	fmt.Scan(&name)
	fmt.Print("\n Welcome to the sim\n")
	x := greet(prefix, name)
	fmt.Println(x)
}
func greet(s string, s1 string) string {
	fmt.Println(s, " ", s1)
	var a string
	return a

}
