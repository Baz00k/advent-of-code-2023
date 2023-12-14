package day06

func SolvePartB(input []string) int {
	// A race consists of a pair of time and distance
	// Race is won by the boat that has traveled the furthest in the given time
	//
	// You start the race with 0 speed and 0 distance
	// When the race starts, for each millisecond you can increase your speed by 1
	// When you stop increasing your speed, you start moving with that speed
	// Once you start moving, you can't change your speed
	//
	// Record distance is the current best distance traveled in the given time
	// Our task is to find in how many ways we can beat that record

	times, recordDistances := parseInput(input, true)
	numberOfWays := make([]int, len(times))

	for i := 0; i < len(times); i++ {
		// Instead of brute forcing
		// We can find the lowest and highest speed we can use to beat the record
		// The number of ways to beat the record is the difference between the highest and lowest speed

		// We can use binary search to find the lowest and highest speed
		lowest := findLowestSpeedToBeatRecord(times[i], recordDistances[i])
		highest := findHighestSpeedToBeatRecord(times[i], recordDistances[i])

		numberOfWays[i] = highest - lowest + 1
	}

	// We need to multiply all the ways to beat the record
	result := 1
	for _, v := range numberOfWays {
		result *= v
	}

	return result
}
