package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var exampleATests = []struct {
	input    string
	expected bool
}{
	{"111111", true},
	{"223450", false},
	{"123789", false},
}

func TestExampleA(t *testing.T) {
	for _, dt := range exampleATests {
		assert.Equal(t, dt.expected, checkCriteria1(dt.input))
	}
}

func TestAnswerA(t *testing.T) {
	assert.Equal(t, 1246, partA())
}

var exampleBTests = []struct {
	input    string
	expected bool
}{
	{"112233", true},
	{"123444", false},
	{"111122", true},
}

func TestExampleB(t *testing.T) {
	for _, dt := range exampleBTests {
		assert.Equal(t, dt.expected, checkCriteria2(dt.input))
	}
}

func TestAnswerB(t *testing.T) {
	assert.Equal(t, 814, partB())
}
