package test

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func read() {
	_, err := ioutil.ReadFile("tmp")
	if err != nil {
		fmt.Println(err)
	}
}

func load() map[string]string {
	r := make(map[string]string, 0)
	file, _ := os.Open("tmp")
	scanner := bufio.NewScanner(file)
	i := 0

	for scanner.Scan() {
		t := strings.Split(scanner.Text(), ",")
		i++
		r[t[0]] = t[1]
	}
	return r
}

func BenchmarkRead(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		read()
	}
}

func BenchmarkLoad(b *testing.B) {
	var r map[string]string
	for i := 0; i <= b.N; i++ {
		r = load()
	}
	assert.NotNil(b, r)
	assert.Greater(b, len(r), 100000)
}

func BenchmarkGet(b *testing.B) {
	r := map[string]string{"10.13.0.17": "5A"}
	var s string
	for i := 0; i <= b.N; i++ {
		s = r["10.13.0.17"]
	}
	assert.NotNil(b, s)
}

func BenchmarkGet2(b *testing.B) {
	r := load()
	var s string
	for i := 0; i <= b.N; i++ {
		s = r["10.13.0.17"]
	}
	assert.NotNil(b, s)
}
