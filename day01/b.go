package day01

func SolvePartB(input []string) int {
	sum := 0

	for _, line := range input {
		var firstDigit, lastDigit int
		lineLength := len(line)

		searchEnd := 1

		// Search the line using a sliding window
		for searchBegin := 0; searchBegin < lineLength; searchEnd++ {
			linePart := line[searchBegin:searchEnd]

			if value, ok := validStrings[linePart]; ok {
				if firstDigit == 0 {
					firstDigit = value
				}

				lastDigit = value
			}

			if searchEnd == lineLength {
				searchBegin++
				searchEnd = searchBegin 
			}
		}

		sum += firstDigit * 10 + lastDigit
	}

	return sum
}

var validStrings = map[string]int{
	"1": 1,
	"2": 2,
	"3": 3,
	"4": 4,
	"5": 5,
	"6": 6,
	"7": 7,
	"8": 8,
	"9": 9,
	"one": 1,
	"two": 2,
	"three": 3,
	"four": 4,
	"five": 5,
	"six": 6,
	"seven": 7,
	"eight": 8,
	"nine": 9,
}