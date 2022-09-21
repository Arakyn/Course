package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Add your number")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	fmt.Println("This is your first number", input)
	fmt.Println("Add your second number")
	reader1 := bufio.NewReader(os.Stdin)
	input1, _ := reader1.ReadString('\n')
	fmt.Println("This your second number", input1)
	x, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 64)
	y, _ := strconv.ParseInt(strings.TrimSpace(input1), 0, 64)
	fmt.Println("remainder is =", x%y)
	fmt.Println("Lets check if your numbers are completely divisible ")

	if x%y == 0 {
		fmt.Println("Number is compleely divisible")
	} else {
		fmt.Println("this number is not completely disivble")
	}

	fmt.Printf("The type of var is %T \n", x)
	const hi float64 = 3.141564532432423
	fmt.Println("Value of pi is,", hi)
}
