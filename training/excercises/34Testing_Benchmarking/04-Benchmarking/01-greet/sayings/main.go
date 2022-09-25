package sayings

import "fmt"

func Greet(s string) string {
	x := fmt.Sprint("Hello ", s)
	return x
}
