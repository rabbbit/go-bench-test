package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func ptrToString(s *string, def string) string {
	if s != nil {
		return *s
	}
	return def
}

func BenchmarkNil(b *testing.B) {
	var a *string
	var s string
	for i := 0; i <= b.N; i++ {
		s = ptrToString(a, "")
	}
	require.Empty(b, s)
}

func BenchmarkNonNil(b *testing.B) {
	input := "baba"

	var a *string = &input
	var s string
	for i := 0; i <= b.N; i++ {
		s = ptrToString(a, "")
	}
	require.NotEmpty(b, s)
}
