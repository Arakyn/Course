package main

import "fmt"

type person struct {
	first   string
	last    string
	favFlav []string
}

func main() {

	p1 := person{
		first:   "Amandeep",
		last:    "Khurana",
		favFlav: []string{"Pineapple", "strawberry"},
	}
	p2 := person{
		first:   "Divesh",
		last:    "Ahuja",
		favFlav: []string{"Chocolate", "Bubblegum"},
	}
	fmt.Println(p1.first, p1.first)
	fmt.Println(p2.first, p2.last)

	for _, v := range p1.favFlav {
		fmt.Println(v)
	}
	for _, v := range p2.favFlav {
		fmt.Println(v)
	}
}
