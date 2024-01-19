package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkString(b *testing.B) {
	var out string
	var bla [10]byte = [10]byte("hello")
	for i := 0; i <= b.N; i++ {
		out = do(bla)
	}
	require.NotEqual(b, out, "")
}
