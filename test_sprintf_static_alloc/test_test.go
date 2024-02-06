package test

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func BenchmarkBase(b *testing.B) {
	var out string
	for i := 0; i <= b.N; i++ {
		out = "hello bye"
	}
	require.Equal(b, "hello bye", out)
}

func BenchmarkConcat(b *testing.B) {
	var out string
	for i := 0; i <= b.N; i++ {
		out = "hello " + "bye"
	}
	require.Equal(b, "hello bye", out)
}

func BenchmarkSprintf(b *testing.B) {
	var out string
	for i := 0; i <= b.N; i++ {
		out = fmt.Sprintf("%s %s", "hello", "bye")
	}
	require.Equal(b, "hello bye", out)
}

func BenchmarkBuilder(b *testing.B) {
	var out string
	for n := 0; n < b.N; n++ {
		var strBuilder strings.Builder
		strBuilder.WriteString("hello ")
		strBuilder.WriteString("bye")
		out = strBuilder.String()
	}
	require.Equal(b, "hello bye", out)
}

func BenchmarkBuilderGrow(b *testing.B) {
	var out string
	for n := 0; n < b.N; n++ {
		var strBuilder strings.Builder
		strBuilder.Grow(9)
		strBuilder.WriteString("hello ")
		strBuilder.WriteString("bye")
		out = strBuilder.String()
	}
	require.Equal(b, "hello bye", out)
}

func BenchmarkBuilderReset(b *testing.B) {
	var out string
	var strBuilder strings.Builder
	for n := 0; n < b.N; n++ {
		strBuilder.Grow(9)
		strBuilder.WriteString("hello ")
		strBuilder.WriteString("bye")
		out = strBuilder.String()
		strBuilder.Reset()
	}
	require.Equal(b, "hello bye", out)
}

func BenchmarkDynamicConcat(b *testing.B) {
	var out string
	for i := 0; i <= b.N; i++ {
		out = "hello " + strconv.Itoa(i)
	}
	require.NotZero(b, out)
}

func BenchmarkDynamicSprintf(b *testing.B) {
	var out string
	for i := 0; i <= b.N; i++ {
		out = fmt.Sprintf("%s %s", "hello", strconv.Itoa(i))
	}
	require.NotZero(b, out)
}

func BenchmarkDynamicBuilder(b *testing.B) {
	var out string
	for n := 0; n < b.N; n++ {
		var strBuilder strings.Builder
		strBuilder.WriteString("hello ")
		strBuilder.WriteString(strconv.Itoa(n))
		out = strBuilder.String()
	}
	require.NotZero(b, out)
}

func BenchmarkDynamicBuilderGrow(b *testing.B) {
	var out string
	for n := 0; n < b.N; n++ {
		var strBuilder strings.Builder
		strBuilder.Grow(9)
		strBuilder.WriteString("hello ")
		strBuilder.WriteString(strconv.Itoa(n))
		out = strBuilder.String()
	}
	require.NotZero(b, out)
}

func BenchmarkDynamicBuilderReset(b *testing.B) {
	var out string
	var strBuilder strings.Builder
	for n := 0; n < b.N; n++ {
		strBuilder.Grow(9)
		strBuilder.WriteString("hello ")
		strBuilder.WriteString(strconv.Itoa(n))
		out = strBuilder.String()
		strBuilder.Reset()
	}
	require.NotZero(b, out)
}

func BenchmarkRollingConcat(b *testing.B) {
	a := []string{"a", "b", "c", "d"}
	var out string
	for i := 0; i <= b.N; i++ {
		for _, a := range a {
			out = "hello " + a
		}
	}
	require.Equal(b, "hello d", out)
}

func BenchmarkRollingSprintf(b *testing.B) {
	var out string
	a := []string{"a", "b", "c", "d"}
	for i := 0; i <= b.N; i++ {
		for _, a := range a {
			out = fmt.Sprintf("%s %s", "hello", a)
		}
	}
	require.Equal(b, "hello d", out)
}
