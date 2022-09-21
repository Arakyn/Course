package main

import "fmt"

type vehicle struct {
	doors  int
	colour string
}
type truck struct {
	vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	toyota := truck{
		vehicle: vehicle{
			doors:  4,
			colour: "Yellow",
		},
		fourWheel: true,
	}

	jaguar := sedan{
		vehicle: vehicle{
			doors:  2,
			colour: "black",
		},
		luxury: true,
	}
	fmt.Println(toyota)
	fmt.Println("Is toyota a 4 door", toyota.fourWheel)
	fmt.Println(jaguar)
	fmt.Println("is jaguar luxury", jaguar.luxury)
}
