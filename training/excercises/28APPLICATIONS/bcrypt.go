package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Println(bs)
	loginPword := `password1235`
	err1 := bcrypt.CompareHashAndPassword(bs, []byte(loginPword))
	if err1 != nil {
		fmt.Println("YOU CANT LOGIN")
		return
	}
	fmt.Println("You are logged in")
}
