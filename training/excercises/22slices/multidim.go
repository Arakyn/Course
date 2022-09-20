package main

import "fmt"

func main() {
	// multi dimensional slices

	james := []string{"James", " Bond", "Chocolate", "Martini"}
	miss := []string{"Miss", "Penny", "Hazelnnut", "Bubblegum"}

	fmt.Println(len(james))
	fmt.Println(len(miss))
	fmt.Println(james)
	fmt.Println(miss)
	fmt.Println(cap(james))
	fmt.Println(cap(miss))
	xp := [][]string{james, miss}
	fmt.Println(xp)
}
