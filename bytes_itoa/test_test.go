package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = [10]byte{}

func func1(id int) string {
	in := input
	in[0] = 'a'
	return string(in[:])
}

func func2(id int) string {
	in := make([]byte, 10)
	in[0] = 'a'
	return string(in[:])
}

func func3(id int) string {
	in := [10]byte{}
	in[0] = 'a'
	return string(in[:])
}

func func4(out []byte) string {
	out[3] = 'a'
	return string(out)
}

func Benchmark1(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = func1(1)
	}
	assert.NotEqual(b, 0, r)
}

func Benchmark2(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = func2(1)
	}
	assert.NotEqual(b, 0, r)
}

func Benchmark3(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = func3(1)
	}
	assert.NotEqual(b, 0, r)
}

func Benchmark4(b *testing.B) {
	in := make([]byte, 10)
	var r string
	for i := 0; i < b.N; i++ {
		r = func4(in)
	}
	assert.NotEqual(b, 0, r)
}
