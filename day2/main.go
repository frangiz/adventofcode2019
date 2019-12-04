package main

import (
	"github.com/frangiz/adventofcode2019/pkg/aoc"
)

func runIntCode(input []int) []int {
	ip := 0
	opcode := input[ip]
	for opcode != 99 {
		a := input[ip+1]
		b := input[ip+2]
		res := input[ip+3]
		if opcode == 1 {
			input[res] = input[a] + input[b]
		} else if opcode == 2 {
			input[res] = input[a] * input[b]
		}
		ip = (ip + 4) % len(input)
		opcode = input[ip]
	}
	return input
}

func partA() int {
	input := aoc.ReadInputIntsLineSplitOn("input.txt", ",")
	input[1] = 12
	input[2] = 2
	return runIntCode(input)[0]
}

func partB() int {
	input := aoc.ReadInputIntsLineSplitOn("input.txt", ",")
	for noun := 0; noun <= 99; noun++ {
		for verb := 0; verb <= 99; verb++ {
			tmp := make([]int, len(input))
			copy(tmp, input)
			tmp[1] = noun
			tmp[2] = verb
			if runIntCode(tmp)[0] == 19690720 {
				return 100*noun + verb
			}
		}
	}
	return 0
}

func main() {
}
