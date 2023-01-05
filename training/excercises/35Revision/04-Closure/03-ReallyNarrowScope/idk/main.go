package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []float64{1, 5, 9}
	b := []float64{2, 6, 7}
	a = append(a, b...)
	sort.Float64s(a)
	middle := (a[2] + a[3]) / 2
	fmt.Println(middle)

}
