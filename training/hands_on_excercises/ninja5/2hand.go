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
	x := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}
	fmt.Println(x[p1.last])
	fmt.Println(x[p2.last])
	for k, v := range x {
		fmt.Println("This is the key", k)
		for _, v1 := range v.favFlav {
			fmt.Println("This is the favourite flavours:", v1)
		}
	}
}
