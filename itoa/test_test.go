package test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

const prefix = "pre-"

var prefix2 = []byte("pre-")

func t(i int) string {
	return prefix + strconv.Itoa(i)
}

func BenchmarkNaive(b *testing.B) {
	h := make(map[string]int, 0)
	for i := 0; i <= b.N; i++ {
		h[t(100)] = 10
	}
	assert.NotEqual(b, 0, len(h))
}
