package main

import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
		if i%2 != 0 {
			continue
		}
		if i == 50 {
			break
		}
		fmt.Println(i)
	}
}
