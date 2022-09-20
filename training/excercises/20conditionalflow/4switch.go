package main

import "fmt"

func main() {
	x := "This should print"
	y := "this should'nt print."
	switch {
	case false:
		fmt.Println("Hello Faggot")
	case !false:
		fmt.Println("Hello Faggot2")
	case true:
		fmt.Println("this should print")
	case !true: // this is no default fall through in this code
		fmt.Println("this shouldnt print")
	case (2 == 4):
		fmt.Println(y)
	case (2 == 2):
		fmt.Println(x)
	case !(2 == 2):
		fmt.Println(y)
	case !(2 == 4):
		fmt.Println(x)

	}
}
