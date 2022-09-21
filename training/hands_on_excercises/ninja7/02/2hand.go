package main

import "fmt"

type person struct {
	name string
}

func main() {
	p1 := person{
		"Aman",
	}
	changeMe(&p1)
	fmt.Println("from main", p1.name)

}
func changeMe(x *person) {
	fmt.Printf("%T\n", x)
	fmt.Println(x)
	(*x).name = "Lmao name changed to faggot"
	fmt.Println(x.name)
}
