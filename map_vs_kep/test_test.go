package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type store struct {
	map [string]string
}

func (s *store) Add(map[string]string) {
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
