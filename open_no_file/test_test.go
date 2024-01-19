package test

import (
	"os"
	"testing"
)

func BenchmarkOpen(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		os.Open("/i/dont/exist")
	}
}

func BenchmarkStat(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		os.Stat("/i/dont/exist")
	}
}
