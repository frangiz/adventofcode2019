package aoc

import (
	"strconv"
)

// Atoi parses a string to int
func Atoi(s string) int {
	i, err := strconv.Atoi(s)
	Check(err)
	return i
}

// AbsInt returns the absolute value of x
func AbsInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// MinInt returns the smallest value of a, b
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns the largest value of a, b
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
