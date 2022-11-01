package main

import "fmt"

func main() {

	s1 := reverseword("Cat")
	fmt.Println(s1)
}
func reverseword(word string) string {
	j := 0
	r := []byte(word)
	var r1 = make([]byte, len(r))

	for i := len(r) - 1; i >= 0; i-- {

		fmt.Println(r)
		fmt.Println(len(r))
		fmt.Println("Value of j", j)
		fmt.Println(i)
		r1 = append(r1, r[i])

	}

	return string(r1)
}
