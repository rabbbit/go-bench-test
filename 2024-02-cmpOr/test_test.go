package main

import (
	"cmp"
	"testing"

	"github.com/stretchr/testify/require"
)

type myType struct {
	a string
	b string
	c string
}

var (
	target = myType{
		a: "hello",
		b: "worldd",
		c: "yes",
	}
	target2 = myType{
		a: "helloo",
		b: "world",
		c: "yesss",
	}
)

func cmp1(a, b myType) int {
	if a.a != b.a {
		return cmp.Compare(a.a, b.a)
	}
	if a.a != b.a {
		return cmp.Compare(a.b, b.b)
	}
	return cmp.Compare(a.c, b.c)
}

func cmp2(a, b myType) int {
	return cmp.Or(
		cmp.Compare(a.a, b.a),
		cmp.Compare(a.b, b.b),
		cmp.Compare(a.c, b.c),
	)
}

func cmp3(a, b myType) int {
	if one := cmp.Compare(a.a, b.a); one != 0 {
		return one
	}
	if one := cmp.Compare(a.b, b.b); one != 0 {
		return one
	}
	return cmp.Compare(a.c, b.c)
}

func BenchmarkManual(b *testing.B) {
	var r int
	for i := 0; i <= b.N; i++ {
		r = cmp1(target, target2)
	}
	require.Equal(b, -1, r)
}

func BenchmarkManual2(b *testing.B) {
	var r int
	for i := 0; i <= b.N; i++ {
		r = cmp1(target2, target)
	}
	require.Equal(b, 1, r)
}

func BenchmarkManual3(b *testing.B) {
	var r int
	for i := 0; i <= b.N; i++ {
		r = cmp1(target, target)
	}
	require.Equal(b, 0, r)
}

func BenchmarkDheerendra(b *testing.B) {
	var r int
	for i := 0; i <= b.N; i++ {
		r = cmp3(target, target2)
	}
	require.Equal(b, -1, r)
}

func BenchmarkDheerendra2(b *testing.B) {
	var r int
	for i := 0; i <= b.N; i++ {
		r = cmp3(target2, target)
	}
	require.Equal(b, 1, r)
}

func BenchmarkDheerendra3(b *testing.B) {
	var r int
	for i := 0; i <= b.N; i++ {
		r = cmp3(target, target)
	}
	require.Equal(b, 0, r)
}

func BenchmarkOr(b *testing.B) {
	var r int
	for i := 0; i <= b.N; i++ {
		r = cmp2(target, target2)
	}
	require.Equal(b, -1, r)
}

func BenchmarkOr2(b *testing.B) {
	var r int
	for i := 0; i <= b.N; i++ {
		r = cmp2(target2, target)
	}
	require.Equal(b, 1, r)
}

func BenchmarkOr3(b *testing.B) {
	var r int
	for i := 0; i <= b.N; i++ {
		r = cmp2(target, target)
	}
	require.Equal(b, 0, r)
}
