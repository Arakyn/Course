package main

import "fmt"

var y = 44 // this shit can be used in the whole code
var z int  // declared a variable with identifier Z which is of type int
// not assigning a value will lead to it having a value of Zero of its type

func main() {
	fmt.Println("Hello Worldcxdsss ")
	x := 42 // short term declaration, it can only be used in this func main()
	fmt.Println(x)
	fmt.Println(y)
	foo()

}
func foo() {
	fmt.Println(`Again:`, y)
	x := 21
	fmt.Println("x is a bitch", x)
	yes := make([]int, 6, 6)
	yes1 := append(yes, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(yes)
	fmt.Println(yes1)
	for _, i := range yes1 {
		fmt.Println(i)
	}
}
