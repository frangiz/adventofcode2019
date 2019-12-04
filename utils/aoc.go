package aoc

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

// Point stuff
type Point struct {
	X, Y int
}

// ManhattanDist stuff
func ManhattanDist(p1, p2 Point) int {
	return Abs(p1.X-p2.X) + Abs(p2.Y-p2.Y)
}

// ManhattanDistSingle stuff
func ManhattanDistSingle(p Point) int {
	return Abs(p.X) + Abs(p.Y)
}
