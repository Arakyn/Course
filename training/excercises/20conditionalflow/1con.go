package main

import "fmt"

func main() {
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if true {
		fmt.Println("001")
	}
	if false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("001")
	}
	if !false {
		fmt.Println("002")
	}
	if !true {
		fmt.Println("001")
	}
	if !false {
		fmt.Println("002")

	}
	if 2 == 2 {
		fmt.Println("Why not")
	}
	if !(2 == 2) {
		fmt.Println("Why tho not")
	}
	if 2 != 2 {
		fmt.Println("Why lol not")
	}
	if !(2 != 2) {
		fmt.Println("Why lol not")
	}

}
