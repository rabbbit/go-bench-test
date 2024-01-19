package test

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type valAlias string

type store struct {
	v int
}

func update(s *store, v valAlias) {
	s.v, _ = strconv.Atoi(string(v))
}

type valStruct struct {
	v string
}

func updateStruct(s *store, vs valStruct) {
	s.v, _ = strconv.Atoi(vs.v)
}

type valStruct2 struct {
	v string
	b []byte
}

func updateStruct2(s *store, vs valStruct2) {
	s.v, _ = strconv.Atoi(vs.v)
}

func BenchmarkCurrent(b *testing.B) {
	s := &store{}
	for i := 0; i <= b.N; i++ {
		v := valAlias("3")
		update(s, v)
	}
	assert.Equal(b, 3, s.v)
}

func BenchmarkStruct(b *testing.B) {
	s := &store{}
	for i := 0; i <= b.N; i++ {
		v := valStruct{v: "3"}
		updateStruct(s, v)
	}
	assert.Equal(b, 3, s.v)
}

func BenchmarkStruct2(b *testing.B) {
	s := &store{}
	for i := 0; i <= b.N; i++ {
		v := valStruct2{v: "3"}
		updateStruct2(s, v)
	}
	assert.Equal(b, 3, s.v)
}
