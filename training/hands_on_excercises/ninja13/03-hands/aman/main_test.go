package aman

import "testing"

type data struct {
	data   []int
	answer float64
}

func TestCentered(t *testing.T) {
	d1 := data{[]int{1, 2, 3, 4}, 2.5}
	d2 := data{[]int{10, 20, 30, 40}, 25}
	d3 := data{[]int{5, 5, 5, 5}, 5}
	d4 := data{[]int{0, 1, 1, 34534}, 1}
	datatable := []data{d1, d2, d3, d4}
	for _, v := range datatable {
		x := CenteredAvg(v.data)
		if x != v.answer {
			t.Error("Expected", v.answer, "Got", x)
		}
	}

}
func BenchmarkCenteredAvg(b *testing.B) {
	d1 := []int{1, 2, 3, 4}
	for i := 0; i < b.N; i++ {
		CenteredAvg(d1)
	}

}
