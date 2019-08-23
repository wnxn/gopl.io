package ex5_test

import (
	"github.com/adonovan/gopl.io/ch8/ex5"
	"testing"
)

// sequential 230742875 ns/op
func Benchmark(b *testing.B){
	ex5.CreateNImage(1)
}

// parallel 2818828387 ns/op
func BenchmarkCreateNParallelImage(b *testing.B) {
	ex5.CreateNParallelImage(1)
}