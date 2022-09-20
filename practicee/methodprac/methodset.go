package main

import (
	"fmt"
	"math"
)

type square struct {
	sides float64
}
type circle struct {
	radius float64
}
type shapes interface {
	area() float64
}

func areaprint(s shapes) {
	fmt.Println("Area of Shape is", s.area())
}

func main() {
	s1 := square{5}
	circle2 := circle{14}
	areaprint(&s1)
	areaprint(&circle2)
	areaprint(&s1)

}
func (s *square) area() float64 {
	sides := s.sides * s.sides
	return sides
}
func (r *circle) area() float64 {
	radius := math.Pi * r.radius * r.radius
	return radius
}
