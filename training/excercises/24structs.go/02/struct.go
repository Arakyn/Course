package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func main() {
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   52,
	}
	p2 := person{
		first: "Amandeep",
		last:  "Khurana",
		age:   17,
	}

	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.last, p2.first, p2.age)
}
