package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"gopkg.in/yaml.v2"
	// "github.com/goccy/go-yaml"
)

type Large struct {
	Hellos []string
}

type Small struct {
	Name string
}

var test = `
hellos:
- hi
- hello
`

func BenchmarkBase(b *testing.B) {
	var bb Large
	var err error
	for i := 0; i <= b.N; i++ {
		err = yaml.Unmarshal([]byte(test), &bb)
	}
	require.NoError(b, err)
	assert.Len(b, bb.Hellos, 12)
}

type Large2 struct {
	A      string
	Hellos []string `yaml:"hellos"`
}

func (l *Large2) UnmarshalYAML(unmarshal func(interface{}) error) error {
	// l = &Large2{}
	target := make(map[string]string, 12)
	err := unmarshal(&target)
	if err != nil {
		fmt.Println(err)
	}
	// l.Hellos = target
	return nil
}

func BenchmarkCustom(b *testing.B) {
	var bb Large2
	var err error
	// bb.Hellos = make([]string, 4)
	for i := 0; i <= b.N; i++ {
		err = yaml.Unmarshal([]byte(test), &bb)
	}
	require.NoError(b, err)
	assert.Len(b, bb.Hellos, 12)
}
