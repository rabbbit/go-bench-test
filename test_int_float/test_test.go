package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkInt(b *testing.B) {
	var out int
	for i := 0; i <= b.N; i++ {
		out = decayI(i)
	}
	require.NotEqual(b, out, "")
}

func BenchmarkFloat(b *testing.B) {
	var out int
	for i := 0; i <= b.N; i++ {
		out = decayF(i)
	}
	require.NotEqual(b, out, "")
}

func TTestMe(t *testing.T) {
	startI := 1000
	startF := 1000

	for i := 0; i < 1000; i++ {
		startI = decayI(startI)
		startF = decayF(startF)
		fmt.Println(startI, startF)
		assert.Equal(t, startI, startF, i)
	}
}
