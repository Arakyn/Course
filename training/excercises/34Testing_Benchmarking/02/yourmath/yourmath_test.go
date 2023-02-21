package yourmath

import (
	"testing"

	mymaths "github.com/Arakyn/Course/training/excercises/33Documentation/02-godoc/01-maths"
)

type test struct {
	data   []int
	answer int
}

func TestSum(t *testing.T) {
	tests := []test{
		{[]int{1, 2}, 3}, {[]int{2, 5}, 7}, {[]int{10, 50}, 60},
		{[]int{-1, 0}, -1},
	}
	for _, v := range tests {
		x := mymaths.Sum(v.data...)
		if x != v.answer {
			t.Error("Expected", v.answer, "got ", x)
		}

	}
	v := mymaths.Sum(2, 3)
	if v != 5 {
		t.Error("Expected 5 got v", v)
	}

}
