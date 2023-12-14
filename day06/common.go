package day06

import (
	"strconv"
	"strings"
)

func parseInput(input []string, ignoreSpaces bool) (times []int, recordDistances []int) {
	var timesString []string
	if ignoreSpaces {
		timesString = []string{strings.ReplaceAll(strings.Split(input[0], ": ")[1], " ", "")}
	} else {
		timesString = strings.Split(strings.Split(input[0], ": ")[1], " ")
	}

	for _, v := range timesString {
		if v == "" {
			continue
		}
		time, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		times = append(times, time)
	}

	var distancesString []string
	if ignoreSpaces {
		distancesString = []string{strings.ReplaceAll(strings.Split(input[1], ": ")[1], " ", "")}
	} else {
		distancesString = strings.Split(strings.Split(input[1], ": ")[1], " ")
	}

	for _, v := range distancesString {
		if v == "" {
			continue
		}

		distance, err := strconv.Atoi(v)

		if err != nil {
			panic(err)
		}

		recordDistances = append(recordDistances, distance)
	}

	return
}

func findLowestSpeedToBeatRecord(time int, recordDistance int) int {
	left := 1
	right := time

	for left < right {
		mid := left + (right-left)/2
		if calculateDistanceTraveled(time, mid) < recordDistance {
			left = mid + 1
		} else {
			right = mid
		}
	}

	// Check if the lowest speed found still beats the record
	if calculateDistanceTraveled(time, left) <= recordDistance {
		return left + 1
	}

	return left
}

func findHighestSpeedToBeatRecord(time int, recordDistance int) int {
	left := 1
	right := time

	for left < right {
		mid := left + (right-left+1)/2
		if calculateDistanceTraveled(time, mid) > recordDistance {
			left = mid
		} else {
			right = mid - 1
		}
	}

	// Check if the highest speed found still beats the record
	if calculateDistanceTraveled(time, left) <= recordDistance {
		return left - 1
	}

	return left
}

func calculateDistanceTraveled(time int, speed int) int {
	return speed * (time - speed)
}
