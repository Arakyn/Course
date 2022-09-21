// package math takes in  integers and adds them
package Maths

// func maths takes a int and returns a added value idk mai to itnerne t hui
func Maths(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum

}
