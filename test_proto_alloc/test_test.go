package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkInt(b *testing.B) {
	var a []int32
	for i := 0; i <= b.N; i++ {
		a = []int32{0, 0, 0, 1}
	}
	require.NotEqual(b, a, "")
}

func BenchmarkIntFull(b *testing.B) {
	var a []int32
	for i := 0; i <= b.N; i++ {
		a = []int32{1, 2, 3, 4}
	}
	require.NotEqual(b, a, "")
}

type test struct {
	index int32
	value int32
}

func BenchmarkStruct(b *testing.B) {
	var a []*test
	for i := 0; i <= b.N; i++ {
		t := &test{3, 1}
		a = []*test{t}
	}
	require.NotEqual(b, a, "")
}

func BenchmarkStructFull(b *testing.B) {
	var a []*test
	for i := 0; i <= b.N; i++ {
		a = []*test{{0, 1}, {1, 1}, {2, 1}, {3, 1}}
	}
	require.NotEqual(b, a, "")
}
