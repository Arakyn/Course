package main

import (
	"fmt"

	"github.com/github.com/Arakyn/Course/training/hands_on_excercises/ninja13/02-hands/quote"
	"github.com/github.com/Arakyn/Course/training/hands_on_excercises/ninja13/02-hands/word"
)

func main() {
	fmt.Println(word.Count(quote.SunAlso))
	for k, v := range word.UseCount(quote.SunAlso) {
		fmt.Println(v, k)
	}

}
