package main

import "fmt"

func main() {

	y := []string{"James", "Bond", "Shaken, not Stirred"}
	z := []string{"Miss", "Moneypenny", "Helloooo, James"}
	jb := [][]string{y, z}
	fmt.Println(jb)
	for i, v := range jb {
		fmt.Println(i, v)
		for i, val := range v {
			fmt.Printf("\tIndex Position %v\t Value %v\n", i, val)
		}
	}

}
