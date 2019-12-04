package main

import (
	"github.com/frangiz/adventofcode2019/utils"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

type wireStep struct {
	wireID, steps int
}

type frontPanel struct {
	visited map[point][]wireStep
}

func makeFrontPanel() frontPanel {
	return frontPanel{make(map[point][]wireStep)}
}

func (fp frontPanel) outlineWire(wireID int, wire string) {
	pos := point{0, 0}
	fp.visited[pos] = append(fp.visited[pos], wireStep{wireID, 0})
	totalSteps := 0
	for _, dir := range strings.Split(wire, ",") {
		if strings.HasPrefix(dir, "R") {
			steps, _ := strconv.Atoi(dir[1:])
			for i := 1; i <= steps; i++ {
				totalSteps++
				pos = point{pos.x + 1, pos.y}
				fp.visited[pos] = append(fp.visited[pos], wireStep{wireID, totalSteps})
			}
		}
		if strings.HasPrefix(dir, "U") {
			steps, _ := strconv.Atoi(dir[1:])
			for i := 1; i <= steps; i++ {
				totalSteps++
				pos = point{pos.x, pos.y - 1}
				fp.visited[pos] = append(fp.visited[pos], wireStep{wireID, totalSteps})
			}
		}
		if strings.HasPrefix(dir, "L") {
			steps, _ := strconv.Atoi(dir[1:])
			for i := 1; i <= steps; i++ {
				totalSteps++
				pos = point{pos.x - 1, pos.y}
				fp.visited[pos] = append(fp.visited[pos], wireStep{wireID, totalSteps})
			}
		}
		if strings.HasPrefix(dir, "D") {
			steps, _ := strconv.Atoi(dir[1:])
			for i := 1; i <= steps; i++ {
				totalSteps++
				pos = point{pos.x, pos.y + 1}
				fp.visited[pos] = append(fp.visited[pos], wireStep{wireID, totalSteps})
			}
		}
	}
}

func (fp frontPanel) smallestDistance() int {
	shortestDist := 9999999
	for k, v := range fp.visited {
		if (len(v) > 1 && k != point{0, 0}) {
			wireIds := make(map[int]bool)
			for _, entry := range v {
				wireIds[entry.wireID] = true
			}
			if len(wireIds) > 1 {
				shortestDist = utils.Min(shortestDist, utils.Abs(k.x)+utils.Abs(k.y))
			}
		}
	}
	return shortestDist
}

func (fp frontPanel) fewestCombinedSteps() int {
	combinedSteps := 9999999
	for k, v := range fp.visited {
		if (len(v) > 1 && k != point{0, 0}) {
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
				combinedSteps = utils.Min(combinedSteps, steps)
			}
		}
	}
	return combinedSteps
}

func partA() int {
	input := utils.ReadInputAsStr("input.txt")
	fp := makeFrontPanel()
	for id, wire := range input {
		fp.outlineWire(id, wire)
	}
	return fp.smallestDistance()
}

func partB() int {
	input := utils.ReadInputAsStr("input.txt")
	fp := makeFrontPanel()
	for id, wire := range input {
		fp.outlineWire(id, wire)
	}
	return fp.fewestCombinedSteps()
}

func main() {
}
