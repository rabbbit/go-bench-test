package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const _alphanums = "bcdfghjklmnpqrstvwxz2456789"

const _idLength = 5
const l = len(_alphanums)

const max = 27 * 27 * 27 * 27

var max2 = [5]int{1, 27, 27 * 27, 27 * 27 * 27, 27 * 27 * 27 * 27}

const (
	m0 = 27 * 27 * 27 * 27
	m1 = 27 * 27 * 27
	m2 = 27 * 27
	m3 = 27
)

func func1(index int) string {
	base := 1
	id := make([]byte, _idLength)
	for i := 0; i < _idLength; i++ {
		rIndex := (index / base) % len(_alphanums)
		id[_idLength-i-1] = _alphanums[rIndex]
		base *= len(_alphanums)
	}
	return string(id)
}

func func2(index int) string {
	base := 1
	id := make([]byte, _idLength)
	for i := 0; i < _idLength; i++ {
		rIndex := (index / base) % l
		id[_idLength-i-1] = _alphanums[rIndex]
		base *= l
	}
	return string(id)
}

// max = 531441
// 8765748

func func3(index int) string {
	id := make([]byte, _idLength)
	base := 1
	for i := 0; i < _idLength; i++ {
		n := index / base
		rIndex := n - len(_alphanums)*(n/len(_alphanums))
		id[_idLength-i-1] = _alphanums[rIndex]
		base *= len(_alphanums)
	}
	return string(id)
}

func func4(index int) string {
	id := make([]byte, _idLength)

	for i := 5; i >= 1; i-- {
		curr := max2[i-1]
		n := index / curr
		id[5-i] = _alphanums[n]
		index = index - n*curr
	}
	return string(id)
}

func func5(index int) string {
	id := make([]byte, _idLength)

	for i := 5; i >= 1; i-- {
		curr := max2[i-1]
		id[5-i] = _alphanums[index/curr]
		index = index % curr
	}
	return string(id)
}

func func7(index int) string {
	id := make([]byte, _idLength)

	idx := index / m0
	id[0] = _alphanums[idx]
	index = index - m0*idx

	idx = index / m1
	id[1] = _alphanums[idx]
	index = index - m1*idx

	idx = index / m2
	id[2] = _alphanums[idx]
	index = index - m2*idx

	idx = index / m3
	id[3] = _alphanums[idx]
	index = index - m3*idx

	id[4] = _alphanums[index]

	return string(id)
}

func func6(index int) string {
	id := make([]byte, _idLength)

	id[0] = _alphanums[index/m0]
	index = index % m0

	id[1] = _alphanums[index/m1]
	index = index % m1

	id[2] = _alphanums[index/m2]
	index = index % m2

	id[3] = _alphanums[index/m3]
	index = index % m3

	id[4] = _alphanums[index]

	return string(id)
}

func BenchmarkUnroll2(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = func7(8765748)
	}
	assert.Equal(b, "vrmmm", r)
}

func BenchmarkUnroll(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = func6(8765748)
	}
	assert.Equal(b, "vrmmm", r)
}

func BenchmarkReverse3(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = func5(8765748)
	}
	assert.Equal(b, "vrmmm", r)
}

func BenchmarkReverse2(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = func4(8765748)
	}
	assert.Equal(b, "vrmmm", r)
}

func BenchmarkModulo(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = func1(8765748)
	}
	assert.Equal(b, "vrmmm", r)
}

func BenchmarkModulo2(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = func2(8765748)
	}
	assert.Equal(b, "vrmmm", r)
}

func BenchmarkReverse(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = func3(8765748)
	}
	assert.Equal(b, "vrmmm", r)
}
