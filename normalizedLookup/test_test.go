package test

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var headers = map[string]string{
	"bla":   "bla",
	"blue":  "blue",
	"hello": "hello",
	"1":     "hello",
	"2":     "hello",
	"3":     "hello",
	"4":     "hello",
}

func normalizedLookup(headers map[string]string, key string) string {
	var value string
	n := len(key)
	for k, v := range headers {
		if len(k) != n {
			continue
		}
		if strings.ToLower(k) == key {
			value = v
			break
		}
	}
	return value
}

func normalizedLookup2(headers map[string]string, key string) string {
	var value string
	for k, v := range headers {
		if strings.ToLower(k) == key {
			value = v
			break
		}
	}
	return value
}

func normalizedLookup3(headers map[string]string, key string) string {
	return headers[strings.ToLower(key)]
}

func Benchmark1(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = normalizedLookup(headers, "hello")
	}
	assert.NotEqual(b, "", r)
}

func Benchmark2(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = normalizedLookup2(headers, "hello")
	}
	assert.NotEqual(b, "", r)
}

func Benchmark3(b *testing.B) {
	var r string
	for i := 0; i < b.N; i++ {
		r = normalizedLookup3(headers, "Hello")
	}
	assert.NotEqual(b, "", r)
}
