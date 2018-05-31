package main

import "testing"

func BenchmarkFibTail1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib_tail(i%200, 1, 0)
	}
}

// func BenchmarkFibRecursive1000(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		fib_recursive(i % 100)
// 	}
// }

func BenchmarkFibIter1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fib_iter(i % 200)
	}
}
