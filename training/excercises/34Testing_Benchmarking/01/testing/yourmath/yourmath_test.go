package yourmath

import "testing"

func TestAverage(t *testing.T) {

	v := Add(1, 2, 4, 5)
	if v != 12 {
		t.Error("Expected 3 got", v)
	}
}
