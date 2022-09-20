package main

import "fmt"

type customErr struct {
	err string
}

func main() {
	ce1 := customErr{
		err: "LMAOOOOOOOOOOOOOOOO",
	}
	foo(&ce1)
}
func (e *customErr) Error() string {
	x := fmt.Sprintln("Hey this is the Error of CustomErr somethin g weird happening", e.err)
	return x
}
func foo(e error) {
	fmt.Println("Error", e, "\n", e.(*customErr).err)
}
