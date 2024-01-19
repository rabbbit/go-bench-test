package test

import (
	"strconv"
	"strings"
)

func func1(i int) string {
	return "q=" + strconv.Itoa(i)
}

func func2(s string) string {
	return strings.ToLower(s)
}
