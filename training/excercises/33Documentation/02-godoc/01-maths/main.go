package maths

func Sum(x ...int) int {
	sum1 := 0
	for _, v := range x {
		sum1 += v
	}
	return sum1
}
