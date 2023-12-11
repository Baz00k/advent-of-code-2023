package day04

import (
	"strconv"
	"strings"
)

type scratchCard struct {
	cardNumbers    []int
	winningNumbers []int
	amount         int
}

func decodeCard(line string) *scratchCard {
	// Line is in format:
	// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
	// Where numbers before the pipe are winning numbers and numbers after the pipe are the card numbers
	parts := strings.Split(line, ":")
	numbers := strings.Split(parts[1], "|")

	if len(numbers) != 2 {
		panic("Invalid line")
	}

	scratchCard := scratchCard{amount: 1}

	// Convert the winning numbers to integers
	for _, winningNumber := range strings.Split(numbers[0], " ") {
		if winningNumber == "" {
			continue
		}

		number, err := strconv.Atoi(winningNumber)

		if err != nil {
			panic(err)
		}

		scratchCard.winningNumbers = append(scratchCard.winningNumbers, number)
	}

	// Convert the card numbers to integers
	for _, cardNumber := range strings.Split(numbers[1], " ") {
		if cardNumber == "" {
			continue
		}

		number, err := strconv.Atoi(cardNumber)

		if err != nil {
			panic(err)
		}

		scratchCard.cardNumbers = append(scratchCard.cardNumbers, number)
	}

	return &scratchCard
}
