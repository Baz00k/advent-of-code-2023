package day03

func SolvePartB(input [][]byte) int {
	sum := 0

	for y := 0; y < len(input); y++ {
		for x := 0; x < len(input[y]); x++ {
			isGear, gearRatio := isGear(input, x, y)
			
			if isGear {
				sum += gearRatio
			}
		}
	}

	return sum
}