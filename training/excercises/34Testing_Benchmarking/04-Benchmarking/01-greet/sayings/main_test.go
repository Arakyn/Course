package sayings

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	x := Greet("Aman")
	if x != "Hello Aman" {
		t.Error("lmao", x)

	}

}
func ExampleGreet() {
	fmt.Println(Greet("Aman"))
	// Output:
	// Hello Aman
}
func BenchmarkGreet(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Greet("Aman")
	}
}
