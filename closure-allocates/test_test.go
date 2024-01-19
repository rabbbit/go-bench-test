package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkClosure(b *testing.B) {
	b.ReportAllocs()

	x := &outer{
		in: []*inner{
			{a: 1},
			{a: 2},
			{a: 3},
		},
	}

	w := &worker{
		in: &inner{},
	}
	for i := 0; i < b.N; i++ {
		assert.NotEqual(b, 0, run(x, w))
	}
}
