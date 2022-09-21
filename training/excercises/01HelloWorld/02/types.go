package main

import "fmt"

var y = 23
var z string = `Jame said, "Shaken, not Stirred." ` // declared variable with identifier of Z with type string
//	and i assigned the value of string to it

var ptr = &z

func main() {
	fmt.Println("Hello World")
	fmt.Println(y)
	fmt.Printf("this value is of =  %T \n", y)
	fmt.Printf("this value is of =  %T \n", z)
	fmt.Printf("this value is of =  %T \n", *ptr) // this is pointing to Z
	fmt.Printf("this value is of =  %T \n", ptr)  // this is telling us ptr so its a *string while Z is a string
	fmt.Println(ptr)
	fmt.Println(*ptr)
	foo(ptr)
	fmt.Println(ptr)
	fmt.Println(*ptr)

}
func foo(ptr *string) *string {
	*ptr = "Stirred not shaken asshole"

	return ptr
}
