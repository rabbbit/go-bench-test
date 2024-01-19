package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type (
	err1 struct{}
	err2 struct{}
	err3 struct{}
)

func (err1) Error() string { return "err1" }
func (err2) Error() string { return "err2" }
func (err3) Error() string { return "err3" }

func assert1(e error) int {
	switch e.(type) {
	case err1:
		return 1
	case err2:
		return 2
	case err3:
		return 3
	}
	return 4
}

type retBehaviour uint8

const (
	ret1 retBehaviour = iota
	ret2
	ret3
)

func assert2(r retBehaviour) int {
	switch r {
	case ret1:
		return 1
	case ret2:
		return 2
	case ret3:
		return 3
	}
	return 4
}

func BenchmarkType(b *testing.B) {
	err := err3{}
	var ret int
	for i := 0; i <= b.N; i++ {
		ret = assert1(err)
	}
	assert.Equal(b, 3, ret)
}

func BenchmarkEnum(b *testing.B) {
	r := retBehaviour(2)
	var ret int
	for i := 0; i <= b.N; i++ {
		ret = assert2(r)
	}
	assert.Equal(b, 3, ret)
}
