package test

import (
	"testing"

	"github.com/stretchr/testify/require"
	// "github.com/goccy/go-yaml"
)

var a = map[string]struct{}{
	"a": struct{}{},
	"b": struct{}{},
	"c": struct{}{},
	"d": struct{}{},
}

func BenchmarkRange(b *testing.B) {
	var ii int
	for i := 0; i <= b.N; i++ {
		ii = 0
		for range a {
			ii++
		}
	}
	require.Equal(b, 4, ii)
}

func BenchmarkIter(b *testing.B) {
	var ii int
	for i := 0; i <= b.N; i++ {
		ii = 0
		for j := 0; j < len(a); j++ {
			ii++
		}
	}
	require.Equal(b, 4, ii)
}
