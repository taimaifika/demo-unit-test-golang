package benchmark_test

import (
	math "demo-unit-test/12-benchmark"
	"testing"
)

// Benchmark
func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		math.Add(1, 2)
	}
}
