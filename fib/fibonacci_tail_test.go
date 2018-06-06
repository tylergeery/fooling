package fib

import (
	"math/big"
	"testing"
)

func BenchmarkFibTail1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib_tail(i%200, big.NewInt(1), big.NewInt(0))
	}
}

func BenchmarkFibIter1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib_iter(i % 200)
	}
}
