package day04

import (
	"github.com/Baz00k/advent-of-code-2023/helpers"
)

func SolvePartA(input []string) int {
	sum := 0

	for _, line := range input {
		// We need to check if any of the card numbers are winning numbers
		card := decodeCard(line)
		cardValue := 0

		for _, number := range card.cardNumbers {
			if helpers.Contains(card.winningNumbers, number) {
				if cardValue == 0 {
					cardValue = 1
				} else {
					cardValue *= 2
				}
			}
		}

		sum += cardValue
	}

	return sum
}
