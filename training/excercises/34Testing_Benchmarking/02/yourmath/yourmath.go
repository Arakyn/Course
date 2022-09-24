package yourmath

// Add adds shit
func Add(x ...int) int {
	sum := 0

	for _, v := range x {

		sum += v
	}
	return sum
}
