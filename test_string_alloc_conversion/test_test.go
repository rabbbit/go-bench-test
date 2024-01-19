package test

import (
	"bytes"
	"testing"
)

var (
	caller = []byte("hello")
)

func cmp(target []byte) bool {
	return bytes.HasPrefix(caller, target)
}

func cmp2(target string) bool {
	return bytes.HasPrefix(caller, []byte(target))
}

func BenchmarkString(b *testing.B) {
	source := "h"
	for i := 0; i <= b.N; i++ {
		cmp2(source)
	}
}

func BenchmarkByte(b *testing.B) {
	source := []byte("h")
	for i := 0; i <= b.N; i++ {
		cmp(source)
	}
}
