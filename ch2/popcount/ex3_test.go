package popcount

import (
	"testing"
)

func TestPopCount(t *testing.T) {
	if PopCount(0xff) != PopCountLoop(0xff) {
		t.Fatalf("should equal: single %v, loop %v", PopCount(0xff), PopCountLoop(0xff))
	}
}

func BenchmarkPopCount(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCount(uint64(i))
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 1; i < b.N; i++ {
		PopCountLoop(uint64(i))
	}
}
