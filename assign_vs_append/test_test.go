package test

import (
	"testing"
)

var result []int

const size = 32

func doAssign() {
	data := make([]int, size)
	for i := 0; i < size; i++ {
		data[i] = i
	}
	result = data
}

func doAppend() {
	data := make([]int, 0, size)
	for i := 0; i < size; i++ {
		data = append(data, i)
	}
	result = data
}

func BenchmarkAssign(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doAssign()
	}
}

func BenchmarkAppend(b *testing.B) {
	for i := 0; i < b.N; i++ {
		doAppend()
	}
}
