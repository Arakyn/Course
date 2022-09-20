package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter your name")
	x, _ := reader.ReadString('\n')

	fmt.Println("Hello", x)

	for x != "festival" {

		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Please enter your name")
		x, _ := reader.ReadString('\n')
		fmt.Println("Hello", x)
		if strings.TrimSpace(x) == "festival" {
			break
		}

	}

}

// breaking loops with input.
