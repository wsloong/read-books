package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(10)
	}
}

func BenchmarkPopCountWithLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithLoop(10)
	}
}

func BenchmarkPopCountWithBitMove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithBitMove(10)
	}
}

func BenchmarkPopCountWithCleanLowest(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountWithCleanLowest(10)
	}
}
