package main

import "fmt"

type person struct {
	First string
	Last  string
}
type human interface {
	speak()
}

func main() {
	p1 := person{"Aman", "Man"}
	saySomething(&p1)

}
func (s *person) speak() {
	fmt.Println("Hello my name is,", s.First, s.Last)
}
func saySomething(h human) {
	h.speak()
}
