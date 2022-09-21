package main

import "fmt"

type human interface {
	speak() // basically this means anybody who has the fucntion to speak is also a type of human
	//  a value can be of more than one type

} // types
type person struct {
	first string
	last  string
}
type secret struct {
	person
	ltk bool
}
type hotdog int

// functions for those types
func (s secret) speak() {
	fmt.Println("I am", s.first, s.last)
} // interface ka function
func bar(h human) { // cannot use methods with interfaces probably
	switch h.(type) {
	case person:
		fmt.Println("I am a person", h.(person).first) // this is assertion
	case secret:
		fmt.Println("I am a secret agent", h.(secret).first)
	default:
	}
	fmt.Println("I was passed into bar and i am speaking from bar, I called human", h)
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
	p3 := person{
		first: "Lmao ",
		last:  "Ded",
	}
	fmt.Println(p3)
	fmt.Println(p1)
	p1.speak()
	p2.speak()
	p3.speak()
	bar(p1)
	bar(p2)
	bar(p3)
	// conversion
	var x hotdog = 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	var y int
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
func (p1 person) speak() {
	fmt.Println("I am not a secret agent", p1.first, p1.last)
}
