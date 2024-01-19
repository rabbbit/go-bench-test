package test

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	q0 = "q=0"
	q1 = "q=1"
	q2 = "q=2"
	q3 = "q=3"
	q4 = "q=4"
	q5 = "q=5"
)

var (
	bq0 = []byte(q0)
	bq1 = []byte(q1)
	bq2 = []byte(q2)
	bq3 = []byte(q3)
	bq4 = []byte(q4)
	bq5 = []byte(q5)

	m = map[string]wrapper{
		q0: wrapper(q0),
		q1: wrapper(q1),
		q2: wrapper(q2),
		q3: wrapper(q3),
		q4: wrapper(q4),
		q5: wrapper(q5),
	}
)

type wrapper string

func mapped(a []byte) wrapper {
	if r, ok := m[string(a)]; ok {
		return r
	}
	return wrapper(a)
}

func naive(a []byte) wrapper {
	return wrapper(a)
}

func interned(a []byte) wrapper {
	switch {
	case bytes.Equal(a, bq0):
		return wrapper(q0)
	case bytes.Equal(a, bq1):
		return wrapper(q1)
	case bytes.Equal(a, bq2):
		return wrapper(q2)
	case bytes.Equal(a, bq3):
		return wrapper(q3)
	case bytes.Equal(a, bq4):
		return wrapper(q4)
	case bytes.Equal(a, bq5):
		return wrapper(q5)
	}
	return wrapper(a)
}

var (
	inputEmpty = []byte{}
	input0     = []byte("q=0")
	input1     = []byte("q=1")
	input2     = []byte("q=2")
	input3     = []byte("q=3")
	input4     = []byte("q=4")
	input5     = []byte("q=5")
	input6     = []byte("q=6")
	input11    = []byte("q=11")

	inputs = [][]byte{
		inputEmpty,
		input0,
		input1,
		input2,
		input3,
		input4,
		input5,
		input6,
		input11,
	}
)

func BenchmarkNaive(b *testing.B) {
	var r wrapper

	type benchCase struct {
		name string
		f    func([]byte) wrapper
	}

	var benchCases = []benchCase{
		{"naive", naive},
		{"interned", interned},
		{"mapped", mapped},
	}

	for _, input := range inputs {
		b.Run(string(input), func(b *testing.B) {
			for _, bc := range benchCases {
				b.Run(bc.name, func(b *testing.B) {
					for i := 0; i <= b.N; i++ {
						r = bc.f(input)
					}
				})
			}
		})
	}
	assert.NotEqual(b, "1", r)
}
