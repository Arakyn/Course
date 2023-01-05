package main

import "fmt"

type movie struct {
	movie  string
	year   int
	actors []actor
}
type actor struct {
	name string
	role string
}

func (m movie) string() string {
	return fmt.Sprintf("%v (%v) \n %v", m.movie, m.year, m.actors)
}
func (a actor) string() string {
	return fmt.Sprintf("%v as %v", a.name, a.role)
}
func main() {

	m1 := movie{
		movie: "the Matrix",
		year:  1999,
		actors: []actor{

			{name: "Keanu Reeves", role: "Neo"},
			{name: "Laurence Fishburne", role: "Morpheus"},
		},
	}
	x, x1 := m1.string(), m1.actors[0].string()
	fmt.Println(x)
	fmt.Println(x1)

}
