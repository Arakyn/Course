package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, err := os.Create("lmao.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	r := strings.NewReader("Wasssupp")
	io.Copy(f, r)
}
