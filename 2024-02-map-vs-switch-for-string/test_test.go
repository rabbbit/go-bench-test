package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var m = map[string]int{
	"0":                    1,
	"something-longer":     1,
	"something-very-large": 1,
	"something-longer1":    1,
	"something-longer2":    1,
	"something-longer3":    1,
	"something-longer5":    1,
	"something-longer6":    1,
	"something-longer7":    1,
	"something-longer8":    1,
	"something-longer9":    1,
	"something-longer10":   1,
	"something-longer11":   1,
	"something-longer12":   1,
}

func func1(key string) (int, bool) {
	v, ok := m[key]
	return v, ok
}

func func2(key string) (int, bool) {
	switch key {
	case "0":
		return 1, true
	case "something-longer":
		return 1, true
	case "something-very-large":
		return 1, true
	case "something-longer1":
		return 1, true
	case "something-longer2":
		return 1, true
	case "something-longer3":
		return 1, true
	case "something-longer4":
		return 1, true
	case "something-longer5":
		return 1, true
	case "something-longer6":
		return 1, true
	case "something-longer7":
		return 1, true
	case "something-longer8":
		return 1, true
	case "something-longer9":
		return 1, true
	case "something-longer10":
		return 1, true
	case "something-longer11":
		return 1, true
	case "something-longer12":
		return 1, true
	}
	return 0, false
}

func BenchmarkMapHit(b *testing.B) {
	var r int
	var ok bool
	for i := 0; i <= b.N; i++ {
		r, ok = func1("something-very-large")
	}
	require.Equal(b, r, 1)
	require.True(b, ok)
}

func BenchmarkMapMiss(b *testing.B) {
	var r int
	var ok bool
	for i := 0; i <= b.N; i++ {
		r, ok = func1("something-very-large-missing")
	}
	require.Equal(b, r, 0)
	require.False(b, ok)
}

func BenchmarkSwitchHit(b *testing.B) {
	var r int
	var ok bool
	for i := 0; i <= b.N; i++ {
		r, ok = func2("something-very-large")
	}
	require.Equal(b, r, 1)
	require.True(b, ok)
}

func BenchmarkSwitchMiss(b *testing.B) {
	var r int
	var ok bool
	for i := 0; i <= b.N; i++ {
		r, ok = func2("something-very-large-missing")
	}
	require.Equal(b, r, 0)
	require.False(b, ok)
}
