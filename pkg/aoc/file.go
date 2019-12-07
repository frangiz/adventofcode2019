package aoc

import (
	"io/ioutil"
	"strings"
)

// ReadInputAsStr stuff
func ReadInputAsStr(path string) string {
	bytes, _ := ioutil.ReadFile(path)
	return string(bytes)
}

// ReadInputAsStrArray stuff
func ReadInputAsStrArray(path string) []string {
	bytes, _ := ioutil.ReadFile(path)
	return strings.Split(string(bytes), "\r\n")
}

// ReadInputAsIntArray stuff
func ReadInputAsIntArray(path string) []int {
	ints := []int{}
	for _, line := range ReadInputAsStrArray(path) {
		ints = append(ints, Atoi(line))
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
		ints = append(ints, Atoi(line))
	}
	return ints
}
