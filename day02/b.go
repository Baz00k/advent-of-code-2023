package day02

import (
	"strconv"
	"strings"
)

func SolvePartB(input []string) int {
	sum := 0

	for _, line := range input {
		// Line is in the format of "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
		input := strings.Split(line, ": ")[1]
		subsets := strings.Split(input, "; ")

		maxCubesPerColor := map[string]int{
			"red": 0,
			"green": 0,
			"blue": 0,
		}

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

				if amount > maxCubesPerColor[color] {
					maxCubesPerColor[color] = amount
				}
			}
		}

		power := 1
		for _, value := range maxCubesPerColor {
			power *= value
		}

		sum += power
	}

	return sum
}