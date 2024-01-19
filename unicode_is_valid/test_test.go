package test

import (
	"testing"
	"unicode"
)

func Valid(x []byte) bool {
	state := uint64(1 * 6)
	for _, b := range x {
		state = dfa[b] >> (state & 63)
	}
	return (state & 63) == 1*6
}

const (
	s00 = 0 | (1*6)<<(1*6)
	sC0 = 0
	sC2 = 0 | (2*6)<<(1*6)
	sE0 = 0 | (3*6)<<(1*6)
	sE1 = 0 | (4*6)<<(1*6)
	sED = 0 | (5*6)<<(1*6)
	sEE = sE1
	sF0 = 0 | (6*6)<<(1*6)
	sF1 = 0 | (7*6)<<(1*6)
	sF4 = 0 | (8*6)<<(1*6)
	sF5 = 0

	s80 = 0 | (1*6)<<(2*6) | (2*6)<<(4*6) | (4*6)<<(5*6) | (4*6)<<(7*6) | (4*6)<<(8*6)
	s90 = 0 | (1*6)<<(2*6) | (2*6)<<(4*6) | (4*6)<<(5*6) | (4*6)<<(6*6) | (4*6)<<(7*6)
	sA0 = 0 | (1*6)<<(2*6) | (2*6)<<(3*6) | (2*6)<<(4*6) | (4*6)<<(6*6) | (4*6)<<(7*6)
)

var dfa = [256]uint64{
	s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00,
	s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00,
	s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00,
	s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00,
	s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00,
	s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00,
	s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00,
	s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00, s00,
	s80, s80, s80, s80, s80, s80, s80, s80, s80, s80, s80, s80, s80, s80, s80, s80,
	s90, s90, s90, s90, s90, s90, s90, s90, s90, s90, s90, s90, s90, s90, s90, s90,
	sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0,
	sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0, sA0,
	sC0, sC0, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2,
	sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2, sC2,
	sE0, sE1, sE1, sE1, sE1, sE1, sE1, sE1, sE1, sE1, sE1, sE1, sE1, sED, sEE, sEE,
	sF0, sF1, sF1, sF1, sF4, sF5, sF5, sF5, sF5, sF5, sF5, sF5, sF5, sF5, sF5, sF5,
}

func goodName2(name string) bool {
	if len(name) == 0 {
		return false
	}
	if name[0] >= '0' && name[0] <= '9' {
		return goodName(name)
	}
	for i := 0; i < len(name); i++ {
		if !((name[i] >= 'a' && name[i] <= 'z') || (name[i] >= 'A' && name[i] <= 'Z') || name[i] == '_' || (name[i] >= '0' && name[i] <= '9')) {
			return goodName(name)
		}
	}
	return true
}

func goodName(name string) bool {
	if name == "" {
		return false
	}
	for i, r := range name {
		switch {
		case r == '_':
		case i == 0 && !unicode.IsLetter(r):
			return false
		case !unicode.IsLetter(r) && !unicode.IsDigit(r):
			return false
		}
	}
	return true
}

func Benchmark_goodName_AsciiOnly(b *testing.B) {
	b.ReportAllocs()
	var x bool
	for i := 0; i < b.N; i++ {
		x = goodName("User_NameFunc1")
	}
	if !x {
		b.Fail()
	}
}
func Benchmark_goodName_NonAscii(b *testing.B) {
	b.ReportAllocs()
	var x bool
	for i := 0; i < b.N; i++ {
		x = goodName("कखग_१")
	}
	if !x {
		b.Fail()
	}
}

func Benchmark_goodName_AsciiOnly2(b *testing.B) {
	b.ReportAllocs()
	var x bool
	for i := 0; i < b.N; i++ {
		x = goodName2("User_NameFunc1")
	}
	if !x {
		b.Fail()
	}
}
func Benchmark_goodName_NonAscii2(b *testing.B) {
	b.ReportAllocs()
	var x bool
	for i := 0; i < b.N; i++ {
		x = goodName2("कखग_१")
	}
	if !x {
		b.Fail()
	}
}

func Benchmark_valid_AsciiOnly(b *testing.B) {
	b.ReportAllocs()
	var x bool
	for i := 0; i < b.N; i++ {
		x = Valid([]byte("User_NameFunc1"))
	}
	if !x {
		b.Fail()
	}
}
func Benchmark_valid_NonAscii(b *testing.B) {
	b.ReportAllocs()
	var x bool
	for i := 0; i < b.N; i++ {
		x = Valid([]byte("कखग_१"))
	}
	if !x {
		b.Fail()
	}
}

func Test_goodName(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{
			"valid_name",
			"validname12",
			true,
		},
		{
			"empty_string",
			"",
			false,
		},
		{
			"just_whitespace",
			" ",
			false,
		},
		{
			"starts_with_number",
			"1invalid",
			false,
		},
		{
			"has_underscore",
			"still_valid_as_underscores_are_fine",
			true,
		},
		{
			"multiple_non_consecutive_digits",
			"v1a2l3id",
			true,
		},
		{
			"embedded whitespace",
			"v alid",
			false,
		},
		{
			"starts_with_underscore",
			"_123",
			true,
		},
		{
			"non_ascii_valid",
			"कखग_१", // "कखग" are all Category L (letters) and "१" is Category Nd (digit) in unicode spec
			true,
		},
		{
			"non_ascii_invalid",
			"emoji_is_not_category_l_for_unicode_☹", // ☹ is part of Category_So in unicode spec, so should fail
			false,
		},
		{
			"valid_unicode_number_but_not_digit",
			// Note
			// that
			// unicode.isNumber("½")
			// is
			// true,
			// but
			// unicode.isDigit("½")
			// is
			// false
			"invalid_½", // "½" is Category_No (which is different from a digit - Category Nd)
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := goodName(tt.args); got != tt.want {
				t.Errorf("goodName() = %v, want %v", got, tt.want)
			}
			if got := Valid([]byte(tt.args)); got != tt.want {
				t.Errorf("validName() = %v, want %v", got, tt.want)
			}
		})
	}
}
