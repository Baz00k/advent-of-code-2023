package day01

func SolvePartA(input []string) int {
	sum := 0

	for _, line := range input {
		var firstDigitChar, lastDigitChar byte

		for _, char := range line {
			// If the character is not a digit, skip it
			if char < 48 || char > 57 {
				continue
			}

			// If the first digit is not set, set 
			if firstDigitChar == 0 {
				firstDigitChar = byte(char)
			}

			lastDigitChar = byte(char)
		}

		firstDigit := int(firstDigitChar - 48)
		lastDigit := int(lastDigitChar - 48)

		// Create a new integer from the first and last digits
		sum += firstDigit * 10 + lastDigit
	}

	return sum
}
