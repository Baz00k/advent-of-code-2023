package day03

import "github.com/Baz00k/advent-of-code-2023/helpers"

func isDigit(b byte) bool {
	return b >= '0' && b <= '9'
}

// A sybol is anything that is not a digit or a dot
func isSymbol(b byte) bool {
	return !isDigit(b) && b != '.'
}

var directions = [][2]int{
	{0, -1},  // Up
	{1, 0},   // Right
	{0, 1},   // Down
	{-1, 0},  // Left
	{1, -1},  // Top right
	{1, 1},   // Bottom right
	{-1, 1},  // Bottom left
	{-1, -1}, // Top left
}

// Return iterator with coordinates of all adjacent cells
func getAdjacentCells(input [][]byte, x int, y int) [](struct{ x, y int }) {
	adjacentCells := [](struct{ x, y int }){}

	for _, direction := range directions {
		adjacentX := x + direction[0]
		adjacentY := y + direction[1]

		if adjacentX >= 0 && adjacentX < len(input[0]) && adjacentY >= 0 && adjacentY < len(input) {
			adjacentCells = append(adjacentCells, struct{ x, y int }{adjacentX, adjacentY})
		}
	}

	return adjacentCells
}

// Function that checks if any cell adjacent to the provided coordinates is a symbol
func isAdjacentToSymbol(input [][]byte, x int, y int) bool {
	for _, cell := range getAdjacentCells(input, x, y) {
		if isSymbol(input[cell.y][cell.x]) {
			return true
		}
	}

	return false
}

func isEndOfNumber(input [][]byte, x int, y int) bool {
	// If the cell to the right isn't a digit, the current digit is the end of the number
	return x == len(input[y])-1 || !isDigit(input[y][x+1])
}

func isStartOfNumber(input [][]byte, x int, y int) bool {
	// If the cell to the left isn't a digit, the current digit is the start of the number
	return x == 0 || !isDigit(input[y][x-1])
}

func getNumber(input [][]byte, x int, y int) int {
	if !isDigit(input[y][x]) {
		return 0
	}

	// Get the start and end of the number
	start := x
	end := x

	for !isStartOfNumber(input, start, y) {
		start--
	}

	for !isEndOfNumber(input, end, y) {
		end++
	}

	// Convert the number to an integer
	number := 0

	for i := start; i <= end; i++ {
		number *= 10
		number += int(input[y][i] - '0')
	}

	return number
}

// A gear is a '*' character that is adjacent to exactly 2 numbers.
// Function takes the input, the coordinates of the cell to check, and returns if the cell is a gear and the gear ratio
func isGear(input [][]byte, x int, y int) (bool, int) {
	if input[y][x] != '*' {
		return false, 0
	}

	numbers := [](int){}
	for _, cell := range getAdjacentCells(input, x, y) {
		number := getNumber(input, cell.x, cell.y)

		// Append only unique numbers
		// This assumes that gears are made of 2 UNIQUE numbers
		// This approach is far from perfect, but it works for the input and is fast enough
		// PR's are welcome
		if number != 0 && !helpers.Contains(numbers, number) {
			numbers = append(numbers, number)
		}
	}

	if len(numbers) != 2 {
		return false, 0
	}

	ratio := 1
	for _, number := range numbers {
		ratio *= number
	}

	return true, ratio
}
