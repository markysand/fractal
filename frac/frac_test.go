package frac

import "testing"

// 3043531 ns/op
// 2838700 ns/op
// 2819896 ns/op

func BenchmarkFrac(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VFractal.CreateImage(3, 5)
	}
}
