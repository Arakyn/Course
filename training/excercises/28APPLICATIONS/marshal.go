package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{"Aman", "Singh", 41}
	p2 := person{"Miss", "Moneypenny", 31}
	people := []person{p1, p2}
	fmt.Println(people)
	b, err := json.Marshal(people)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(string(b))

}
