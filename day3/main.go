package main

import (
	"github.com/frangiz/adventofcode2019/pkg/aoc"
	"strings"
)

type wireStep struct {
	wireID, steps int
}

var dirs = map[string]aoc.Vector2D{
	"R": aoc.MakeVector2D(1, 0),
	"L": aoc.MakeVector2D(-1, 0),
	"U": aoc.MakeVector2D(0, -1),
	"D": aoc.MakeVector2D(0, 1),
}

type frontPanel struct {
	visited map[aoc.Vector2D][]wireStep
}

func makeFrontPanel() frontPanel {
	return frontPanel{make(map[aoc.Vector2D][]wireStep)}
}

func (fp frontPanel) outlineWire(wireID int, wire string) {
	pos := aoc.MakeVector2D(0, 0)
	fp.visited[pos] = append(fp.visited[pos], wireStep{wireID, 0})
	totalSteps := 0
	for _, path := range strings.Split(wire, ",") {
		dir := dirs[path[:1]]
		for i := 1; i <= aoc.Atoi(path[1:]); i++ {
			totalSteps++
			pos = pos.Add(dir)
			fp.visited[pos] = append(fp.visited[pos], wireStep{wireID, totalSteps})
		}
	}
}

func (fp frontPanel) smallestDistance() int {
	shortestDist := 9999999
	for k, v := range fp.visited {
		if len(v) > 1 && k != aoc.MakeVector2D(0, 0) {
			wireIds := make(map[int]bool)
			for _, entry := range v {
				wireIds[entry.wireID] = true
			}
			if len(wireIds) > 1 {
				shortestDist = aoc.MinInt(shortestDist, k.ManhattanLength())
			}
		}
	}
	return shortestDist
}

func (fp frontPanel) fewestCombinedSteps() int {
	combinedSteps := 9999999
	for k, v := range fp.visited {
		if len(v) > 1 && k != aoc.MakeVector2D(0, 0) {
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
				combinedSteps = aoc.MinInt(combinedSteps, steps)
			}
		}
	}
	return combinedSteps
}

func partA() int {
	input := aoc.ReadInputAsStrArray("input.txt")
	fp := makeFrontPanel()
	for id, wire := range input {
		fp.outlineWire(id, wire)
	}
	return fp.smallestDistance()
}

func partB() int {
	input := aoc.ReadInputAsStrArray("input.txt")
	fp := makeFrontPanel()
	for id, wire := range input {
		fp.outlineWire(id, wire)
	}
	return fp.fewestCombinedSteps()
}

func main() {
}
