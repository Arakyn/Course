package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	password := `password12333`
	un := "Aman"
	bs, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MaxCost)
	unbyte := []byte(un)
	_, err := os.Create("dataa")
	f, _ := os.Open("dataa")
	if err != nil {
		fmt.Println("Errorrorororo at file creating.")
	}
	ioutil.WriteFile("dataa", bs, fs.ModeAppend)
	ioutil.WriteFile("dataa", unbyte, fs.ModeAppend)
	read, err1 := ioutil.ReadFile("dataa")
	if err1 != nil {
		fmt.Println("Error while reading,")

	}
	defer f.Close()
	fmt.Println(string(read))
	fmt.Println("Closed")
}
