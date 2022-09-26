package mymaththree

import (
	"fmt"
	"testing"
)

func TestYears(t *testing.T) {
	x := Years(7)
	if x != 49 {
		t.Error("Expected 49 got", x)
	}

}
func TestYearstwo(t *testing.T) {
	x := YearsTwo(56)
	if x != 8 {
		t.Error("Expected 8 got ", x)
	}

}

func ExampleYears() {
	fmt.Println(Years(8))
	// Output:
	// 56
}
func ExampleYearsTwo() {
	fmt.Println(YearsTwo(63))
	// Output:
	// 9
}
func BenchmarkYears(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Years(7)
	}

}
func BenchmarkYearsTwo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		YearsTwo(63)
	}

}
