package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var tests []string = []string{"a", "b", "c"}

func test(caller []byte) bool {
	for _, t := range tests {
		if strings.HasPrefix(string(caller), t) {
			return true
		}
	}
	return false
}

func BenchmarkNil(b *testing.B) {
	var r bool
	for i := 0; i <= b.N; i++ {
		r = test([]byte("d"))
	}
	require.False(b, r)
}
