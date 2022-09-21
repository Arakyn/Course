package main

import "fmt"

var y = 422

func main() {
	fmt.Println("Hello Word")
	fmt.Printf("%T \n", y)
	fmt.Printf("%o \n", y)
	fmt.Printf("%b \n", y)
	y = 911
	fmt.Printf("%x \n", y)
	fmt.Printf("%d \n", y)
	fmt.Printf("%#x\n%b\n%d\n", y, y, y)
	s := fmt.Sprintf("%#x\t%b\t%d\n", y, y, y)
	fmt.Println(s)
	fmt.Printf("%v \n", y)
}
