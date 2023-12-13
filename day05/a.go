package day05

import (
	"strconv"
	"strings"
)

func SolvePartA(input []string) int {
	// First line is the list of seeds:
	// seeds: 79 14 55 ...
	seedsLine := strings.Split(input[0], " ")

	var seeds []valuesRange

	for _, seed := range seedsLine[1:] {
		seedInt, err := strconv.Atoi(seed)

		if err != nil {
			panic(err)
		}

		seeds = append(seeds, valuesRange{
			start: seedInt,
			end:   seedInt,
		})
	}

	var remapper remapChain
	lastRemapper := &remapper
	newRemapperNext := false

	for _, line := range input[1:] {

		// Every next line is either:
		// Indication of which to which range to remap
		// An actual remap
		// Or a blank line

		// We can ignore blank lines
		if line == "" {
			continue
		}

		if !newRemapperNext && strings.HasSuffix(line, ":") {
			newRemapperNext = true
			continue
		}

		if newRemapperNext && len(lastRemapper.remaps) > 0 {
			newRemapper := remapChain{}

			lastRemapper.next = &newRemapper
			lastRemapper = &newRemapper
		}

		newRemapperNext = false
		lastRemapper.addRemap(remapRangeFromLine(line))
	}

	remapper.sortRemaps()

	var locations []valuesRange

	// Now we have a chain of remaps, we can iterate over the seeds to get the final value
	for _, seed := range seeds {
		locations = append(locations, remapper.remap(seed)...)
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
