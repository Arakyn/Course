package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("Hey my name is ,", p.first, p.last, "and i am ", p.age, "years old.")
}
func main() {
	p1 := person{
		first: "Jungngesh",
		last:  "Gaitonde",
		age:   42,
	}
	p2 := person{
		first: "Amandeep",
		last:  "KJai",
		age:   18,
	}
	p1.speak()
	p2.speak()

}
