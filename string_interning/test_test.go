package test

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func BenchmarkNaive(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		out := make(map[string]string, 100_000)
		f, err := os.Open("server_types")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			s := strings.Split(scanner.Text(), ",")
			out[s[0]] = s[1]
		}
		assert.NotEqual(b, 0, len(out))
		assert.Equal(b, "5G", out["10.13.0.130"])
	}
	s := time.Now()
	runtime.GC()
	fmt.Println("GC took", time.Now().Sub(s))
}

func BenchmarkBytes(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		out := make(map[string]string, 100_000)
		f, err := os.Open("server_types")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			b := scanner.Bytes()
			i := bytes.IndexByte(b, ',')
			out[string(b[:i])] = string(b[i+1:])
		}
		assert.NotEqual(b, 0, len(out))
		assert.Equal(b, "5G", out["10.13.0.130"])
	}
	s := time.Now()
	runtime.GC()
	fmt.Println("GC took", time.Now().Sub(s))
}

type interner map[string]string

func (si interner) intern(b []byte) string {
	if interned, ok := si[string(b)]; ok {
		return interned
	}
	s := string(b)
	si[s] = s
	return s
}

func newInterner() interner {
	i := make(map[string]string, 100)
	return interner(i)
}

func BenchmarkInterned(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		in := newInterner()
		out := make(map[string]string, 100_000)
		f, err := os.Open("server_types")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			b := scanner.Bytes()
			i := bytes.IndexByte(b, ',')
			out[string(b[:i])] = in.intern(b[i+1:])
		}
		assert.NotEqual(b, 0, len(out))
		assert.Equal(b, "5G", out["10.13.0.130"])
	}
	s := time.Now()
	runtime.GC()
	fmt.Println("GC took", time.Now().Sub(s))
}

func Benchmark15BytesInterned(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		in := newInterner()
		out := make(map[[15]byte]string, 100_000)
		f, err := os.Open("server_types")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			b := scanner.Bytes()
			i := bytes.IndexByte(b, ',')
			var v [15]byte
			copy(v[:], b[:i])
			out[v] = in.intern(b[i+1:])
		}
		assert.NotEqual(b, 0, len(out))
	}
	s := time.Now()
	runtime.GC()
	fmt.Println("GC took", time.Now().Sub(s))
}

func Benchmark4BytesInterned(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		in := newInterner()
		out := make(map[[4]byte]string, 100_000)
		f, err := os.Open("server_types")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			b := scanner.Bytes()
			var v [4]byte
			i := 0
			j := bytes.IndexByte(b[i:], '.')
			vv, _ := atoi(b[i:j])
			v[0] = vv
			i += j + 1
			j = bytes.IndexByte(b[i:], '.')
			vv, _ = atoi(b[i : i+j])
			v[1] = vv
			i += j + 1
			j = bytes.IndexByte(b[i:], '.')
			vv, _ = atoi(b[i : i+j])
			v[2] = vv
			i += j + 1
			j = bytes.IndexByte(b[i:], ',')
			vv, _ = atoi(b[i : i+j])
			v[3] = vv

			out[v] = in.intern(b[i+j+1:])
		}
		assert.NotEqual(b, 0, len(out))
		assert.Equal(b, "5G", out[[4]byte{10, 13, 0, 130}])
	}
	s := time.Now()
	runtime.GC()
	fmt.Println("GC took", time.Now().Sub(s))
}

func Benchmark4BytesInternedAlt(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		in := newInterner()
		out := make(map[[4]byte]string, 100_000)
		f, err := os.Open("server_types")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()

		scanner := bufio.NewScanner(f)
		for scanner.Scan() {
			var v [4]byte
			b := scanner.Bytes()
			j := bytes.IndexByte(b, '.')
			vv, _ := atoi(b[:j])
			v[0] = vv
			b = b[j+1:]
			j = bytes.IndexByte(b, '.')
			vv, _ = atoi(b[:j])
			v[1] = vv
			b = b[j+1:]
			j = bytes.IndexByte(b, '.')
			vv, _ = atoi(b[:j])
			v[2] = vv
			b = b[j+1:]
			j = bytes.IndexByte(b, ',')
			vv, _ = atoi(b[:j])
			v[3] = vv

			out[v] = in.intern(b[j+1:])
		}
		assert.NotEqual(b, 0, len(out))
		assert.Equal(b, "5G", out[[4]byte{10, 13, 0, 130}])
	}
	s := time.Now()
	runtime.GC()
	fmt.Println("GC took", time.Now().Sub(s))
}

func atoi(s []byte) (byte, bool) {
	// We could return MaxInt, but returning error matches `strconv.atoi`.
	if len(s) > 10 {
		return byte(0), false
	}

	n := 0
	for _, v := range s {
		if v == ',' {
			return byte(n), true
		}
		v -= '0'
		if v > 9 {
			return 0, false
		}
		n = n*10 + int(v)
	}

	return byte(n), true
}
