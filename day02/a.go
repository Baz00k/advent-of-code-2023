package day02

import (
	"strconv"
	"strings"
)

// A map of the color and the maximum amount of cubes that can be used
var maxCubes = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
}

func SolvePartA(input []string) int {
	sum := 0

	for index, line := range input {
		// Line is in the format of "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
		input := strings.Split(line, ": ")[1]
		subsets := strings.Split(input, "; ")

		isGameValid := true

		// Subset is in the format of "3 blue, 4 red"
		for _, subset := range subsets {
			cubes := strings.Split(subset, ", ")

			// Cube is in the format of "3 blue"
			for _, cube := range cubes {
				cubeParts := strings.Split(cube, " ")

				amount, err := strconv.Atoi(cubeParts[0])
				color := cubeParts[1]

				if err != nil {
					panic(err)
				}

				if amount > maxCubes[color] {
					isGameValid = false
					break
				}
			}
		}

		if isGameValid {
			sum += index + 1
		}
	}
	
	return sum
}
