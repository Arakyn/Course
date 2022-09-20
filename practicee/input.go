package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")
	reader := bufio.NewReader(os.Stdin)
	inpu, _ := reader.ReadString('\n')
	fmt.Println(inpu)

	x := map[int]string{
		1: "AMan",
		2: "Lmaoo",
	}
	fmt.Println(x[1])

}
