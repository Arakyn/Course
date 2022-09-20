package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	f, err := os.Open("lmao.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	bs, err1 := ioutil.ReadAll(f)
	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(string(bs))
}
