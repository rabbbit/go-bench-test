package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkOne(b *testing.B) {
	var r bool
	for i := 0; i <= b.N; i++ {
		r = test("a")
	}
	require.True(b, r)
}

func BenchmarkTwo(b *testing.B) {
	var r bool
	for i := 0; i <= b.N; i++ {
		r = test2("a")
	}
	require.True(b, r)
}
