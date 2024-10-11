package main

import (
	"maps"
	"slices"
	"sort"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

func add0(in map[string]int) []string {
	r := make([]string, 0, len(in))
	for k := range in {
		r = append(r, k)
	}
	sort.Strings(r)
	return r
}

func add1(in map[string]int) []string {
	r := make([]string, 0, len(in))
	for k := range in {
		r = append(r, k)
	}
	slices.Sort(r)
	return r
}

func add2(in map[string]int) []string {
	return slices.Sorted(maps.Keys(in))
}

func add3(in map[string]int) []string {
	r := make([]string, 0, len(in))
	r = slices.AppendSeq(r, maps.Keys(in))
	slices.Sort(r)
	return r
}

func add4(in map[string]int) []string {
	r := slices.AppendSeq(
		slices.Grow([]string{}, len(in)),
		maps.Keys(in),
	)
	slices.Sort(r)
	return r
}

func prep() map[string]int {
	a := make(map[string]int)
	for i := 0; i < 1000; i++ {
		a["13131322323"+strconv.Itoa(i)] = i
	}
	return a
}

func Benchmark0(b *testing.B) {
	p := prep()
	var r []string
	for i := 0; i <= b.N; i++ {
		r = add0(p)
	}
	require.True(b, len(r) > 100)
}

func Benchmark1(b *testing.B) {
	p := prep()
	var r []string
	for i := 0; i <= b.N; i++ {
		r = add1(p)
	}
	require.True(b, len(r) > 100)
}

func Benchmark2(b *testing.B) {
	p := prep()
	var r []string
	for i := 0; i <= b.N; i++ {
		r = add2(p)
	}
	require.True(b, len(r) > 100)
}

func Benchmark3(b *testing.B) {
	p := prep()
	var r []string
	for i := 0; i <= b.N; i++ {
		r = add3(p)
	}
	require.True(b, len(r) > 100)
}

func Benchmark4(b *testing.B) {
	p := prep()
	var r []string
	for i := 0; i <= b.N; i++ {
		r = add4(p)
	}
	require.True(b, len(r) > 100)
}
