package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleATests = []struct {
	input    []string
	expected int
}{
	{[]string{"R8,U5,L5,D3", "U7,R6,D4,L4"}, 6},
	{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 159},
	{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 135},
}

func TestExampleA(t *testing.T) {
	for _, dt := range exampleATests {
		fp := makeFrontPanel()
		for id, wire := range dt.input {
			fp.outlineWire(id, wire)
		}
		assert.Equal(t, dt.expected, fp.smallestDistance())
	}
}

func TestAnswerA(t *testing.T) {
	assert.Equal(t, 1195, partA())
}

var exampleBTests = []struct {
	input    []string
	expected int
}{
	{[]string{"R8,U5,L5,D3", "U7,R6,D4,L4"}, 30},
	{[]string{"R75,D30,R83,U83,L12,D49,R71,U7,L72", "U62,R66,U55,R34,D71,R55,D58,R83"}, 610},
	{[]string{"R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51", "U98,R91,D20,R16,D67,R40,U7,R15,U6,R7"}, 410},
}

func TestExampleB(t *testing.T) {
	for _, dt := range exampleBTests {
		fp := makeFrontPanel()
		for id, wire := range dt.input {
			fp.outlineWire(id, wire)
		}
		assert.Equal(t, dt.expected, fp.fewestCombinedSteps())
	}
}

func TestAnswerB(t *testing.T) {
	assert.Equal(t, 91518, partB())
}
