package main

import "fmt"

func main() {
	var x (map[string]int) = map[string]int{

		"Aman": 5,
		"Lmao": 23,
	}
	for i, value := range x {
		fmt.Println(i, "=", value)
	}

}
