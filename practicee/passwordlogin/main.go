package main

import (
	"fmt"
)

var sliceee []map[string]string
var slicecatch []map[string]string

func main() {
	var val int

	for i := 0; i < 120; i-- {
		fmt.Println("Welcome to the Login Machine")
		fmt.Println("Please Select a Option")
		fmt.Println("1. Signup        2. Login           3. Exit")
		fmt.Println("---------------------------------------------")
		fmt.Scan(&val)
		if val == 1 {
			slicecatch = signup(sliceee)
		} else if val == 2 {
			login(slicecatch)
		} else if val == 3 {
			break
		} else {
			fmt.Println("Enter a Valid Option")
		}

	}
}
func signup(slice []map[string]string) []map[string]string {
	fmt.Println("Please Enter Your Username.")
	var x string
	fmt.Scan(&x)
	var password string

	fmt.Println("Please Enter your password")
	fmt.Scan(&password)

	user := map[string]string{
		x: password,
	}
	fmt.Println("Signing you up")
	slice1 := append(slice, user)
	return slice1

}
func login(slice []map[string]string) {
	fmt.Println("Logging you up")
	fmt.Println("Please Enter your Username")
	var usernamelogin string
	var passwordlogin string
	fmt.Scan(&usernamelogin)
	fmt.Println("Please Enter your password")
	fmt.Scan(&passwordlogin)

	fmt.Println("Doing the work")
	for _, v := range slice {
		for i, _ := range v {
			if v[i] == passwordlogin {
				fmt.Println("You are logged in.")
			} else {
				fmt.Println("Password is incorrect.")
			}

		}
	}

}
