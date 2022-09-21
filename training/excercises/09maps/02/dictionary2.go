package main

import "fmt"

func main() {
	var mp map[string]string = map[string]string{

		"pear":   "Green",
		"Apple":  "red",
		"grapes": "purple",
	}

	for fruit, x := range mp {
		fmt.Println(fruit, "=", x)
	}
}
