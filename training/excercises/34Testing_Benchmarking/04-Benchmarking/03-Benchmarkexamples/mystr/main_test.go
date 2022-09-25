package mystr

import (
	"strings"
	"testing"
)

const s = "We ask ourselves, Who am I to be brilliant, gorgeous, talented, fabulous? Actually, who are you not to be? Your playing small does not serve the world. There is nothing enlightened about shrinking so that other people won't feel insecure around you. We are all meant to shine, as children do. We were born to make manifest the glory that is within us. It's not just in some of us; it's in everyone. And as we let our own light shine, we unconsciously give other people permission to do the same. As we are liberated from our own fear, our presence automatically liberates others. - Marianne Williamson"

func BenchmarkCat(b *testing.B) {
	xs := strings.Split(s, " ")
	for i := 0; i < b.N; i++ {
		Cat(xs)
	}

}
func BenchmarkJooin(b *testing.B) {
	xs := strings.Split(s, " ")
	for i := 0; i < b.N; i++ {
		Join(xs)
	}

}

// $ go test -bench Jooin
// goos: windows
// goarch: amd64
// pkg: github.com/github.com/Arakyn/Course/training/excercises/34Testing_Benchmarking/04-Benchmarking/03-Benchmarkexamples/mystr
// cpu: AMD Ryzen 5 5600X 6-Core Processor
// BenchmarkJooin-12        1753702               687.3 ns/op
// PASS
// ok

// $ go test -bench Cat
// goos: windows
// goarch: amd64
// pkg: github.com/github.com/Arakyn/Course/training/excercises/34Testing_Benchmarking/04-Benchmarking/03-Benchmarkexamples/mystr
// cpu: AMD Ryzen 5 5600X 6-Core Processor
// BenchmarkCat-12        ?????????????????????    99382  ?????????????????            12328 ns/op
// PASS
