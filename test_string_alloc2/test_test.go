package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type stringer interface {
	String() string
}

type a struct {
	s string
}

func (a a) String() string { return a.s }

type bs struct {
	b []byte
}

func (b bs) String() string { return string(b.b) }

func BenchmarkA(b *testing.B) {
	r := ""
	var aa a
	for i := 0; i <= b.N; i++ {
		aa = a{s: "a"}
		r = aa.String()
	}
	assert.NotEqual(b, "1", r)
}
func BenchmarkAIface(b *testing.B) {
	r := ""
	var aa stringer
	for i := 0; i <= b.N; i++ {
		aa = &a{s: "a"}
		r = aa.String()
	}
	assert.NotEqual(b, "1", r)
}

func BenchmarkB(b *testing.B) {
	r := ""
	var aa bs
	for i := 0; i <= b.N; i++ {
		aa = bs{b: []byte("bb")}
		r = aa.String()
	}
	assert.NotEqual(b, "1", r)
}

func BenchmarkBIface(b *testing.B) {
	r := ""
	var aa stringer
	for i := 0; i <= b.N; i++ {
		aa = bs{b: []byte("bb")}
		r = aa.String()
	}
	assert.NotEqual(b, "1", r)
}
