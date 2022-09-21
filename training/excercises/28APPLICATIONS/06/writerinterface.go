package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Hello World")
	fmt.Fprintln(os.Stdout, "Hello World")
	io.WriteString(os.Stdout, "Hello Habbibibiib\n")
	io.WriteString(os.Stdout, "MFKER HELLO SAMAJ AA GAYA?????????")
}
