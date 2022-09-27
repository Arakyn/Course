package word

import (
	"fmt"
	"testing"
)

var xs = `Aman Aman Aman Joker Joker Joker Taliban`

func TestCount(t *testing.T) {

	x := Count(xs)
	if x != 7 {
		t.Error("Wanted 7 got ", x)
	}

}
func TestUseCount(t *testing.T) {

	x := UseCount(xs)

	if x["Aman"] != 3 {
		t.Error("Expected 3 got something else")
	}

}
func BenchmarkCount(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Count(xs)
	}

}
func BenchmarkUsecount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		UseCount(xs)
	}

}
func ExampleCount() {
	xs := `Aman Aman Aman Joker Joker Joker Taliban`
	fmt.Println(Count(xs))
	// Output:
	// 7

}
