package test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/multierr"
)

const _alphanums = "bcdfghjklmnpqrstvwxz2456789"

const _idLength = 5
const l = len(_alphanums)

const (
	_placeValue0 = _placeValue1 * len(_alphanums)
	_placeValue1 = _placeValue2 * len(_alphanums)
	_placeValue2 = _placeValue3 * len(_alphanums)
	_placeValue3 = len(_alphanums)
)

func func1(id string) (int, error) {
	index := 0

	rIndex := strings.IndexByte(_alphanums, id[0])
	if rIndex == -1 {
		return -1, fmt.Errorf("%q contains invalid character %q", id, id[0])
	}
	index += (_placeValue0 * rIndex)

	rIndex = strings.IndexByte(_alphanums, id[1])
	if rIndex == -1 {
		return -1, fmt.Errorf("%q contains invalid character %q", id, id[1])
	}
	index += (_placeValue1 * rIndex)

	rIndex = strings.IndexByte(_alphanums, id[2])
	if rIndex == -1 {
		return -1, fmt.Errorf("%q contains invalid character %q", id, id[2])
	}
	index += (_placeValue2 * rIndex)

	rIndex = strings.IndexByte(_alphanums, id[3])
	if rIndex == -1 {
		return -1, fmt.Errorf("%q contains invalid character %q", id, id[3])
	}
	index += (_placeValue3 * rIndex)

	rIndex = strings.IndexByte(_alphanums, id[4])
	if rIndex == -1 {
		return -1, fmt.Errorf("%q contains invalid character %q", id, id[4])
	}
	index += rIndex

	return index, nil
}

func func2(id string) (int, error) {
	index := 0
	var err error

	add := func(b byte, step int) int {
		rIndex := strings.IndexByte(_alphanums, b)
		if rIndex == -1 {
			err = multierr.Append(err, fmt.Errorf("%q contains invalid character %q", id, b))
		}
		return (step * rIndex)
	}

	index += add(id[0], _placeValue0)
	index += add(id[1], _placeValue1)
	index += add(id[2], _placeValue2)
	index += add(id[3], _placeValue3)
	index += add(id[4], 1)

	return index, nil
}

func add2(b byte, step int, err *[5]error) int {
	rIndex := strings.IndexByte(_alphanums, b)
	if rIndex == -1 {
		err[0] = fmt.Errorf("contains invalid character %q", b)
	}
	return (step * rIndex)
}

func func3(id string) (int, error) {
	index := 0
	var err [5]error

	index += add2(id[0], _placeValue0, &err)
	index += add2(id[1], _placeValue1, &err)
	index += add2(id[2], _placeValue2, &err)
	index += add2(id[3], _placeValue3, &err)
	index += add2(id[4], 1, &err)

	return index, err[0]
}

func BenchmarkWrapper(b *testing.B) {
	var r int
	var err error
	for i := 0; i < b.N; i++ {
		r, err = func3("99999")
	}
	require.NoError(b, err)
	assert.Equal(b, 14348906, r)
}

func BenchmarkUnrolled(b *testing.B) {
	var r int
	var err error
	for i := 0; i < b.N; i++ {
		r, err = func1("99999")
	}
	require.NoError(b, err)
	assert.Equal(b, 14348906, r)
}

func BenchmarkClosure(b *testing.B) {
	var r int
	var err error
	for i := 0; i < b.N; i++ {
		r, err = func2("99999")
	}
	require.NoError(b, err)
	assert.Equal(b, 14348906, r)
}
