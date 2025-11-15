package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type keyT struct {
	a, b, c string
}

func BenchmarkLookup(b *testing.B) {
	key1 := []byte("abcde")
	key2 := "caller-idsdgsdfgsdgsdgfsdgfsdgfsdfgsdgs"
	key3 := "stagingdfgdgdgd"
	b.Run("struct", func(b *testing.B) {
		cache := map[keyT]bool{}
		for b.Loop() {
			key := keyT{a: string(key1), b: key2, c: key3}
			_, ok := cache[key]
			if !ok {
				cache[key] = true
			}
		}
		assert.True(b, len(cache) > 0)
	})
	b.Run("conca", func(b *testing.B) {
		cache := map[string]bool{}
		for b.Loop() {
			key := string(key1) + ":" + key2 + ":" + key3
			_, ok := cache[key]
			if !ok {
				cache[key] = true
			}
		}
	})
	b.Run("conca2", func(b *testing.B) {
		cache := map[string]bool{}
		for b.Loop() {
			key := string(key1) + key2 + key3
			_, ok := cache[key]
			if !ok {
				cache[key] = true
			}
		}
	})
	b.Run("struct", func(b *testing.B) {
		cache := map[keyT]bool{}
		for b.Loop() {
			_, ok := cache[keyT{a: string(key1), b: key2, c: key3}]
			if !ok {
				cache[keyT{a: string(key1), b: key2, c: key3}] = true
			}
		}
	})
	b.Run("conca", func(b *testing.B) {
		cache := map[string]bool{}
		for b.Loop() {
			_, ok := cache[string(key1)+":"+key2+":"+key3]
			if !ok {
				cache[string(key1)+":"+key2+":"+key3] = true
			}
		}
	})
	b.Run("conca2", func(b *testing.B) {
		cache := map[string]bool{}
		for b.Loop() {
			_, ok := cache[string(key1)+key2+key3]
			if !ok {
				cache[string(key1)+key2+key3] = true
			}
		}
	})
}
