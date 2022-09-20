package main

import "fmt"

func main() {
	x23 := []string{"James", "Charles", "Babushka"}
	lul(mapp, x23...)
}
func mapp(x []string) {
	for _, v := range x {
		fmt.Println("The Value is  \n", v)
	}
}
func lul(f func(x []string), x ...string) {
	total := 0
	for i, _ := range x {
		total += i

	}
	fmt.Println("The total number of values are these", total)
	f(x)
}
