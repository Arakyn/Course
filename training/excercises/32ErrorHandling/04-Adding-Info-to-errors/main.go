package main

import (
	"errors"
	"log"
)

// very nice guys
func main() {
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
	// hello fags
}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("Square root of a negative number cant do")
	}
	return 42, nil
}
