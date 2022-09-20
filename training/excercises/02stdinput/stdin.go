package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Println("Welcome to User Input")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please Rate our Pizza")
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating our pizza", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error")

	} else {
		fmt.Println("Added 1 to your rating", numRating*5)
	}
	newreader := bufio.NewReader(os.Stdin)

	fmt.Println("What is your name")
	input1, _ := newreader.ReadString('\n')
	fmt.Println("Hi", input1)

}
