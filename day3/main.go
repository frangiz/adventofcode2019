package main

import (
	"github.com/frangiz/adventofcode2019/utils"
	"strconv"
	"strings"
)

type wireStep struct {
	wireID, steps int
}

var dirs = map[string]aoc.Point{
	"R": aoc.Point{X: 1, Y: 0},
	"L": aoc.Point{X: -1, Y: 0},
	"U": aoc.Point{X: 0, Y: -1},
	"D": aoc.Point{X: 0, Y: 1},
}

type frontPanel struct {
	visited map[aoc.Point][]wireStep
}

func makeFrontPanel() frontPanel {
	return frontPanel{make(map[aoc.Point][]wireStep)}
}

func (fp frontPanel) outlineWire(wireID int, wire string) {
	pos := aoc.Point{X: 0, Y: 0}
	fp.visited[pos] = append(fp.visited[pos], wireStep{wireID, 0})
	totalSteps := 0
	for _, path := range strings.Split(wire, ",") {
		dir := dirs[path[:1]]
		steps, _ := strconv.Atoi(path[1:])
		for i := 1; i <= steps; i++ {
			totalSteps++
			pos = aoc.Point{X: pos.X + dir.X, Y: pos.Y + dir.Y}
			fp.visited[pos] = append(fp.visited[pos], wireStep{wireID, totalSteps})
		}
	}
}

func (fp frontPanel) smallestDistance() int {
	shortestDist := 9999999
	for k, v := range fp.visited {
		if (len(v) > 1 && k != aoc.Point{X: 0, Y: 0}) {
			wireIds := make(map[int]bool)
			for _, entry := range v {
				wireIds[entry.wireID] = true
			}
			if len(wireIds) > 1 {
				shortestDist = aoc.Min(shortestDist, aoc.ManhattanDistSingle(k))
			}
		}
	}
	return shortestDist
}

func (fp frontPanel) fewestCombinedSteps() int {
	combinedSteps := 9999999
	for k, v := range fp.visited {
		if (len(v) > 1 && k != aoc.Point{X: 0, Y: 0}) {
			wireIds := make(map[int]int)
			for _, entry := range v {
				if _, ok := wireIds[entry.wireID]; !ok {
					wireIds[entry.wireID] = entry.steps
				}
			}
			if len(wireIds) > 1 {
				steps := 0
				for _, v2 := range wireIds {
					steps += v2
				}
				combinedSteps = aoc.Min(combinedSteps, steps)
			}
		}
	}
	return combinedSteps
}

func partA() int {
	input := aoc.ReadInputAsStr("input.txt")
	fp := makeFrontPanel()
	for id, wire := range input {
		fp.outlineWire(id, wire)
	}
	return fp.smallestDistance()
}

func partB() int {
	input := aoc.ReadInputAsStr("input.txt")
	fp := makeFrontPanel()
	for id, wire := range input {
		fp.outlineWire(id, wire)
	}
	return fp.fewestCombinedSteps()
}

func main() {
}
