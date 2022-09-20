package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `pagalkutta1233`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println("Error idiot", err)

	}
	fmt.Println(bs)
	err1 := bcrypt.CompareHashAndPassword(bs, []byte(s))
	if err1 != nil {
		fmt.Println("Password is incorrect")
		return
	}
	fmt.Println("You are logged in.")

}
