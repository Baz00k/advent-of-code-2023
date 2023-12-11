package day03

func SolvePartA(input [][]byte) int {
	sum := 0

	// A number is a part number if any digit is adjacent to a symbol
	isPartNumber := false
	partNumber := 0

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			if !isDigit(input[y][x]) {
				continue
			}
			
			partNumber = partNumber * 10 + int(input[y][x] - '0')

			// If we already know that the current digit is part of a part number, we don't need to check if it is adjacent to a symbol
			if !isPartNumber && isAdjacentToSymbol(input, x, y) {
				isPartNumber = true
			}

			if isEndOfNumber(input, x, y) {
				if isPartNumber {
					// Add the part number to the sum
					sum += partNumber
				} 

				// Reset the part number
				partNumber = 0
				isPartNumber = false
			}
		}
	}

	return sum
}