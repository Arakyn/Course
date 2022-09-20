package main

import "fmt"

func main() {
	f := 17
	fmt.Printf("%d\t%b\t%#x\n", f, f, f)
	i := f << 1
	fmt.Printf("%d\t%b\t%#x\t", i, i, i)

}
