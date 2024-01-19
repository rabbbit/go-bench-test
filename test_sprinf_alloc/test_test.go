package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkSprintf(b *testing.B) {
	var out string
	for i := 0; i <= b.N; i++ {
		out = concat4(i)
	}
	require.NotEqual(b, out, "")
}

func BenchmarkBuilder(b *testing.B) {
	var out string
	for i := 0; i <= b.N; i++ {
		out = concat(i)
	}
	require.NotEqual(b, out, "")
}

func BenchmarkBuilderPool(b *testing.B) {
	var out string
	for i := 0; i <= b.N; i++ {
		out = concat3(i)
	}
	require.NotEqual(b, out, "")
}

func BenchmarkConcat(b *testing.B) {
	var out string
	for i := 0; i <= b.N; i++ {
		out = concat2(i)
	}
	require.NotEqual(b, out, "")
}
