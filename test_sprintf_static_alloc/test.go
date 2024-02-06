package test

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

const (
	q      = "q"
	format = "q=%d"
)

func concat(i int) string {
	var sb strings.Builder
	sb.Grow(10)
	sb.WriteString(q)
	sb.WriteString(strconv.Itoa(i))
	return sb.String()
}

func concat2(i int) string {
	return q + strconv.Itoa(i)
}

var bufPool = sync.Pool{
	New: func() interface{} {
		// The Pool's New function should generally only return pointer
		// types, since a pointer can be put into the return interface
		// value without an allocation:
		return new(strings.Builder)
	},
}

func concat3(i int) string {
	b := bufPool.Get().(*strings.Builder)
	b.Reset()
	b.Grow(10)
	b.WriteString(q)
	b.WriteString(strconv.Itoa(i))

	s := b.String()
	bufPool.Put(b)
	return s
}

func concat4(i int) string {
	return fmt.Sprintf(format, i)
}

func RollingConcat() {
	a := []string{"a", "b", "c", "d"}
	var out string
	for _, a := range a {
		out = "hello " + a
	}
	if out == "" {
		panic("hi")
	}
}

func RollingSprintf() {
	var out string
	a := []string{"a", "b", "c", "d"}
	for _, a := range a {
		out = fmt.Sprintf("%s %s", "hello", a)
	}
	if out == "" {
		panic("hi")
	}
}
