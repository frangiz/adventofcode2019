package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// ReadInputAsStr stuff
func ReadInputAsStr(path string) []string {
	bytes, _ := ioutil.ReadFile(path)
	return strings.Split(string(bytes), "\r\n")
}

// ReadInputAsInt stuff
func ReadInputAsInt(path string) []int {
	ints := []int{}
	for _, line := range ReadInputAsStr(path) {
		invVal, _ := strconv.Atoi(line)
		ints = append(ints, invVal)
	}
	return ints
}

// ReadInputLineSplitOn stuff
func ReadInputLineSplitOn(path, seperator string) []string {
	bytes, _ := ioutil.ReadFile(path)
	return strings.Split(string(bytes), seperator)
}

// ReadInputIntsLineSplitOn stuff
func ReadInputIntsLineSplitOn(path, seperator string) []int {
	ints := []int{}
	for _, line := range ReadInputLineSplitOn(path, seperator) {
		invVal, _ := strconv.Atoi(line)
		ints = append(ints, invVal)
	}
	return ints
}

// Min stuff
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// Abs stuff
func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
