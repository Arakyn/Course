package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var jsonBlob = []byte(`[
		{"Name": "Platypus", "Order": "One" },
		{"Name":"Quoll",     "Order": "Two" }
	]`)
	fmt.Println(jsonBlob)

	type Animal struct {
		Name  string
		Order string
	}
	var animals []Animal
	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("Error", err)

	}
	fmt.Printf("%+v\n", animals)

}
