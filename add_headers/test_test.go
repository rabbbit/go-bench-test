package test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkBaseline(b *testing.B) {
	out := make(map[string]string, 1)
	for i := 0; i < b.N; i++ {
		out["me"] = "hello"
	}
	assert.NotEqual(b, "", out["me"])
}

func BenchmarkStatic1(b *testing.B) {
	out := make(map[string]string, 1)
	for i := 0; i < b.N; i++ {
		do1(out, "me", "hello")
	}
	assert.NotEqual(b, "", out["me"])
}

func BenchmarkStatic2(b *testing.B) {
	out := make(map[string]string, 1)
	for i := 0; i < b.N; i++ {
		do2(out, "me", "hello")
	}
	assert.NotEqual(b, "", out["me"])
}

func BenchmarkStaticMore(b *testing.B) {
	out := make(map[string]string, 1)
	d := doer{}
	for i := 0; i < b.N; i++ {
		d.doMore(out)
	}
	assert.NotEqual(b, "", out["me"])
}

func BenchmarkDynamic1(b *testing.B) {
	out := make(map[string]string, 1)
	for i := 0; i < b.N; i++ {
		do1(out, "me", strconv.Itoa(i))
	}
	assert.NotEqual(b, "", out["me"])
}

func BenchmarkDynamic2(b *testing.B) {
	out := make(map[string]string, 1)
	for i := 0; i < b.N; i++ {
		do2(out, "me", strconv.Itoa(i))
	}
	assert.NotEqual(b, "", out["me"])
}
