package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type seceretAgent struct {
	person
	ltk   bool
	first int
}

func main() {
	sa1 := seceretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   42,
		},
		ltk:   true,
		first: 2004,
	}
	fmt.Println(sa1)
	fmt.Println(sa1.person.first, sa1.first)

}
