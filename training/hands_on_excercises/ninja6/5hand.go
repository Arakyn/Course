package main

import (
	"fmt"
	"math"
)

type square struct {
	sides float64
}
type circle struct {
	radii float64
}
type shape interface {
	areaCalc() float64
}

func main() {
	s1 := square{
		sides: 5,
	}
	circle2 := circle{
		radii: 14.5,
	}
	info(s1)
	info(circle2)

}
func (s square) areaCalc() float64 {

	return s.sides * s.sides
}
func (c circle) areaCalc() float64 {

	return c.radii * c.radii * math.Pi

}
func info(sh shape) {
	fmt.Println(sh.areaCalc())
}
