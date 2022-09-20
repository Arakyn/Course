package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("CPU:", runtime.NumCPU())
	fmt.Println("ARCH:", runtime.GOARCH)
	fmt.Println("OS:", runtime.GOOS)
	fmt.Scanln("")
}
