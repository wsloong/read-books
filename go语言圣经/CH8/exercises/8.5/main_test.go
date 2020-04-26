package main

import "testing"


func BenchmarkManderbrotSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ManderbrotSingle()
	}
}

func BenchmarkManderbrotWithGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ManderbrotWithGoroutine()
	}
}