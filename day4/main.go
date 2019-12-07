package main

import (
	"github.com/frangiz/adventofcode2019/pkg/aoc"
	"strconv"
	"strings"
)

func checkCriteria1(str string) bool {
	hasDoubleVal := false
	for i := range str {
		if i == 0 {
			continue
		}
		if str[i] == str[i-1] {
			hasDoubleVal = true
		}
		if str[i-1] > str[i] {
			return false
		}
	}
	return hasDoubleVal
}

func checkCriteria2(str string) bool {
	matching := false
	gotDouble := false
	for i := 0; i < len(str)-1; i++ {
		if str[i] > str[i+1] {
			return false
		}
		if str[i] == str[i+1] {
			matching = true
			if (i > 0 && str[i-1] == str[i]) || (i < len(str)-2 && str[i] == str[i+2]) {
				matching = false
			}
			if matching {
				gotDouble = true
			}
		}
	}
	return gotDouble
}

func partA() int {
	input := aoc.ReadInputAsStr("input.txt")
	ranges := strings.Split(input, "-")
	min := aoc.Atoi(ranges[0])
	max := aoc.Atoi(ranges[1])
	count := 0
	for n := min; n < max; n++ {
		if checkCriteria1(strconv.Itoa(n)) {
			count++
		}
	}

	return count
}

func partB() int {
	input := aoc.ReadInputAsStr("input.txt")
	ranges := strings.Split(input, "-")
	min := aoc.Atoi(ranges[0])
	max := aoc.Atoi(ranges[1])
	count := 0
	for n := min; n < max; n++ {
		if checkCriteria2(strconv.Itoa(n)) {
			count++
		}
	}

	return count
}

func main() {
}
