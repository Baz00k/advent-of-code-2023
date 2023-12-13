package day05

import (
	"strconv"
	"strings"
)

func SolvePartB(input []string) int {
	// First line is the list of seeds:
	// seeds: 79 14 55 ...
	seedsLine := strings.Split(input[0], " ")

	var rangeStart int
	var seedRanges []valuesRange

	for _, seed := range seedsLine[1:] {
		seedInt, err := strconv.Atoi(seed)

		if err != nil {
			panic(err)
		}

		if rangeStart == 0 {
			rangeStart = seedInt
			continue
		}

		seedRanges = append(seedRanges, valuesRange{
			start: rangeStart,
			end:   rangeStart + seedInt - 1,
		})

		rangeStart = 0
	}

	var remapper remapChain
	lastRemapper := &remapper

	for _, line := range input[1:] {

		// Every next line is either:
		// Indication of which to which range to remap
		// An actual remap
		// Or a blank line

		// We can ignore blank lines
		if line == "" {
			continue
		}

		if strings.HasSuffix(line, ":") {
			if len(lastRemapper.remaps) == 0 {
				continue
			}

			newRemapper := remapChain{}

			lastRemapper.next = &newRemapper
			lastRemapper = &newRemapper
			continue
		}

		lastRemapper.addRemap(remapRangeFromLine(line))
	}

	remapper.sortRemaps()

	var locations []valuesRange

	// Now we have a chain of remaps, we can iterate over the seeds to get the final value
	for _, seedRange := range seedRanges {
		locations = append(locations, remapper.remap(seedRange)...)
	}

	// Return lowest value
	lowest := locations[0].start

	for _, location := range locations[1:] {
		if location.start < lowest {
			lowest = location.start
		}
	}

	return lowest
}
