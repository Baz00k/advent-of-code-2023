package day04

import "github.com/Baz00k/advent-of-code-2023/helpers"

func SolvePartB(input []string) int {
	sum := 0
	cards := [](*scratchCard){}

	for _, line := range input {
		cards = append(cards, decodeCard(line))
	}

	for i, card := range cards {
		winningNumbers := 0
		for _, number := range card.cardNumbers {
			if helpers.Contains(card.winningNumbers, number) {
				winningNumbers++
			}
		}

		for j := i + 1; j < i+winningNumbers+1; j++ {
			cards[j].amount += card.amount
		}

		sum += card.amount
	}

	return sum
}
