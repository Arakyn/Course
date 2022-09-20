package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println("Context:\t", ctx)
	fmt.Println("Error:   \t", ctx.Err())
	fmt.Printf("Context Type: \t%T\n", ctx)
	fmt.Println("---------------------------")
	ctx1, cancel := context.WithCancel(ctx)
	fmt.Println("Context:\t", ctx1)
	fmt.Println("Error:   \t", ctx1.Err())
	fmt.Printf("Context Type: \t%T\n", ctx1)
	fmt.Println("Cancel :\t", cancel)
	fmt.Printf("Cancel type:\t%T\n\t", cancel)
	fmt.Println("----------------------------------")
	cancel()
	fmt.Println("Context:\t", ctx1)
	fmt.Println("Error:   \t", ctx1.Err())
	fmt.Printf("Context Type: \t%T\n", ctx1)
	fmt.Println("Cancel :\t", cancel)
	fmt.Printf("Cancel type:\t%T\n\t", cancel)
}
