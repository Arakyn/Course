package main

import "fmt"

func main() {
	s := `Hello Playground 世界`
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Printf("%T\n", bs)
	fmt.Println(bs)
	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U\n", s[i])

	}

	for i, v := range s {
		fmt.Printf("At point index %d we have %#U\n ", i, v)
	}

}
