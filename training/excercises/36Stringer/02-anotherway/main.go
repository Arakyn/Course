package main

import (
	"fmt"
	"strings"
)

type movie struct {
	moviename string
	year      int
	cast      caststruct
}
type caststruct []actors
type actors struct {
	actorname string
	role      string
}

func (m movie) String() string {
	return fmt.Sprintf("\t%v\n Made in (%v) \n %v", m.moviename, m.year, m.cast)
}
func (cc caststruct) String() string {
	var caststrings []string
	for _, castmembers := range cc {
		caststrings = append(caststrings, castmembers.String())

	}
	cast := strings.Join(caststrings, ",\n ")
	return cast
}
func (a actors) String() string {
	return fmt.Sprintf("%v as %v", a.actorname, a.role)
}

func main() {
	m1 := movie{
		moviename: "the Matrix",
		year:      1999,
		cast: caststruct{
			{actorname: "Keanu Reeves", role: "Neo"},
			{actorname: "Laurence Fishburne", role: "Morpheus"},
		},
	}

	x := m1.cast.String()
	x1 := m1.String()
	fmt.Println(x)
	fmt.Println(x1)

}
