package main

import (
	"errors"
	"fmt"
	"log"
)

var errorlul = errors.New("Cant squarroot a negative number you retard")

func main() {
	fmt.Printf("%T\n", errorlul)
	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}

}
func sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errorlul
	}
	return 42, nil
}
