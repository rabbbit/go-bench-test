package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type smthWrapper struct {
	s [1000]*something
	m []*something
}

type something struct {
	a    string
	b    string
	c    int
	d    int
	want bool
}

// prepare a slice with 4000 elements where 650
// are interesting.
func prep() []*something {
	var prep []*something
	for i := 0; i < 4000; i++ {
		s := something{}
		if i > 3300 {
			s.want = true
		}
		prep = append(prep, &s)
	}
	return prep
}

func BenchmarkHeap(b *testing.B) {
	prep := prep()
	wrapper := &smthWrapper{}
	wrapper.m = wrapper.s[:0]
	for i := 0; i <= b.N; i++ {
		for _, item := range prep {
			if item.want == true {
				wrapper.m = append(wrapper.m, item)
			}
		}
	}
	require.True(b, len(wrapper.m) > 0)
}

func BenchmarkZero(b *testing.B) {
	prep := prep()
	var mm [1000]*something
	m := mm[:0]
	for i := 0; i <= b.N; i++ {
		for _, item := range prep {
			if item.want == true {
				m = append(m, item)
			}
		}
	}
	require.True(b, len(m) > 0)
}

func BenchmarkOne(b *testing.B) {
	prep := prep()
	var m []*something
	for i := 0; i <= b.N; i++ {
		m = make([]*something, 0, 1000)
		for _, item := range prep {
			if item.want == true {
				m = append(m, item)
			}
		}
	}
	require.True(b, len(m) > 0)
}

func BenchmarkOneA(b *testing.B) {
	prep := prep()
	var m []something
	for i := 0; i <= b.N; i++ {
		m = make([]something, 0, 1000)
		for _, item := range prep {
			if item.want == true {
				m = append(m, *item)
			}
		}
	}
	require.True(b, len(m) > 0)
}

func BenchmarkTwo(b *testing.B) {
	prep := prep()
	var m []*something
	for i := 0; i <= b.N; i++ {
		n := 0
		for _, item := range prep {
			if item.want == true {
				n += 1
			}
		}
		m = make([]*something, 0, n)
		for _, item := range prep {
			if item.want == true {
				m = append(m, item)
			}
		}
	}
	require.True(b, len(m) > 0)
}

func BenchmarkThree(b *testing.B) {
	prep := prep()
	var m []*something
	for i := 0; i <= b.N; i++ {
		m = make([]*something, 0, 0)
		for _, item := range prep {
			if item.want == true {
				m = append(m, item)
			}
		}
	}
	require.True(b, len(m) > 0)
}
