package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to Divison")
	fmt.Println("Add your first number")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	x, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 64)
	fmt.Println("Add your second number")

	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	y, _ := strconv.ParseInt(strings.TrimSpace(input1), 0, 64)
	if (x % y) != 0 {
		fmt.Println("mfker put decent numbers")

	} else {
		quo := x % y
		fmt.Println("Here is remainder", quo)
	}
	reader3 := bufio.NewReader(os.Stdin)
	input4, _ := reader3.ReadString('\n')
	fmt.Println(input4)
}
