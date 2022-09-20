package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++ {
		if x%4 == 0 {
			fmt.Println("This number ", x, "is fully divided by 4 and there is no remainder")
			continue
		}
		fmt.Println("The remainder of the number", x, "When divided by 4 is", x%4)

	}
}
