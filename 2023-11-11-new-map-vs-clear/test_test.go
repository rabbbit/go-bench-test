package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkOne(b *testing.B) {
	var m map[string]struct{}
	for i := 0; i <= b.N; i++ {
		m = make(map[string]struct{}, 100)
		for j := 0; j < 100; j++ {
			m[strconv.Itoa(j)] = struct{}{}
		}
	}
	require.True(b, len(m) > 0)
}

func BenchmarkTwo(b *testing.B) {
	m := make(map[string]struct{})
	for i := 0; i <= b.N; i++ {
		clear(m)
		for j := 0; j < 100; j++ {
			m[strconv.Itoa(j)] = struct{}{}
		}
	}
	require.True(b, len(m) > 0)
}
