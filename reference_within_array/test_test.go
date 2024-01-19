package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Benchmark1(b *testing.B) {
	var t tstruct
	for i := 0; i < b.N; i++ {
		t = test1()
	}
	assert.NotEqual(b, t, tstruct{})
}

func BenchmarkAppend(b *testing.B) {
	var t tstruct
	for i := 0; i < b.N; i++ {
		t = test1()
	}
	assert.NotEqual(b, t, tstruct{})
}

func BenchmarkPointer(b *testing.B) {
	var t *tstruct
	for i := 0; i < b.N; i++ {
		t = testPointer()
	}
	assert.NotEqual(b, *t, tstruct{})
}

func BenchmarkPointer2(b *testing.B) {
	var t *tstruct
	for i := 0; i < b.N; i++ {
		t = testPointer2()
	}
	assert.NotEqual(b, *t, tstruct{})
}

func BenchmarkPointer3(b *testing.B) {
	var t *tstruct
	for i := 0; i < b.N; i++ {
		t = testPointer3()
	}
	assert.NotEqual(b, *t, tstruct{})
}

func BenchmarkPointer4(b *testing.B) {
	var t *tstruct
	for i := 0; i < b.N; i++ {
		t = testPointer4()
	}
	assert.NotEqual(b, *t, tstruct{})
}
