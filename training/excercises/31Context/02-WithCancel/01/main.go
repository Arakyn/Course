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
	ctx1, ctx2 := context.WithCancel(ctx)
	fmt.Println(ctx1)
	fmt.Println(ctx.Err())
	fmt.Printf("%T\n", ctx1) // in this we are throwing away the cancel func. which should not be done
	fmt.Println(ctx2)
}
