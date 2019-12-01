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
