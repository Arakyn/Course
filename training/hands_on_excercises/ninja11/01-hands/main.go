package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First   string
	Last    string
	Sayings []string
}

func main() {
	p1 := person{
		First:   "James",
		Last:    "Bond",
		Sayings: []string{"Shaken, not stirred", "Any last wishes?", "Never say never"},
	}
	p2 := person{
		First:   "Aman",
		Last:    "KASH",
		Sayings: []string{"WE HAWTT", "WTF JENNIFFER", "OMG"},
	}
	people := []person{p1, p2}

	bs, err := json.Marshal(people)
	if err != nil {
		fmt.Println("error while marshaling", err)
	}
	fmt.Println(string(bs))

}
