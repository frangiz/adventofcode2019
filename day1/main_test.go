package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleATests = []struct {
	mass int
	fuel int
}{
	{12, 2},
	{14, 2},
	{1969, 654},
	{100756, 33583},
}

func TestExampleA(t *testing.T) {
	for _, dt := range exampleATests {
		actual := calcFuel(dt.mass)
		if actual != dt.fuel {
			t.Errorf("calcFuel(%d): expected %d, actual %d", dt.mass, dt.fuel, actual)
		}
		assert.Equal(t, dt.fuel, actual)
	}
}

func TestAnswerA(t *testing.T) {
	assert.Equal(t, 3399394, partA())
}

var exampleBTests = []struct {
	mass int
	fuel int
}{
	{14, 2},
	{1969, 966},
	{100756, 50346},
}

func TestExampleB(t *testing.T) {
	for _, dt := range exampleBTests {
		actual := calcFuelPartB(dt.mass)
		if actual != dt.fuel {
			t.Errorf("calcFuel(%d): expected %d, actual %d", dt.mass, dt.fuel, actual)
		}
		assert.Equal(t, dt.fuel, actual)
	}
}

func TestAnswerB(t *testing.T) {
	assert.Equal(t, 5096223, partB())
}
