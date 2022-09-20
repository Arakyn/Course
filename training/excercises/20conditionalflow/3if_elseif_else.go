package main

import "fmt"

func main() {
	x := 45
	if x == 40 {
		fmt.Println("our value was 40")
	} else if x == 41 {
		fmt.Println("our value was 41")

	} else if x == 42 {
		fmt.Println("our value was 42")
	} else if x == 43 {
		fmt.Println("Our value was 43")

	} else if x == 44 {
		fmt.Println("Our Value was 44")
	} else {
		fmt.Println("The value was 40,41,42,43 or 44")
	}

}
