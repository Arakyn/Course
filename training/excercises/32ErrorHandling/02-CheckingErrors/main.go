package main

import "fmt"

func main() {
	var answer1, answer2, answer3 string
	fmt.Println("Your Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Fav Food: ")
	_, err1 := fmt.Scan(&answer2)
	if err1 != nil {
		fmt.Println(err)
	}
	fmt.Println("Fav Sport: ")
	_, err2 := fmt.Scan(&answer3)
	if err2 != nil {
		fmt.Println(err)
	}
	fmt.Println(answer1, answer2, answer3)

}
