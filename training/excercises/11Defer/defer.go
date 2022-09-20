package main

import "fmt"

func main() {

	//  FirstRun()
	// SecondRun()
	// out put is I ran second
	//I ran First
	defer FirstRun()
	SecondRun()
	// out put is
	// I ran second
	// I ran First
}
func FirstRun() {

	fmt.Println("I ran First")

}
func SecondRun() {

	fmt.Println("I ran second")
}
