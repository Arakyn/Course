package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	fmt.Println(ctx)
	fmt.Println(ctx.Err())
	fmt.Printf("%T\n", ctx)
	fmt.Println("---------------------------")

}
