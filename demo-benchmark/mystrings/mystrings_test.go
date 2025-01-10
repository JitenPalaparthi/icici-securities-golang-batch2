package mystrings_test

import (
	"demo/mystrings"
	"testing"
)

// func BenchmarkSomething(b *testing.B) {

// }

func BenchmarkConcatenateUsingPlusOperator(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		mystrings.ConcatenateUsingPlusOperator(1000)
	}
}

func BenchmarkConcatenateUsingBuilder(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		mystrings.ConcatenateUsingBuilder(1000)
	}
}
