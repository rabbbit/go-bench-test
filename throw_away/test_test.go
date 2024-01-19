package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkWrapper(b *testing.B) {
	var out string
	for i := 0; i < b.N; i++ {
		out = func2("load")
	}
	assert.NotEqual(b, "", out)
}
