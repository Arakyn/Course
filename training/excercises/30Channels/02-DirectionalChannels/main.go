package main

import "fmt"

func main() {
	//  c := make(chan int, 2) a normal channel it is
	c := make(chan int)
	c1 := make(chan<- int, 2) // a channel that only takes in values

	c2 := make(<-chan int, 3) // this is a recieve only channel it only takes in values
	//   c1 <- 23  // this doesnt work we cant send values to a send only channel

	//  fmt.Println(<-c)
	// fmt.Println(<-c)  // invalid operation: cannot receive from send-only channel c (variable of type chan<- int)compilerInvalidReceive
	// this gives us problem
	fmt.Println("-----------------------------------")
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", c2)
	c1 = c
	fmt.Printf("%T\n", c1)
	fmt.Printf("%T\n", c)
	fmt.Println("_____________________________________________")
	fmt.Printf("c\t%T\n", (chan<- int)(c))

}
