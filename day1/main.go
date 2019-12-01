package main

import (
	"github.com/frangiz/adventofcode2019/utils"
)

func calcFuel(mass int) int {
	return (mass / 3) - 2
}

func calcFuelPartB(mass int) int {
	fuel := calcFuel(mass)
	if fuel <= 0 {
		return 0
	}
	return fuel + calcFuelPartB(fuel)
}

func partA() int {
	input := utils.ReadInputAsInt("input.txt")
	sum := 0
	for _, mass := range input {
		sum += calcFuel(mass)
	}
	return sum
}

func partB() int {
	input := utils.ReadInputAsInt("input.txt")
	sum := 0
	for _, mass := range input {
		sum += calcFuelPartB(mass)
	}
	return sum
}

func main() {
}
