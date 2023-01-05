package main

import "fmt"

type person struct {
	first string
	last  string
}
type secret struct {
	person
	ltk bool
}

func (s secret) speak() {
	fmt.Println("I am", s.first, s.last)
}
func (r secret) greet() {
	fmt.Println("Its nice to meet you, My name is", r.first, r.last)
}

// func (r reciever) indetifier(paramters) return(s){  code     }

func main() {
	p1 := secret{
		person: person{
			first: "James",
			last:  "Bond",
		},
		ltk: true,
	}
	p2 := secret{
		person: person{
			first: "Chamar",
			last:  "Babua",
		},
		ltk: true,
	}
	fmt.Println(p1)
	p1.speak()
	p2.speak()
	p1.greet()
}
