package aoc

import (
	"strings"
)

// IntArrayEquals stuff
func IntArrayEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// StrToIntArray stuff
func StrToIntArray(str string) []int {
	res := make([]int, len(str))
	for i, n := range strings.Split(str, "") {
		res[i] = Atoi(n)
	}
	return res
}
