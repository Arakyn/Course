// this is package Maths and does mathematical operations
package Mathss

// Func Averages takes in any number of numbers and gives average
func Averages(x ...float64) float64 {

	var sum float64
	for _, v := range x {
		sum += v
	}
	return sum / float64(len(x))
}
