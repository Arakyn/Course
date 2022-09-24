// Package acdc takes in a slice of integers and does very creepy things to them.
package acdc

// Sum takes in unlimited integers and returns the sum of them cause he is a bitch.go get -u golang.org/x/lint/golint
func Sum(n ...int) int {

	s := 0
	for _, v := range n {
		s += v
	}
	return s
}
