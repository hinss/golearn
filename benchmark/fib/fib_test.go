package fib_test

import (
	"learngo/benchmark/fib"
	"testing"
)

func BenchmarkFib20(b *testing.B) {
	for n:=0; n < b.N; n++ {
		fib.Fib(20)
	}
}
