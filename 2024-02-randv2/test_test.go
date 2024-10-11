package main

import (
	"math/rand"
	randv2 "math/rand/v2"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkV1(b *testing.B) {
	var instanceID uint64 = 123131
	r := randv2.New(randv2.NewPCG(instanceID, instanceID+1))

	var a uint32
	for i := 0; i <= b.N; i++ {
		a = r.Uint32()
	}
	require.NotEqual(b, -1, a)
}

func BenchmarkV2(b *testing.B) {
	var instanceID int64 = 123131
	r := rand.New(rand.NewSource(instanceID))

	var a uint32
	for i := 0; i <= b.N; i++ {
		a = r.Uint32()
	}
	require.NotEqual(b, -1, a)
}

func BenchmarkV2Global(b *testing.B) {
	var a uint32
	for i := 0; i <= b.N; i++ {
		a = randv2.Uint32()
	}
	require.NotEqual(b, -1, a)
}

func BenchmarkV2IntN(b *testing.B) {
	var a int
	for i := 0; i <= b.N; i++ {
		a = int(randv2.IntN(120))
	}
	require.NotEqual(b, -1, a)
}

func twoInt32(n int32) (int32, int32) {
	i := randv2.Int32N(n)
	j := i + 1 + randv2.Int32N(n-1)
	if j >= n {
		j -= n
	}
	return i, j
}

func BenchmarkV2TwoInt(b *testing.B) {
	var a, c int32
	for i := 0; i <= b.N; i++ {
		a, c = twoInt32(120)
	}
	require.NotEqual(b, -1, a)
	require.NotEqual(b, -1, c)
}
