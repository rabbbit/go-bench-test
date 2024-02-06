package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func add1([]int) int {
	j := 0
	for i := 0; i < 1000; i++ {
		j += i
	}
	return j
}

func add2(...int) int {
	j := 0
	for i := 0; i < 1000; i++ {
		j += i
	}
	return j
}

//go:noinline
func add3(a, b int) int {
	return a + b
}

//go:noinline
func add3a(nums ...int) int {
	return nums[0] + nums[1]
}

func prep() []int {
	a := []int{}
	for i := 0; i < 1000; i++ {
		a = append(a, i)
	}
	return a
}

func BenchmarkSmth(b *testing.B) {
	p := prep()
	var r int
	for i := 0; i <= b.N; i++ {
		r = add1(p)
	}
	require.True(b, r > 1000)
}

func BenchmarkVariadic(b *testing.B) {
	p := prep()
	var r int
	for i := 0; i <= b.N; i++ {
		for j := 0; j < len(p); j++ {
			r = add2(p[j])
		}
	}
	require.True(b, r > 1000)
}

func BenchmarkVariadic2(b *testing.B) {
	p := prep()
	var r int
	for i := 0; i <= b.N; i++ {
		r = add2(p...)
	}
	require.True(b, r > 1000)
}

func BenchmarkZap(b *testing.B) {
	p := prep()
	var r int
	for i := 0; i <= b.N; i++ {
		for j := 0; j < len(p); j += 2 {
			r = add3(p[j], p[j+1])
		}
	}
	require.True(b, r > 1000)
}

func BenchmarkZapVar(b *testing.B) {
	p := prep()
	var r int
	for i := 0; i <= b.N; i++ {
		for j := 0; j < len(p); j += 2 {
			r = add3a(p[j], p[j+1])
		}
	}
	require.True(b, r > 1000)
}

func BenchmarkZapVar2(b *testing.B) {
	p := prep()
	pp := [][]int{}
	for j := 0; j < len(p)-1; j += 2 {
		pp = append(pp, []int{p[j], p[j+1]})
	}
	var r int
	for i := 0; i <= b.N; i++ {
		for j := 0; j < len(pp); j++ {
			r = add3a(pp[j]...)
		}
	}
	require.True(b, r > 1000)
}
