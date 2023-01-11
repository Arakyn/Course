package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println(start)
	for i := 0; i < 100000000; i++ {

	}
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println(elapsed)

	fmt.Println(time.Now().Format("01-02-2006 15:04:05"))
	presentTime := time.Now().Format("Monday 01-02-2006 15:04:05 Jan MST")
	fmt.Println(presentTime)
	createdDate := time.Date(2004, time.December, 31, 12, 23, 34, 0, time.UTC)
	fmt.Println(createdDate.Format("Monday 01-02-06 Jan MST 15:04:05"))
}
