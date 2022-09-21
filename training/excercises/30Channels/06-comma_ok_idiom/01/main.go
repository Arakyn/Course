package main

import "fmt"

func main() {

	even := make(chan int)
	oddd := make(chan int)
	quit := make(chan bool)
	go send(even, oddd, quit)
	fmt.Println("_____________________________________")
	recieve(even, oddd, quit)
	fmt.Println("About to exit./")

}
func send(e, o chan<- int, q chan<- bool) {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}

	}

	q <- true
	q <- false
	close(e)
	close(o)
	close(q)

}

func recieve(e, o <-chan int, q <-chan bool) {
	for {
		select {
		case v := <-e:
			fmt.Println("From The Even Channel:", v)
		case v := <-o:
			fmt.Println("From The Oddd Channel:", v)
		case i, ok := <-q:
			if ok {
				fmt.Println("From ok comma,", i, ok)

			} else {
				fmt.Println("From Okay comma two.", i, ok)

			}
			return

		}

	}
}
