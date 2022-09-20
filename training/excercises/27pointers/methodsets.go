package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}
type shape interface {
	area() float64
}

func (x *circle) area() float64 {
	return x.radius * x.radius * math.Pi
}
func info(s shape) {
	fmt.Println("Area is circle", s.area())
}
func main() {
	c := circle{
		5,
	}

	info(&c)
	fmt.Printf("%T\n", &c)
	fmt.Println(c.area())
}
