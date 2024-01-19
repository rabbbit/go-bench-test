package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type data struct {
	a string
	b string
	c int
}

func func1(in data) data {
	in.a = "a"
	in.b = "b"
	in.c = 123
	return in
}

func entry() data {
	d := data{}
	return func1(d)
}

func func1Alt(in *data) *data {
	in.a = "a"
	in.b = "b"
	in.c = 123
	return in
}

func entryAlt() *data {
	d := &data{}
	d = func1Alt(d)
	return d
}

func func1Alt2(in *data) {
	in.a = "a"
	in.b = "b"
	in.c = 123
}

func entryAlt2() *data {
	d := &data{}
	func1Alt(d)
	return d
}

func BenchmarkCopy(b *testing.B) {
	var d data
	for i := 0; i < b.N; i++ {
		d = entry()
	}
	assert.Equal(b, data{"a", "b", 123}, d)
}

func BenchmarkRef(b *testing.B) {
	var d *data
	for i := 0; i < b.N; i++ {
		d = entryAlt()
	}
	assert.Equal(b, data{"a", "b", 123}, *d)
}

func BenchmarkRefNoReturn(b *testing.B) {
	var d *data
	for i := 0; i < b.N; i++ {
		d = entryAlt2()
	}
	assert.Equal(b, data{"a", "b", 123}, *d)
}
