package frac

import "testing"

func BenchmarkFrac(b *testing.B) {
	for i := 0; i < b.N; i++ {
		VFractal.CreateImage(3, 5)
	}
}
