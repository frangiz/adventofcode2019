package main

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

var exampleATests = []struct {
	input    []int
	expected []int
}{
	{[]int{1, 0, 0, 0, 99}, []int{2, 0, 0, 0, 99}},
	{[]int{2, 3, 0, 3, 99}, []int{2, 3, 0, 6, 99}},
	{[]int{2, 4, 4, 5, 99, 0}, []int{2, 4, 4, 5, 99, 9801}},
	{[]int{1, 1, 1, 4, 99, 5, 6, 0, 99}, []int{30, 1, 1, 4, 2, 5, 6, 0, 99}},
}

func TestExampleA(t *testing.T) {
	for _, dt := range exampleATests {
		actual := runIntCode(dt.input)
		if !reflect.DeepEqual(actual, dt.expected) {
			t.Errorf("runIntCode(%d): expected %d, actual %d", dt.input, dt.expected, actual)
		}
		assert.Equal(t, true, reflect.DeepEqual(actual, dt.expected))
	}
}

func TestAnswerA(t *testing.T) {
	assert.Equal(t, 12490719, partA())
}

func TestAnswerB(t *testing.T) {
	assert.Equal(t, 2003, partB())
}
